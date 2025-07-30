// Package server @Author:冯铁城 [17615007230@163.com] 2025-07-30 16:09:02
package server

import (
	"github.com/gin-gonic/gin"
	"net-project-edu_manage/handler"
)

// InitRouter 初始化路由
func initRouter() {

	//1.生成拦截器
	interceptor := interceptor()

	//2.定义version1版本路由组
	v1 := router.Group("api/v1")

	//3.定义v1-系统管理-用户管理路由组
	initSystemUserRouter(v1, interceptor)
}

// 初始化系统管理-用户管理路由
func initSystemUserRouter(v *gin.RouterGroup, interceptor gin.HandlerFunc) {

	//1.定义用户管理路由组
	systemUsers := v.Group("system-users")

	//2.使用拦截器
	systemUsers.Use(interceptor)

	//3.定义接口路由
	systemUsers.POST("", handler.AddSystemUser)
	systemUsers.DELETE("", nil)
	systemUsers.GET(":id", nil)
	systemUsers.PUT(":id", nil)
	systemUsers.GET("", nil)
}
