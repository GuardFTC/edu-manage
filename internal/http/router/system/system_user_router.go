// Package system @Author:冯铁城 [17615007230@163.com] 2025-07-31 11:11:11
package system

import (
	"net-project-edu_manage/internal/http/handler/system"

	"github.com/gin-gonic/gin"
)

// InitSystemUserRouter 初始化系统管理-用户管理路由
func InitSystemUserRouter(v *gin.RouterGroup, tokenHandler gin.HandlerFunc) *gin.RouterGroup {

	//1.定义用户管理路由组
	systemUserRouter := v.Group("system-users")
	systemUserRouter.Use(tokenHandler)

	//2.定义接口路由
	systemUserRouter.POST("", system.AddSystemUser)
	systemUserRouter.DELETE("", system.DeleteSystemUser)
	systemUserRouter.GET(":id", system.GetSystemUser)
	systemUserRouter.PUT(":id", system.UpdateSystemUser)
	systemUserRouter.GET("", system.PageSystemUser)

	//3.返回
	return systemUserRouter
}
