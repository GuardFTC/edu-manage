// Package router @Author:冯铁城 [17615007230@163.com] 2025-08-04 16:18:00
package router

import (
	"github.com/gin-gonic/gin"
	"net-project-edu_manage/handler/handler"
)

// InitAcademicYearRouter 初始化学年路由
func initAcademicYearRouter(v *gin.RouterGroup, tokenHandler gin.HandlerFunc) *gin.RouterGroup {

	//1.定义用户管理路由组
	academicYearRouter := v.Group("academic-years")
	academicYearRouter.Use(tokenHandler)

	//2.定义接口路由
	academicYearRouter.POST("", handler.AddAcademicYear)
	academicYearRouter.DELETE("", nil)
	academicYearRouter.GET(":id", nil)
	academicYearRouter.PUT(":id", nil)
	academicYearRouter.GET("", nil)

	//3.返回
	return academicYearRouter
}
