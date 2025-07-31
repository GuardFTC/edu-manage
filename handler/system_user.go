// Package handler @Author:冯铁城 [17615007230@163.com] 2025-07-30 17:32:15
package handler

import (
	"github.com/gin-gonic/gin"
	"net-project-edu_manage/common/res"
	"net-project-edu_manage/model/dto"
	systemUser "net-project-edu_manage/service"
	"net-project-edu_manage/util"
)

// AddSystemUser 新增系统用户
func AddSystemUser(c *gin.Context) {

	//1.创建DTO
	var systemUserDTO *dto.SystemUserDTO

	//2.校验参数并绑定
	if err := c.ShouldBindJSON(&systemUserDTO); err != nil {
		util.FailResToC(c, res.BadRequestFail, util.FormatMsg(err.Error()))
		return
	}

	//3.保存系统用户
	if err := systemUser.Add(c, systemUserDTO); err != nil {
		util.FailResToC(c, res.ServerFail, err.Error())
		return
	}

	//4.返回
	util.SuccessResToC(c, res.CreateSuccess, systemUserDTO)
}
