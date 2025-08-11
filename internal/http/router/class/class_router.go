// Package class @Author:冯铁城 [17615007230@163.com] 2025-08-11 11:01:55
package class

import (
	"net-project-edu_manage/internal/http/handler/class"

	"github.com/gin-gonic/gin"
)

// InitClassRouter 初始化班级管理路由
func InitClassRouter(v *gin.RouterGroup, tokenHandler gin.HandlerFunc) *gin.RouterGroup {

	//1.定义班级管理路由组
	classRouter := v.Group("classes")
	classRouter.Use(tokenHandler)

	//2.定义接口路由
	classRouter.POST("", class.AddClass)
	classRouter.DELETE("", nil)
	classRouter.GET(":id", nil)
	classRouter.PUT(":id", nil)
	classRouter.GET("", nil)

	//3.返回
	return classRouter
}
