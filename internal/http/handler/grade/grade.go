// Package grade @Author:冯铁城 [17615007230@163.com] 2025-08-08 11:26:06
package grade

import (
	dtoPack "net-project-edu_manage/internal/model/dto/grade"
	"net-project-edu_manage/internal/model/res"
	"net-project-edu_manage/internal/service/grade"

	"github.com/gin-gonic/gin"
)

// gradeService 年级服务
var gradeService = new(grade.GradeService)

// AddGrade 添加年级
func AddGrade(c *gin.Context) {

	//1.创建DTO
	var gradeDto dtoPack.GradeDto

	//2.校验参数并绑定
	if err := c.ShouldBindJSON(&gradeDto); err != nil {
		res.FailResToC(c, res.BadRequestFail, err.Error())
		return
	}

	//3.获取路径参数学年ID
	yearId := c.Param("id")

	//3.保存
	if err := gradeService.Add(c, yearId, &gradeDto); err != nil {
		res.FailResToCByMsg(c, err.Error())
		return
	}

	//4.返回
	res.SuccessResToC(c, res.CreateSuccess, gradeDto)
}
