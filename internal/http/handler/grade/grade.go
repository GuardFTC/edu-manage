// Package grade @Author:冯铁城 [17615007230@163.com] 2025-08-08 11:26:06
package grade

import (
	dtoPack "net-project-edu_manage/internal/model/dto/grade"
	reqPack "net-project-edu_manage/internal/model/request/grade"
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

	//3.保存
	if err := gradeService.Add(c, &gradeDto); err != nil {
		res.FailResToCByMsg(c, err.Error())
		return
	}

	//4.返回
	res.SuccessResToC(c, res.CreateSuccess, gradeDto)
}

// DeleteGrade 删除年级
func DeleteGrade(c *gin.Context) {

	//1.获取参数
	ids := c.QueryArray("id")
	if len(ids) == 0 {
		res.FailResToC(c, res.BadRequestFail, "参数为空")
		return
	}

	//2.删除年级
	if err := gradeService.Delete(c, ids); err != nil {
		res.FailResToCByMsg(c, err.Error())
		return
	}

	//3.返回
	res.SuccessResToC(c, res.DeleteSuccess, nil)
}

// GetGrade 获取年级
func GetGrade(c *gin.Context) {

	//1.获取路径参数年级ID
	id := c.Param("id")

	//2.查询年级
	gradeDto, err := gradeService.Get(c, id)
	if err != nil {
		res.FailResToCByMsg(c, err.Error())
	}

	//3.返回
	res.SuccessResToC(c, res.QuerySuccess, gradeDto)
}

// UpdateGrade 修改年级
func UpdateGrade(c *gin.Context) {

	//1.获取路径参数
	id := c.Param("id")

	//2.创建DTO
	var gradeDto dtoPack.GradeDto

	//3.校验Body参数并绑定
	if err := c.ShouldBindJSON(&gradeDto); err != nil {
		res.FailResToC(c, res.BadRequestFail, err.Error())
		return
	}

	//4.更新年级
	if err := gradeService.Update(c, id, &gradeDto); err != nil {
		res.FailResToCByMsg(c, err.Error())
		return
	}

	//5.返回
	res.SuccessResToC(c, res.UpdateSuccess, gradeDto)
}

func PageGrade(c *gin.Context) {

	//1.创建查询参数
	gradeRequest := reqPack.GradeRequest{}

	//2.校验URL参数并绑定
	if err := c.ShouldBindQuery(&gradeRequest); err != nil {
		res.FailResToC(c, res.BadRequestFail, err.Error())
		return
	}

	//3.分页查询
	pageRes, err := gradeService.Page(c, &gradeRequest)
	if err != nil {
		res.FailResToCByMsg(c, err.Error())
		return
	}

	//4.返回
	res.SuccessResToC(c, res.QuerySuccess, pageRes)
}
