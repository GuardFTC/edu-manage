// Package grade @Author:冯铁城 [17615007230@163.com] 2025-08-08 11:25:02
package grade

import (
	"errors"
	"net-project-edu_manage/internal/infrastructure/db"
	"net-project-edu_manage/internal/infrastructure/db/master/model"
	"net-project-edu_manage/internal/infrastructure/db/master/query"
	dtoPack "net-project-edu_manage/internal/model/dto/grade"
	reqPack "net-project-edu_manage/internal/model/request/grade"
	"net-project-edu_manage/internal/model/res"
	voPack "net-project-edu_manage/internal/model/vo/grade"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cast"
)

// GradeService 年级服务
type GradeService struct {
	sync sync.Mutex //预留锁 并发高时使用
}

// Add 新增
func (s *GradeService) Add(c *gin.Context, gradeDto *dtoPack.GradeDto) error {
	return db.GetDefaultQuery().Transaction(func(tx *query.Query) error {

		//1.设置创建人与修改人
		gradeDto.SetCreateByAndUpdateBy(c)

		//2.dto to po
		var grade model.Grade
		if err := copier.Copy(&grade, &gradeDto); err != nil {
			return err
		}

		//3.保存
		if err := tx.Grade.WithContext(c).Create(&grade); err != nil {
			return err
		}

		//4.ID回写
		gradeDto.ID = grade.ID

		//5.默认返回
		return nil
	})
}

// Delete 删除年级
func (s *GradeService) Delete(c *gin.Context, ids []string) error {
	return db.GetDefaultQuery().Transaction(func(tx *query.Query) error {

		//1.id string 转 int64
		intIds := cast.ToInt64Slice(ids)

		//2.查询年级存在关联的年级
		count, err := tx.GradeYear.WithContext(c).Where(tx.GradeYear.GradeID.In(intIds...)).Count()
		if err != nil {
			return err
		}
		if count > 0 {
			return errors.New("grade is associated with academic year! can't delete " + res.UnProcessTag)
		}

		//3.删除年级
		if delRes, err := tx.Grade.WithContext(c).Where(tx.Grade.ID.In(intIds...)).Delete(); err != nil {
			return err
		} else {
			log.Printf("删除年级成功,删除数量:%d", delRes.RowsAffected)
		}

		//4.默认返回
		return nil
	})
}

// Get 获取年级
func (s *GradeService) Get(c *gin.Context, id string) (*dtoPack.GradeDto, error) {

	//1.id string 转 int64
	intId := cast.ToInt64(id)

	//2.查询年级
	a := db.GetDefaultQuery().Grade
	grade, err := a.WithContext(c).Where(a.ID.Eq(intId)).First()
	if err != nil {
		return nil, err
	}

	//3.po to dto
	var gradeDto dtoPack.GradeDto
	if err = copier.Copy(&gradeDto, &grade); err != nil {
		return nil, err
	}

	//4.返回dto
	return &gradeDto, nil
}

// Update 修改年级
func (s *GradeService) Update(c *gin.Context, id string, gradeDto *dtoPack.GradeDto) error {
	return db.GetDefaultQuery().Transaction(func(tx *query.Query) error {

		//1.id string 转 int64
		intId := cast.ToInt64(id)

		//2.查询年级
		grade, err := tx.Grade.WithContext(c).Where(tx.Grade.ID.Eq(intId)).First()
		if err != nil {
			return err
		}

		//3.设置ID
		gradeDto.ID = intId

		//4.设置修改人
		gradeDto.SetUpdateBy(c)

		//5.dto to po
		if err = copier.Copy(&grade, &gradeDto); err != nil {
			return err
		}

		//6.更新
		if updateRes, err := tx.Grade.WithContext(c).Where(tx.Grade.ID.Eq(intId)).Updates(&grade); err != nil {
			return err
		} else {
			log.Printf("更新年级成功,更新数量:%d", updateRes.RowsAffected)
		}

		//7.返回
		return nil
	})
}

// Page 分页查询年级
func (s *GradeService) Page(c *gin.Context, request *reqPack.GradeRequest) (*res.PageResult[*voPack.GradeVo], error) {

	//1.分页参数默认处理
	request.DefaultPage()

	//2.设置别名，利于后续Join查询
	g := db.GetDefaultQuery().Grade.As("g")
	s1 := db.GetDefaultQuery().SystemUser.As("s1")
	s2 := db.GetDefaultQuery().SystemUser.As("s2")

	//3.封装查询参数
	context := g.WithContext(c)
	if request.Name != "" {
		context = context.Where(g.Name.Like("%" + request.Name + "%"))
	}

	//4.暂存总数查询参数
	countContext := context

	//5.设置查询字段，排序，分页参数
	context = context.
		Select(g.ID, g.Name, g.CreatedAt, g.UpdatedAt, s1.Name.As("created_user"), s2.Name.As("updated_user")).
		Join(s1, g.CreatedBy.EqCol(s1.Email)).
		Join(s2, g.UpdatedBy.EqCol(s2.Email)).
		Order(g.ID.Desc()).
		Offset(request.GetSkip()).Limit(request.PageSize)

	//6.查询数据
	var gradeVos []*voPack.GradeVo
	if err := context.Scan(&gradeVos); err != nil {
		return nil, err
	}

	//7.查询总数
	total, err := countContext.Count()
	if err != nil {
		return nil, err
	}

	//8.封装分页结果
	return res.CreatePageResult[*voPack.GradeVo](&request.Request, total, gradeVos), nil
}
