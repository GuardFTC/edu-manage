// Package system @Author:冯铁城 [17615007230@163.com] 2025-07-30 19:41:24
package system

import (
	"net-project-edu_manage/internal/common/util"
	"net-project-edu_manage/internal/infrastructure/db"
	"net-project-edu_manage/internal/infrastructure/db/master/model"
	"net-project-edu_manage/internal/infrastructure/db/master/query"
	dtoPack "net-project-edu_manage/internal/model/dto/system"
	reqPack "net-project-edu_manage/internal/model/request/system"
	"net-project-edu_manage/internal/model/res"
	voPack "net-project-edu_manage/internal/model/vo/system"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cast"
)

// SystemUserService 系统用户服务
type SystemUserService struct {
	sync sync.Mutex //预留锁 并发高时使用
}

// Add 新增系统用户
func (sys *SystemUserService) Add(c *gin.Context, systemUserDTO *dtoPack.SystemUserDto) error {
	return db.GetDefaultQuery().Transaction(func(tx *query.Query) error {

		//1.密码加密
		if password, err := util.HashPassword(systemUserDTO.Password); err != nil {
			return err
		} else {
			systemUserDTO.Password = password
		}

		//2.设置创建人、修改人
		systemUserDTO.SetCreateByAndUpdateBy(c)

		//3.dto to po
		var systemUser model.SystemUser
		if err := copier.Copy(&systemUser, &systemUserDTO); err != nil {
			return err
		}

		//4.入库保存
		if err := tx.SystemUser.WithContext(c).Create(&systemUser); err != nil {
			return err
		}

		//5.ID回写,密码清空
		systemUserDTO.ID = systemUser.ID
		systemUserDTO.Password = ""

		//6.默认返回
		return nil
	})
}

// Delete 删除系统用户
func (sys *SystemUserService) Delete(c *gin.Context, ids []string) error {
	return db.GetDefaultQuery().Transaction(func(tx *query.Query) error {

		//1.id string 转 int64
		intIds := cast.ToInt64Slice(ids)

		//2.删除系统用户
		if delRes, err := tx.SystemUser.WithContext(c).Where(tx.SystemUser.ID.In(intIds...)).Delete(); err != nil {
			return err
		} else {
			log.Printf("删除系统用户成功,删除数量:%d", delRes.RowsAffected)
			return nil
		}
	})
}

// Get 获取系统用户
func (sys *SystemUserService) Get(c *gin.Context, id string) (*dtoPack.SystemUserDto, error) {

	//1.id string 转 int64
	intId := cast.ToInt64(id)

	//2.查询系统用户
	s := db.GetDefaultQuery().SystemUser
	systemUser, err := s.WithContext(c).Where(s.ID.Eq(intId)).First()
	if err != nil {
		return nil, err
	}

	//3.po to dto
	var systemUserDTO dtoPack.SystemUserDto
	if err = copier.Copy(&systemUserDTO, &systemUser); err != nil {
		return nil, err
	}

	//4.密码清空
	systemUserDTO.Password = ""

	//5.返回dto
	return &systemUserDTO, nil
}

// Update 修改系统用户
func (sys *SystemUserService) Update(c *gin.Context, id string, systemUserDTO *dtoPack.SystemUserDto) error {
	return db.GetDefaultQuery().Transaction(func(tx *query.Query) error {

		//1.id string 转 int64
		intId := cast.ToInt64(id)

		//2.查询系统用户
		systemUser, err := tx.SystemUser.WithContext(c).Where(tx.SystemUser.ID.Eq(intId)).First()
		if err != nil {
			return err
		}

		//3.ID,密码回写
		systemUserDTO.ID = intId
		systemUserDTO.Password = systemUser.Password

		//4.设置修改人
		systemUserDTO.SetUpdateBy(c)

		//5.dto to po
		if err = copier.Copy(&systemUser, &systemUserDTO); err != nil {
			return err
		}

		//6.更新
		if updateRes, err := tx.SystemUser.WithContext(c).Where(tx.SystemUser.ID.Eq(intId)).Updates(&systemUser); err != nil {
			return err
		} else {
			log.Printf("更新系统用户成功,更新数量:%d", updateRes.RowsAffected)
		}

		//7.密码清空
		systemUserDTO.Password = ""

		//8.返回
		return nil
	})
}

// Page 分页查询系统用户
func (sys *SystemUserService) Page(c *gin.Context, request *reqPack.SystemUserRequest) (*res.PageResult[*voPack.SystemUserVo], error) {

	//1.分页参数默认处理
	request.DefaultPage()

	//2.设置别名，利于后续Join查询
	s := db.GetDefaultQuery().SystemUser.As("s")
	s1 := db.GetDefaultQuery().SystemUser.As("s1")
	s2 := db.GetDefaultQuery().SystemUser.As("s2")

	//3.封装查询参数
	context := s.WithContext(c)
	if request.Name != "" {
		context = context.Where(s.Name.Like("%" + request.Name + "%"))
	}
	if request.Email != "" {
		context = context.Where(s.Email.Like("%" + request.Email + "%"))
	}

	//4.暂存总数查询参数
	countContext := context

	//4.设置查询字段，排序，分页参数
	context = context.
		Select(s.ID, s.Name, s.Email, s.CreatedAt, s.UpdatedAt, s1.Name.As("created_user"), s2.Name.As("updated_user")).
		Join(s1, s.CreatedBy.EqCol(s1.Email)).
		Join(s2, s.UpdatedBy.EqCol(s2.Email)).
		Order(s.ID.Desc()).
		Offset(request.GetSkip()).Limit(request.PageSize)

	//5.查询数据
	var systemUsersVos []*voPack.SystemUserVo
	if err := context.Scan(&systemUsersVos); err != nil {
		return nil, err
	}

	//6.查询总数
	total, err := countContext.Count()
	if err != nil {
		return nil, err
	}

	//7.封装分页结果
	return res.CreatePageResult[*voPack.SystemUserVo](&request.Request, total, systemUsersVos), nil
}
