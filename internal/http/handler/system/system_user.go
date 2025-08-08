// Package system @Author:冯铁城 [17615007230@163.com] 2025-07-30 17:32:15
package system

import (
	dtoPack "net-project-edu_manage/internal/model/dto/system"
	reqPack "net-project-edu_manage/internal/model/request/system"
	"net-project-edu_manage/internal/model/res"
	"net-project-edu_manage/internal/service/system"

	"github.com/gin-gonic/gin"
)

// systemUserService 系统用户服务
var systemUserService = new(system.SystemUserService)

// AddSystemUser 新增系统用户
func AddSystemUser(c *gin.Context) {

	//1.创建DTO
	var systemUserDto *dtoPack.SystemUserDto

	//2.校验参数并绑定
	if err := c.ShouldBindJSON(&systemUserDto); err != nil {
		res.FailResToC(c, res.BadRequestFail, err.Error())
		return
	}

	//3.保存系统用户
	if err := systemUserService.Add(c, systemUserDto); err != nil {
		res.FailResToCByMsg(c, err.Error())
		return
	}

	//4.返回
	res.SuccessResToC(c, res.CreateSuccess, systemUserDto)
}

// DeleteSystemUser 删除系统用户
func DeleteSystemUser(c *gin.Context) {

	//1.获取参数
	ids := c.QueryArray("id")
	if len(ids) == 0 {
		res.FailResToC(c, res.BadRequestFail, "参数为空")
		return
	}

	//2.删除系统用户
	if err := systemUserService.Delete(c, ids); err != nil {
		res.FailResToCByMsg(c, err.Error())
		return
	}

	//3.返回
	res.SuccessResToC(c, res.DeleteSuccess, nil)
}

// GetSystemUser 查询系统用户
func GetSystemUser(c *gin.Context) {

	//1.获取参数
	id := c.Param("id")

	//2.查询系统用户
	systemUserDto, err := systemUserService.Get(c, id)
	if err != nil {
		res.FailResToCByMsg(c, err.Error())
		return
	}

	//3.返回
	res.SuccessResToC(c, res.QuerySuccess, systemUserDto)
}

// UpdateSystemUser 修改系统用户
func UpdateSystemUser(c *gin.Context) {

	//1.获取路径参数
	id := c.Param("id")

	//2.创建DTO
	var systemUserDto *dtoPack.SystemUserDto

	//3.校验Body参数并绑定
	if err := c.ShouldBindJSON(&systemUserDto); err != nil {
		res.FailResToC(c, res.BadRequestFail, err.Error())
		return
	}

	//4.更新系统用户
	if err := systemUserService.Update(c, id, systemUserDto); err != nil {
		res.FailResToCByMsg(c, err.Error())
		return
	}

	//5.返回
	res.SuccessResToC(c, res.UpdateSuccess, systemUserDto)
}

// PageSystemUser 分页查询系统用户
func PageSystemUser(c *gin.Context) {

	//1.创建查询参数
	systemUserRequest := reqPack.SystemUserRequest{}

	//2.校验URL参数并绑定
	if err := c.ShouldBindQuery(&systemUserRequest); err != nil {
		res.FailResToC(c, res.BadRequestFail, err.Error())
		return
	}

	//3.分页查询
	pageRes, err := systemUserService.Page(c, &systemUserRequest)
	if err != nil {
		res.FailResToCByMsg(c, err.Error())
		return
	}

	//4.返回
	res.SuccessResToC(c, res.QuerySuccess, pageRes)
}
