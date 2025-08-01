// Package handler @Author:冯铁城 [17615007230@163.com] 2025-07-30 17:32:15
package handler

import (
	"github.com/gin-gonic/gin"
	"net-project-edu_manage/common/res"
	"net-project-edu_manage/common/util"
	"net-project-edu_manage/model/dto"
	"net-project-edu_manage/model/request"
	systemUser "net-project-edu_manage/service"
)

// AddSystemUser 新增系统用户
func AddSystemUser(c *gin.Context) {

	//1.创建DTO
	var systemUserDto *dto.SystemUserDto

	//2.校验参数并绑定
	if err := c.ShouldBindJSON(&systemUserDto); err != nil {
		util.FailResToC(c, res.BadRequestFail, util.FormatMsg(err.Error()))
		return
	}

	//3.保存系统用户
	if err := systemUser.Add(c, systemUserDto); err != nil {
		util.FailResToCByMsg(c, err.Error())
		return
	}

	//4.返回
	util.SuccessResToC(c, res.CreateSuccess, systemUserDto)
}

// DeleteSystemUser 删除系统用户
func DeleteSystemUser(c *gin.Context) {

	//1.获取参数
	ids := c.QueryArray("id")
	if len(ids) == 0 {
		util.FailResToC(c, res.BadRequestFail, "参数为空")
		return
	}

	//2.删除系统用户
	if err := systemUser.Delete(c, ids); err != nil {
		util.FailResToCByMsg(c, err.Error())
		return
	}

	//3.返回
	util.SuccessResToC(c, res.DeleteSuccess, nil)
}

// GetSystemUser 查询系统用户
func GetSystemUser(c *gin.Context) {

	//1.获取参数
	id := c.Param("id")

	//2.查询系统用户
	systemUserDto, err := systemUser.Get(c, id)
	if err != nil {
		util.FailResToCByMsg(c, err.Error())
		return
	}

	//3.返回
	util.SuccessResToC(c, res.QuerySuccess, systemUserDto)
}

// UpdateSystemUser 修改系统用户
func UpdateSystemUser(c *gin.Context) {

	//1.获取路径参数
	id := c.Param("id")

	//2.创建DTO
	var systemUserDto *dto.SystemUserDto

	//3.校验Body参数并绑定
	if err := c.ShouldBindJSON(&systemUserDto); err != nil {
		util.FailResToC(c, res.BadRequestFail, util.FormatMsg(err.Error()))
		return
	}

	//4.更新系统用户
	if err := systemUser.Update(c, id, systemUserDto); err != nil {
		util.FailResToCByMsg(c, err.Error())
		return
	}

	//5.返回
	util.SuccessResToC(c, res.UpdateSuccess, systemUserDto)
}

// PageSystemUser 分页查询系统用户
func PageSystemUser(c *gin.Context) {

	//1.创建查询参数
	systemUserRequest := request.SystemUserRequest{}

	//2.校验URL参数并绑定
	if err := c.ShouldBindQuery(&systemUserRequest); err != nil {
		util.FailResToC(c, res.BadRequestFail, util.FormatMsg(err.Error()))
		return
	}

	//3.分页查询
	pageRes, err := systemUser.Page(c, &systemUserRequest)
	if err != nil {
		util.FailResToCByMsg(c, err.Error())
		return
	}

	//4.返回
	util.SuccessResToC(c, res.QuerySuccess, pageRes)
}
