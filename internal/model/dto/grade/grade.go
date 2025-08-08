// Package grade @Author:冯铁城 [17615007230@163.com] 2025-08-08 11:38:45
package grade

import (
	"net-project-edu_manage/internal/model/base"
)

// GradeDto 年级DTO
type GradeDto struct {
	Name string `gorm:"column:name" json:"name" form:"name" binding:"required,max=32"` // 名称
	base.Dto
}
