// Package grade @Author:冯铁城 [17615007230@163.com] 2025-08-05 16:26:37
package grade

import (
	"net-project-edu_manage/internal/model/base"
)

// GradeRequest 年级请求参数
type GradeRequest struct {
	Name   string `json:"name" form:"name" binding:"omitempty"`
	IsList bool   `json:"isList" form:"isList" binding:"omitempty"`
	base.Request
}
