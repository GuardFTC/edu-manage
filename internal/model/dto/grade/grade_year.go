// Package grade @Author:冯铁城 [17615007230@163.com] 2025-08-08 16:45:00
package grade

import "net-project-edu_manage/internal/model/base"

// GradeYearDto 学年-年级关联DTO
type GradeYearDto struct {
	GradeId        int64 `json:"gradeId" binding:"required,gt=0"`
	AcademicYearId int64 `json:"academicYearId" binding:"required,gt=0"`
	base.Dto
}
