// Package vo @Author:冯铁城 [17615007230@163.com] 2025-08-01 16:31:11
package grade

import (
	"net-project-edu_manage/internal/model/base"
)

// AcademicYearVo 学年VO
type AcademicYearVo struct {
	Name      string          `json:"name" gorm:"column:name"`
	StartDate base.FormatTime `json:"startDate" gorm:"column:start_date"`
	EndDate   base.FormatTime `json:"endDate" gorm:"column:end_date"`
	base.Vo
}
