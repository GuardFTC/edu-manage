// Package vo @Author:冯铁城 [17615007230@163.com] 2025-08-04 20:29:50
package vo

type BaseVo struct {
	CreatedAt   FormatTime `json:"createdTime" gorm:"column:created_at"`
	UpdatedAt   FormatTime `json:"updatedTime" gorm:"column:updated_at"`
	CreatedUser string     `json:"createdUser" gorm:"column:created_user"`
	UpdatedUser string     `json:"updatedUser" gorm:"column:updated_user"`
}
