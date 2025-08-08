// Package system @Author:冯铁城 [17615007230@163.com] 2025-08-01 16:31:11
package system

import "net-project-edu_manage/internal/model/base"

// SystemUserVo 系统用户VO
type SystemUserVo struct {
	Name  string `json:"name" gorm:"column:name"`
	Email string `json:"email" gorm:"column:email"`
	base.Vo
}
