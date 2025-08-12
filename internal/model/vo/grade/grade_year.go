// Package grade @Author:冯铁城 [17615007230@163.com] 2025-08-01 16:31:11
package grade

import (
	"net-project-edu_manage/internal/model/base"
)

// GradeYearVo 年级-学年关联VO
type GradeYearVo struct {
	GradeId        int64 `json:"gradeId" gorm:"column:grade_id"`
	AcademicYearId int64 `json:"academicYearId" gorm:"column:academic_year_id"`
	base.SimpleVo
}
