// Package handler @Author:冯铁城 [17615007230@163.com] 2025-08-04 16:37:12
package handler

import (
	"github.com/gin-gonic/gin"
	"net-project-edu_manage/common/res"
	"net-project-edu_manage/model/dto"
	"net-project-edu_manage/model/request"
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
		res.FailResToC(c, res.BadRequestFail, err.Error())
		return
	}

	//3.DTO解析时间
	if err := academicYearDTO.ParseDate(); err != nil {
		res.FailResToC(c, res.BadRequestFail, err.Error())
	}

	//4.起止时间校验
	if academicYearDTO.StartDate.After(academicYearDTO.EndDate) {
		res.FailResToC(c, res.BadRequestFail, "开始时间不能大于结束时间")
		return
	}

	//5.保存
	if err := academicYearService.Add(c, &academicYearDTO); err != nil {
		res.FailResToCByMsg(c, err.Error())
		return
	}

	//6.返回
	res.SuccessResToC(c, res.CreateSuccess, academicYearDTO)
}

// DeleteAcademicYear 删除学年
func DeleteAcademicYear(c *gin.Context) {

	//1.获取参数
	ids := c.QueryArray("id")
	if len(ids) == 0 {
		res.FailResToC(c, res.BadRequestFail, "参数为空")
		return
	}

	//2.删除学年
	if err := academicYearService.Delete(c, ids); err != nil {
		res.FailResToCByMsg(c, err.Error())
		return
	}

	//3.返回
	res.SuccessResToC(c, res.DeleteSuccess, nil)
}

// GetAcademicYear 查询单个学年
func GetAcademicYear(c *gin.Context) {

	//1.获取参数
	id := c.Param("id")

	//2.查询学年
	academicYearDto, err := academicYearService.Get(c, id)
	if err != nil {
		res.FailResToCByMsg(c, err.Error())
		return
	}

	//3.返回
	res.SuccessResToC(c, res.QuerySuccess, academicYearDto)
}

// UpdateAcademicYear 修改学年
func UpdateAcademicYear(c *gin.Context) {

	//1.获取路径参数
	id := c.Param("id")

	//2.创建DTO
	var academicYearDTO dto.AcademicYearDto

	//3.校验Body参数并绑定
	if err := c.ShouldBindJSON(&academicYearDTO); err != nil {
		res.FailResToC(c, res.BadRequestFail, err.Error())
		return
	}

	//4.DTO解析时间
	if err := academicYearDTO.ParseDate(); err != nil {
		res.FailResToC(c, res.BadRequestFail, err.Error())
	}

	//5.起止时间校验
	if academicYearDTO.StartDate.After(academicYearDTO.EndDate) {
		res.FailResToC(c, res.BadRequestFail, "开始时间不能大于结束时间")
		return
	}

	//6.更新学年
	if err := academicYearService.Update(c, id, &academicYearDTO); err != nil {
		res.FailResToCByMsg(c, err.Error())
		return
	}

	//7.返回
	res.SuccessResToC(c, res.UpdateSuccess, academicYearDTO)
}

// PageAcademicYear 分页查询学年
func PageAcademicYear(c *gin.Context) {

	//1.创建查询参数
	academicYearRequest := request.AcademicYearRequest{}

	//2.校验URL参数并绑定
	if err := c.ShouldBindQuery(&academicYearRequest); err != nil {
		res.FailResToC(c, res.BadRequestFail, err.Error())
		return
	}

	//3.分页查询
	pageRes, err := academicYearService.Page(c, &academicYearRequest)
	if err != nil {
		res.FailResToCByMsg(c, err.Error())
		return
	}

	//4.返回
	res.SuccessResToC(c, res.QuerySuccess, pageRes)
}
