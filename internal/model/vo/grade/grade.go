// Package grade @Author:冯铁城 [17615007230@163.com] 2025-08-01 16:31:11
package grade

import (
	"net-project-edu_manage/internal/model/base"
)

// GradeVo 年级VO
type GradeVo struct {
	Name string `json:"name" gorm:"column:name"`
	base.Vo
}
