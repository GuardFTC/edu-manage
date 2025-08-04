// Package dto @Author:冯铁城 [17615007230@163.com] 2025-08-04 15:14:29
package dto

import (
	"github.com/gin-gonic/gin"
	"net-project-edu_manage/common/util"
)

type BaseDto struct {
	ID        int64  `json:"id" form:"id" binding:""` // 系统用户ID
	CreatedBy string `json:"-"`                       // 创建人ID
	UpdatedBy string `json:"-"`                       // 最后修改人ID
}

// SetCreateByAndUpdateBy 设置创建人ID和最后修改人ID
func (b *BaseDto) SetCreateByAndUpdateBy(c *gin.Context) {

	//1.获取操作人邮箱
	email, ok := util.GetEmailFromC(c)
	if !ok {
		return
	}

	//2.设置创建人ID和最后修改人ID
	b.CreatedBy = email
	b.UpdatedBy = email
}

// SetUpdateBy 设置最后修改人ID
func (b *BaseDto) SetUpdateBy(c *gin.Context) {

	//1.获取操作人邮箱
	email, ok := util.GetEmailFromC(c)
	if !ok {
		return
	}

	//2.设置创建人ID和最后修改人ID
	b.UpdatedBy = email
}
