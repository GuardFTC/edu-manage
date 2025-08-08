// Package grade @Author:冯铁城 [17615007230@163.com] 2025-08-04 16:18:00
package grade

import (
	"net-project-edu_manage/internal/http/handler/grade"

	"github.com/gin-gonic/gin"
)

// InitAcademicYearRouter 初始化学年路由
func InitAcademicYearRouter(v *gin.RouterGroup, tokenHandler gin.HandlerFunc) *gin.RouterGroup {

	//1.定义用户管理路由组
	academicYearRouter := v.Group("academic-years")
	academicYearRouter.Use(tokenHandler)

	//2.定义接口路由
	academicYearRouter.POST("", grade.AddAcademicYear)
	academicYearRouter.DELETE("", grade.DeleteAcademicYear)
	academicYearRouter.GET(":id", grade.GetAcademicYear)
	academicYearRouter.PUT(":id", grade.UpdateAcademicYear)
	academicYearRouter.GET("", grade.PageAcademicYear)
	academicYearRouter.GET(":id/grades", grade.GetYearGrade)

	//3.返回
	return academicYearRouter
}
