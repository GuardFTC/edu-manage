// Package dto @Author:冯铁城 [17615007230@163.com] 2025-07-30 17:33:59
package dto

import "net-project-edu_manage/internal/model/base"

// SystemUserDto 系统用户DTO
type SystemUserDto struct {
	Name     string `json:"name" form:"name" binding:"required,max=64"`          // 用户名
	Email    string `json:"email" form:"email" binding:"required,email,max=128"` // 邮箱
	Password string `json:"password" form:"password" binding:"required,max=128"` //密码
	base.Dto
}
