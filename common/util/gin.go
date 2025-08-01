// Package util @Author:冯铁城 [17615007230@163.com] 2025-07-30 19:56:47
package util

import (
	"github.com/gin-gonic/gin"
	"net-project-edu_manage/common/res"
	"strings"
)

// FormatMsg 将多行原始错误信息按换行符分割成字符串数组
func FormatMsg(msg string) string {

	//1.按行切分错误信息
	lines := strings.Split(msg, "\n")

	//2.定义结果集
	var result []string

	//3.遍历每一行错误信息,去掉多余前缀，只保留字段和错误说明
	for _, line := range lines {
		if idx := strings.Index(line, "Error:"); idx != -1 {
			result = append(result, "- "+line[idx+len("Error:"):])
		} else {
			result = append(result, "- "+line)
		}
	}

	//5.拼接所有错误信息返回
	return strings.Join(result, "\n")
}

// FailResToCByMsg 错误结果集转成JSON返回
func FailResToCByMsg(c *gin.Context, msg string) {

	//1.定义处理方法
	var f func(msg string) *res.Result

	//2.根据异常信息获取处理方法
	if strings.Contains(msg, "not found") {
		f = res.NotFoundFail
	} else if strings.Contains(msg, "unprocess") {
		f = res.UnProcessFail
	} else {
		f = res.ServerFail
	}

	//3.FailResToC
	FailResToC(c, f, msg)
}

// FailResToC 错误结果集转成JSON返回
func FailResToC(c *gin.Context, f func(msg string) *res.Result, msg string) {

	//1.获取res
	failRes := f(msg)

	//2.返回JSON
	c.JSON(failRes.Code, failRes.ToJson())
}

// SuccessResToC 成功结果集转成JSON返回
func SuccessResToC(c *gin.Context, f func(data interface{}) *res.Result, data interface{}) {

	//1.获取res
	successRes := f(data)

	//2.返回JSON
	c.JSON(successRes.Code, successRes.ToJson())
}
