// Package router @Author:冯铁城 [17615007230@163.com] 2025-07-30 16:09:02
package router

import (
	"net-project-edu_manage/internal/http/interceptor"
	"net-project-edu_manage/internal/http/router/auth"
	"net-project-edu_manage/internal/http/router/grade"
	"net-project-edu_manage/internal/http/router/system"
	"net-project-edu_manage/internal/http/validate"

	"github.com/gin-gonic/gin"
)

// Router 核心路由
var Router *gin.Engine

// InitRouter 初始化路由
func InitRouter() {

	//1.初始化自定义校验器
	validate.InitValidate()

	//2.初始化基础路由组
	initBaseRouter()

	//3.初始化业务模块路由组
	initModelRouter()
}

// initBaseRouter 初始化核心基础路由
func initBaseRouter() {

	//1.创建 Gin 实例
	Router = gin.New()

	//2.使用全局异常处理器
	Router.Use(interceptor.GetErrorHandler())

	//3.使用 Logger 和 Recovery 中间件
	Router.Use(gin.Logger(), gin.Recovery())
}

// initModelRouter 初始化业务模块路由
func initModelRouter() {

	//1.获取Token处理器
	tokenHandler := interceptor.GetTokenHandler()

	//2.定义version1版本路由组
	v1 := Router.Group("api/v1")

	//3.定义v1-鉴权路由组
	auth.InitAuthRouter(v1)

	//4.定义v1-系统管理-用户管理路由组
	system.InitSystemUserRouter(v1, tokenHandler)

	//5.定义v1-年级管理-学年路由组
	grade.InitAcademicYearRouter(v1, tokenHandler)
}
