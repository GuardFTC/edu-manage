// Package vo @Author:冯铁城 [17615007230@163.com] 2025-08-01 16:31:11
package vo

import "time"

// SystemUserVo 系统用户VO
type SystemUserVo struct {
	ID          int64     `json:"id" gorm:"column:id"`
	Name        string    `json:"name" gorm:"column:name"`
	Email       string    `json:"email" gorm:"column:email"`
	CreatedAt   time.Time `json:"createdTime" gorm:"column:created_at"`
	UpdatedAt   time.Time `json:"updatedTime" gorm:"column:updated_at"`
	CreatedUser string    `json:"createdUser" gorm:"column:created_user"`
	UpdatedUser string    `json:"updatedUser" gorm:"column:updated_user"`
}
