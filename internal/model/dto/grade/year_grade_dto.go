// Package grade @Author:冯铁城 [17615007230@163.com] 2025-08-08 16:45:00
package grade

// YearGradeDto 学年-年级关联DTO
type YearGradeDto struct {
	GradeIDs []int64 `json:"gradeIds" binding:"required,gt=0"`
}
