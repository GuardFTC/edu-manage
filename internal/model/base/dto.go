// Package base @Author:冯铁城 [17615007230@163.com] 2025-08-04 15:14:29
package base

import (
	"net-project-edu_manage/internal/common/util"

	"github.com/gin-gonic/gin"
)

// Dto 基础DTO
type Dto struct {
	ID        int64  `json:"id" form:"id" binding:""` // 系统用户ID
	CreatedBy string `json:"-"`                       // 创建人ID
	UpdatedBy string `json:"-"`                       // 最后修改人ID
}

// SetCreateByAndUpdateBy 设置创建人ID和最后修改人ID
func (d *Dto) SetCreateByAndUpdateBy(c *gin.Context) {

	//1.获取操作人邮箱
	email, ok := util.GetEmailFromC(c)
	if !ok {
		return
	}

	//2.设置创建人ID和最后修改人ID
	d.CreatedBy = email
	d.UpdatedBy = email
}

// SetUpdateBy 设置最后修改人ID
func (d *Dto) SetUpdateBy(c *gin.Context) {

	//1.获取操作人邮箱
	email, ok := util.GetEmailFromC(c)
	if !ok {
		return
	}

	//2.设置创建人ID和最后修改人ID
	d.UpdatedBy = email
}
