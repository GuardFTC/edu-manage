// Package base @Author:冯铁城 [17615007230@163.com] 2025-08-01 15:40:27
package base

// Request 基础查询参数
type Request struct {
	PageNum  int `json:"pageNum" form:"pageNum" binding:"omitempty,min=1,max=1000"`
	PageSize int `json:"pageSize" form:"pageSize" binding:"omitempty,min=1,max=100"`
}

// DefaultPage 设置默认参数
func (r *Request) DefaultPage() {
	if r.PageNum <= 0 {
		r.PageNum = 1
	}
	if r.PageSize <= 0 {
		r.PageSize = 10
	}
}

// GetSkip 获取跳过的数据
func (r *Request) GetSkip() int {
	return (r.PageNum - 1) * r.PageSize
}
