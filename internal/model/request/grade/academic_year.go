// Package grade @Author:冯铁城 [17615007230@163.com] 2025-08-05 16:26:37
package grade

import (
	"net-project-edu_manage/internal/model/base"
)

// AcademicYearRequest 学年请求参数
type AcademicYearRequest struct {
	Name           string `json:"name" form:"name" binding:"omitempty"`
	StartDateBegin string `json:"startDateBegin" form:"startDateBegin" binding:"omitempty,datetime=2006-01-02" `
	StartDateEnd   string `json:"startDateEnd" form:"startDateEnd" binding:"omitempty,datetime=2006-01-02" `
	EndDateBegin   string `json:"endDateBegin" form:"endDateBegin" binding:"omitempty,datetime=2006-01-02"`
	EndDateEnd     string `json:"endDateEnd" form:"endDateEnd" binding:"omitempty,datetime=2006-01-02"`
	base.Request
}
