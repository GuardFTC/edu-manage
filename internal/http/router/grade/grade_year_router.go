// Package grade @Author:冯铁城 [17615007230@163.com] 2025-08-08 11:10:34
package grade

import (
	"net-project-edu_manage/internal/http/handler/grade"

	"github.com/gin-gonic/gin"
)

// InitGradeYearRouter 初始化年级-学年管理路由
func InitGradeYearRouter(v *gin.RouterGroup, tokenHandler gin.HandlerFunc) *gin.RouterGroup {

	//1.定义用户管理路由组
	academicYearRouter := v.Group("grade-years")
	academicYearRouter.Use(tokenHandler)

	//2.定义接口路由
	academicYearRouter.POST("", grade.AddGradeYear)
	academicYearRouter.DELETE("", grade.DeleteGradeYear)
	academicYearRouter.PUT(":id", grade.UpdateGradeYear)
	academicYearRouter.GET("", grade.PageGradeYear)

	//3.返回
	return academicYearRouter
}
