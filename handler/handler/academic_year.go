// Package handler @Author:冯铁城 [17615007230@163.com] 2025-08-04 16:37:12
package handler

import (
	"github.com/gin-gonic/gin"
	"net-project-edu_manage/common/res"
	"net-project-edu_manage/common/util"
	"net-project-edu_manage/model/dto"
	"net-project-edu_manage/service"
)

// AcademicYearService 学年服务
var academicYearService = new(service.AcademicYearService)

// AddAcademicYear 添加学年
func AddAcademicYear(c *gin.Context) {

	//1.创建DTO
	var academicYearDTO dto.AcademicYearDto

	//2.校验参数并绑定
	if err := c.ShouldBindJSON(&academicYearDTO); err != nil {
		util.FailResToC(c, res.BadRequestFail, err.Error())
		return
	}

	//3.DTO解析时间
	if err := academicYearDTO.ParseDate(); err != nil {
		util.FailResToC(c, res.BadRequestFail, err.Error())
	}

	//4.起止时间校验
	if academicYearDTO.StartDate.After(academicYearDTO.EndDate) {
		util.FailResToC(c, res.BadRequestFail, "开始时间不能大于结束时间")
		return
	}

	//5.保存
	if err := academicYearService.Add(c, &academicYearDTO); err != nil {
		util.FailResToCByMsg(c, err.Error())
		return
	}

	//6.返回
	util.SuccessResToC(c, res.CreateSuccess, academicYearDTO)
}
