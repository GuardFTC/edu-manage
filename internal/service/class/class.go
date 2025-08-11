// Package class @Author:冯铁城 [17615007230@163.com] 2025-08-11 16:29:18
package class

import (
	"net-project-edu_manage/internal/infrastructure/db"
	"net-project-edu_manage/internal/infrastructure/db/master/model"
	"net-project-edu_manage/internal/infrastructure/db/master/query"
	dtoPack "net-project-edu_manage/internal/model/dto/class"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
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
