// Package system @Author:冯铁城 [17615007230@163.com] 2025-08-01 15:48:56
package system

import "net-project-edu_manage/internal/model/base"

// SystemUserRequest 系统用户查询参数结构体
type SystemUserRequest struct {
	Name  string `json:"name" form:"name" binding:""`
	Email string `json:"email" form:"email" binding:""`
	base.Request
}
