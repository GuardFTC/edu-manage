// Package router @Author:冯铁城 [17615007230@163.com] 2025-07-31 11:11:11
package router

import (
	"net-project-edu_manage/internal/http/handler"

	"github.com/gin-gonic/gin"
)

// 初始化系统管理-用户管理路由
func initAuthRouter(v *gin.RouterGroup) *gin.RouterGroup {

	//1.定义鉴权路由组
	authRouter := v.Group("auth")

	//2.定义接口路由
	authRouter.POST("/login", handler.Login)
	authRouter.POST("/refresh-token", handler.RefreshToken)

	//3.返回
	return authRouter
}
