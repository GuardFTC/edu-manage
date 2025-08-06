// Package res @Author:冯铁城 [17615007230@163.com] 2025-08-01 15:34:45
package res

import (
	"net-project-edu_manage/internal/model/base"
)

// PageResult 分页结果
type PageResult[T any] struct {
	PageNum  int   `json:"pageNum"`
	PageSize int   `json:"pageSize"`
	Total    int64 `json:"total"`
	Data     []T   `json:"data"`
}

// CreatePageResult 创建分页结果
func CreatePageResult[T any](r *base.Request, total int64, data []T) *PageResult[T] {

	//1.空值处理
	if data == nil {
		data = []T{}
	}

	//2.封装返回
	return &PageResult[T]{
		PageNum:  r.PageNum,
		PageSize: r.PageSize,
		Total:    total,
		Data:     data,
	}
}
