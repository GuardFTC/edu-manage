// Package service @Author:冯铁城 [17615007230@163.com] 2025-08-04 16:37:12
package service

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"net-project-edu_manage/core/db"
	"net-project-edu_manage/dao/model"
	"net-project-edu_manage/dao/query"
	"net-project-edu_manage/model/dto"
	"sync"
)

// AcademicYearService 学年服务
type AcademicYearService struct {
	sync sync.Mutex //预留锁 并发高时使用
}

// Add 新增
func (s *AcademicYearService) Add(c *gin.Context, academicYearDto *dto.AcademicYearDto) error {
	return db.Q.Transaction(func(tx *query.Query) error {

		//1.设置名称
		academicYearDto.SetName()

		//2.设置创建人与修改人
		academicYearDto.SetCreateByAndUpdateBy(c)

		//3.dto to po
		var academicYear model.AcademicYear
		if err := copier.Copy(&academicYear, &academicYearDto); err != nil {
			return err
		}

		//4.保存
		if err := tx.AcademicYear.WithContext(c).Create(&academicYear); err != nil {
			return err
		}

		//5.ID回写
		academicYearDto.ID = academicYear.ID

		//6.默认返回
		return nil
	})
}
