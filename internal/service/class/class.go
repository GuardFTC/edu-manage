// Package class @Author:冯铁城 [17615007230@163.com] 2025-08-11 16:29:18
package class

import (
	"errors"
	"net-project-edu_manage/internal/infrastructure/db"
	"net-project-edu_manage/internal/infrastructure/db/master/model"
	"net-project-edu_manage/internal/infrastructure/db/master/query"
	dtoPack "net-project-edu_manage/internal/model/dto/class"
	"net-project-edu_manage/internal/model/res"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cast"
)

// ClassService 班级服务
type ClassService struct {
	sync sync.Mutex //预留锁 并发高时使用
}

// Add 添加班级
func (s *ClassService) Add(c *gin.Context, classDto *dtoPack.ClassDto) error {
	return db.GetDefaultQuery().Transaction(func(tx *query.Query) error {

		//1.通过学年ID、班级ID查询关联ID
		gradeYear, err := tx.WithContext(c).GradeYear.Select(tx.GradeYear.ID).Where(
			tx.GradeYear.AcademicYearID.Eq(classDto.AcademicYearId),
			tx.GradeYear.GradeID.Eq(classDto.GradeId),
		).First()
		if err != nil {
			return err
		}

		//2.设置创建人与修改人
		classDto.SetCreateByAndUpdateBy(c)

		//3.dto to po
		var class model.Class
		if err = copier.Copy(&class, &classDto); err != nil {
			return err
		}

		//4.设置关联ID
		class.GradeYearID = gradeYear.ID

		//5.保存
		if err = tx.Class.WithContext(c).Create(&class); err != nil {
			return err
		}

		//6.ID回写
		classDto.ID = class.ID

		//7.默认返回
		return nil
	})
}

// Delete 删除班级
func (s *ClassService) Delete(c *gin.Context, ids []string) error {
	return db.GetDefaultQuery().Transaction(func(tx *query.Query) error {

		//1.id string 转 int64
		intIds := cast.ToInt64Slice(ids)

		//2.TODO 查询是否有关联的学生，如果有无法删除

		//3.TODO 查询是否有关联的课程，如果有无法删除

		//4.删除班级
		if delRes, err := tx.Class.WithContext(c).Where(tx.Class.ID.In(intIds...)).Delete(); err != nil {
			return err
		} else {
			log.Printf("删除班级成功,删除数量:%d", delRes.RowsAffected)
		}

		//5.默认返回
		return nil
	})
}

// Get 获取班级
func (s *ClassService) Get(c *gin.Context, id string) (*dtoPack.ClassDto, error) {

	//1.id string 转 int64
	intId := cast.ToInt64(id)

	//2.设置别名
	dq := db.GetDefaultQuery()
	cl := dq.Class.As("c")
	gy := dq.GradeYear.As("gy")

	//3.查询
	var classDto dtoPack.ClassDto
	if err := cl.WithContext(c).
		Select(cl.ID, cl.Name, gy.AcademicYearID, gy.GradeID).
		Join(gy, gy.ID.EqCol(cl.GradeYearID)).
		Where(cl.ID.Eq(intId)).
		Scan(&classDto); err != nil {
		return nil, err
	}

	//4.为空返回异常
	if classDto.ID == 0 {
		return nil, errors.New("record " + res.NotFoundTag)
	}

	//5.返回dto
	return &classDto, nil
}
