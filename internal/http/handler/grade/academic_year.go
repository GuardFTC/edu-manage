// Package grade @Author:冯铁城 [17615007230@163.com] 2025-08-04 16:37:12
package grade

import (
	dtoPack "net-project-edu_manage/internal/model/dto/grade"
	reqPack "net-project-edu_manage/internal/model/request/grade"
	"net-project-edu_manage/internal/model/res"
	"net-project-edu_manage/internal/service/grade"

	"github.com/gin-gonic/gin"
)

// academicYearService 学年服务
var academicYearService = new(grade.AcademicYearService)

// AddAcademicYear 添加学年
func AddAcademicYear(c *gin.Context) {

	//1.创建DTO
	var academicYearDTO dtoPack.AcademicYearDto

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
	var academicYearDTO dtoPack.AcademicYearDto

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
	request := reqPack.AcademicYearRequest{}

	//2.校验URL参数并绑定
	if err := c.ShouldBindQuery(&request); err != nil {
		res.FailResToC(c, res.BadRequestFail, err.Error())
		return
	}

	//3.如果判定是执行列表查询，还是分页查询
	var resData any
	var err error
	if request.IsList {
		resData, err = academicYearService.List(c, &request)
	} else {
		resData, err = academicYearService.Page(c, &request)
	}

	//4.异常不为空，则返回异常信息
	if err != nil {
		res.FailResToCByMsg(c, err.Error())
		return
	}

	//5.返回
	res.SuccessResToC(c, res.QuerySuccess, resData)
}
