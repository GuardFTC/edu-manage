// Package grade @Author:冯铁城 [17615007230@163.com] 2025-08-05 16:26:37
package grade

import (
	"net-project-edu_manage/internal/model/base"
)

// GradeYearRequest 年级-学年关联请求参数
type GradeYearRequest struct {
	GradeId        int64 `json:"gradeId" form:"gradeId" binding:"omitempty,gt=0"`
	AcademicYearId int64 `json:"academicYearId" form:"academicYearId" binding:"omitempty,gt=0"`
	base.Request
}
