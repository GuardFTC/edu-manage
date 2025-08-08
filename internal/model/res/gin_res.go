// Package res @Author:冯铁城 [17615007230@163.com] 2025-07-30 19:56:47
package res

import (
	"strings"

	"github.com/gin-gonic/gin"
)

// UnProcessTag 流程异常标记
const UnProcessTag = "unprocess"

// NotFoundTag 无法找到目标资源标记
const NotFoundTag = "not found"

// FailResToCByMsg 错误结果集转成JSON返回
func FailResToCByMsg(c *gin.Context, msg string) {

	//1.定义处理方法
	var f func(msg string) *Result

	//2.根据异常信息获取处理方法
	if strings.Contains(msg, NotFoundTag) {
		f = NotFoundFail
	} else if strings.Contains(msg, UnProcessTag) {
		f = UnProcessFail
	} else {
		f = ServerFail
	}

	//3.FailResToC
	FailResToC(c, f, msg)
}

// FailResToC 错误结果集转成JSON返回
func FailResToC(c *gin.Context, f func(msg string) *Result, msg string) {

	//1.获取res
	failRes := f(msg)

	//2.返回JSON
	c.AbortWithStatusJSON(failRes.Code, failRes.ToJson())
}

// SuccessResToC 成功结果集转成JSON返回
func SuccessResToC(c *gin.Context, f func(data interface{}) *Result, data interface{}) {

	//1.获取res
	successRes := f(data)

	//2.返回JSON
	c.JSON(successRes.Code, successRes.ToJson())
}
