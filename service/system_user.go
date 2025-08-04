// Package service @Author:冯铁城 [17615007230@163.com] 2025-07-30 19:41:24
package service

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"net-project-edu_manage/common/res"
	"net-project-edu_manage/common/util"
	"net-project-edu_manage/core/db"
	"net-project-edu_manage/dao/model/system"
	"net-project-edu_manage/dao/query"
	"net-project-edu_manage/model/dto"
	"net-project-edu_manage/model/request"
	"net-project-edu_manage/model/vo"
	"sync"
)

type SystemUserService struct {
	sync sync.Mutex //预留锁 并发高时使用
}

// Add 新增系统用户
func (sys *SystemUserService) Add(c *gin.Context, systemUserDTO *dto.SystemUserDto) error {
	return db.Q.Transaction(func(tx *query.Query) error {

		//1.密码加密
		if password, err := util.HashPassword(systemUserDTO.Password); err != nil {
			return err
		} else {
			systemUserDTO.Password = password
		}

		//2.设置创建人、修改人
		systemUserDTO.SetCreateByAndUpdateBy(c)

		//3.dto to po
		var systemUser system.SystemUser
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
	return db.Q.Transaction(func(tx *query.Query) error {

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
func (sys *SystemUserService) Get(c *gin.Context, id string) (*dto.SystemUserDto, error) {

	//1.id string 转 int64
	intId := cast.ToInt64(id)

	//2.查询系统用户
	systemUser, err := db.Q.SystemUser.WithContext(c).Where(db.Q.SystemUser.ID.Eq(intId)).First()
	if err != nil {
		return nil, err
	}

	//3.po to dto
	var systemUserDTO dto.SystemUserDto
	if err = copier.Copy(&systemUserDTO, &systemUser); err != nil {
		return nil, err
	}

	//4.密码清空
	systemUserDTO.Password = ""

	//5.返回dto
	return &systemUserDTO, nil
}

// Update 修改系统用户
func (sys *SystemUserService) Update(c *gin.Context, id string, systemUserDTO *dto.SystemUserDto) error {
	return db.Q.Transaction(func(tx *query.Query) error {

		//1.id string 转 int64
		intId := cast.ToInt64(id)

		//2.查询系统用户
		systemUser, err := db.Q.SystemUser.WithContext(c).Where(db.Q.SystemUser.ID.Eq(intId)).First()
		if err != nil {
			return err
		}

		//3.密码回写
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

		//7.ID回写,密码清空
		systemUserDTO.ID = systemUser.ID
		systemUserDTO.Password = ""

		//8.返回
		return nil
	})
}

// Page 分页查询系统用户
func (sys *SystemUserService) Page(c *gin.Context, request *request.SystemUserRequest) (*res.PageResult[*vo.SystemUserVo], error) {

	//1.分页参数默认处理
	request.DefaultPage()

	//2.设置别名，利于后续Join查询
	s := db.Q.SystemUser.As("s")
	s1 := db.Q.SystemUser.As("s1")
	s2 := db.Q.SystemUser.As("s2")

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
	var systemUsersVo []*vo.SystemUserVo
	if err := context.Scan(&systemUsersVo); err != nil {
		return nil, err
	}

	//6.查询总数
	total, err := countContext.Count()
	if err != nil {
		return nil, err
	}

	//7.封装分页结果
	return res.CreatePageResult[*vo.SystemUserVo](&request.BaseRequest, total, systemUsersVo), nil
}
