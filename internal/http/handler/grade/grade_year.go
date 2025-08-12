// Package grade @Author:冯铁城 [17615007230@163.com] 2025-08-12 14:19:01
package grade

import (
	dtoPack "net-project-edu_manage/internal/model/dto/grade"
	reqPack "net-project-edu_manage/internal/model/request/grade"
	"net-project-edu_manage/internal/model/res"
	"net-project-edu_manage/internal/service/grade"

	"github.com/gin-gonic/gin"
)

// gradeYearService 年级-学年关联服务
var gradeYearService = new(grade.GradeYearService)

// AddGradeYear 添加年级-学年关联关系
func AddGradeYear(c *gin.Context) {

	//1.创建DTO
	var gradeYearDto dtoPack.GradeYearDto

	//2.校验参数并绑定
	if err := c.ShouldBindJSON(&gradeYearDto); err != nil {
		res.FailResToC(c, res.BadRequestFail, err.Error())
		return
	}

	//3.保存
	if err := gradeYearService.Add(c, &gradeYearDto); err != nil {
		res.FailResToCByMsg(c, err.Error())
		return
	}

	//4.返回
	res.SuccessResToC(c, res.CreateSuccess, gradeYearDto)
}

// DeleteGradeYear 删除年级-学年关联关系
func DeleteGradeYear(c *gin.Context) {

	//1.获取参数
	ids := c.QueryArray("id")
	if len(ids) == 0 {
		res.FailResToC(c, res.BadRequestFail, "参数为空")
		return
	}

	//2.删除
	if err := gradeYearService.Delete(c, ids); err != nil {
		res.FailResToCByMsg(c, err.Error())
		return
	}

	//3.返回
	res.SuccessResToC(c, res.DeleteSuccess, nil)
}

// UpdateGradeYear 更新年级-学年关联关系
func UpdateGradeYear(c *gin.Context) {

	//1.获取路径参数
	id := c.Param("id")

	//2.创建DTO
	var gradeYearDto dtoPack.GradeYearDto

	//3.校验Body参数并绑定
	if err := c.ShouldBindJSON(&gradeYearDto); err != nil {
		res.FailResToC(c, res.BadRequestFail, err.Error())
		return
	}

	//6.更新
	if err := gradeYearService.Update(c, id, &gradeYearDto); err != nil {
		res.FailResToCByMsg(c, err.Error())
		return
	}

	//7.返回
	res.SuccessResToC(c, res.UpdateSuccess, gradeYearDto)
}

// PageGradeYear 分页查询年级-学年关联关系
func PageGradeYear(c *gin.Context) {

	//1.创建查询参数
	request := reqPack.GradeYearRequest{}

	//2.校验URL参数并绑定
	if err := c.ShouldBindQuery(&request); err != nil {
		res.FailResToC(c, res.BadRequestFail, err.Error())
		return
	}

	//3.分页查询
	resData, err := gradeYearService.Page(c, &request)

	//4.异常不为空，则返回异常信息
	if err != nil {
		res.FailResToCByMsg(c, err.Error())
		return
	}

	//5.返回
	res.SuccessResToC(c, res.QuerySuccess, resData)
}
