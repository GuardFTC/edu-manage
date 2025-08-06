// Package router @Author:冯铁城 [17615007230@163.com] 2025-07-31 11:11:11
package router

import (
	"github.com/gin-gonic/gin"
	"net-project-edu_manage/internal/http/handler"
)

// 初始化系统管理-用户管理路由
func initSystemUserRouter(v *gin.RouterGroup, tokenHandler gin.HandlerFunc) *gin.RouterGroup {

	//1.定义用户管理路由组
	systemUserRouter := v.Group("system-users")
	systemUserRouter.Use(tokenHandler)

	//2.定义接口路由
	systemUserRouter.POST("", handler.AddSystemUser)
	systemUserRouter.DELETE("", handler.DeleteSystemUser)
	systemUserRouter.GET(":id", handler.GetSystemUser)
	systemUserRouter.PUT(":id", handler.UpdateSystemUser)
	systemUserRouter.GET("", handler.PageSystemUser)

	//3.返回
	return systemUserRouter
}
