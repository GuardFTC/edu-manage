// Package router @Author:冯铁城 [17615007230@163.com] 2025-07-31 11:11:11
package router

import (
	"github.com/gin-gonic/gin"
	"net-project-edu_manage/handler/handler"
)

// 初始化系统管理-用户管理路由
func initSystemUserRouter(v *gin.RouterGroup) *gin.RouterGroup {

	//1.定义用户管理路由组
	systemUserRouter := v.Group("system-users")

	//2.定义接口路由
	systemUserRouter.POST("", handler.AddSystemUser)
	systemUserRouter.DELETE("", nil)
	systemUserRouter.GET(":id", nil)
	systemUserRouter.PUT(":id", nil)
	systemUserRouter.GET("", nil)

	//3.返回
	return systemUserRouter
}
