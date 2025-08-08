// Package grade @Author:冯铁城 [17615007230@163.com] 2025-08-08 11:25:02
package grade

import (
	"net-project-edu_manage/internal/infrastructure/db"
	"net-project-edu_manage/internal/infrastructure/db/master/model"
	"net-project-edu_manage/internal/infrastructure/db/master/query"
	dtoPack "net-project-edu_manage/internal/model/dto/grade"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/spf13/cast"
)

// GradeService 年级服务
type GradeService struct {
	sync sync.Mutex //预留锁 并发高时使用
}

// Add 新增
func (s *GradeService) Add(c *gin.Context, yearId string, gradeDto *dtoPack.GradeDto) error {
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

		//4.创建年级-学年关联
		var gradeYear model.GradeYear

		//5.设置创建人与修改人
		if err := copier.Copy(&gradeYear, &gradeDto); err != nil {
			return err
		}

		//6.设置年级ID与学年ID
		gradeYear.GradeID = grade.ID
		gradeYear.AcademicYearID = cast.ToInt64(yearId)

		//7.保存年级-学年关联
		if err := tx.GradeYear.WithContext(c).Create(&gradeYear); err != nil {
			return err
		}

		//8.ID回写
		gradeDto.ID = grade.ID

		//9.默认返回
		return nil
	})
}
