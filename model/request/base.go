// Package request @Author:冯铁城 [17615007230@163.com] 2025-08-01 15:40:27
package request

// BaseRequest 基础查询参数
type BaseRequest struct {
	PageNum  int `json:"pageNum" form:"pageNum" binding:"omitempty,min=1,max=1000"`
	PageSize int `json:"pageSize" form:"pageSize" binding:"omitempty,min=1,max=100"`
}

// DefaultPage 设置默认参数
func (p *BaseRequest) DefaultPage() {
	if p.PageNum <= 0 {
		p.PageNum = 1
	}
	if p.PageSize <= 0 {
		p.PageSize = 10
	}
}

// GetSkip 获取跳过的数据
func (p *BaseRequest) GetSkip() int {
	return (p.PageNum - 1) * p.PageSize
}
