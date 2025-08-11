// Package class @Author:冯铁城 [17615007230@163.com] 2025-08-11 10:57:31
package class

import "net-project-edu_manage/internal/model/base"

// ClassDto 班级DTO
type ClassDto struct {
	Name           string `gorm:"column:name" json:"name" binding:"required,max=32"`                     // 班级名称
	GradeId        int64  `gorm:"column:grade_id" json:"gradeId" binding:"required,gt=0"`                // 年级ID
	AcademicYearId int64  `gorm:"column:academic_year_id" json:"academicYearId" binding:"required,gt=0"` // 学年ID
	base.Dto
}
