// Package handler @Author:冯铁城 [17615007230@163.com] 2025-08-01 17:52:00
package handler

import (
	"github.com/gin-gonic/gin"
	"net-project-edu_manage/common/res"
	"net-project-edu_manage/common/util"
	"net-project-edu_manage/model/dto"
	"net-project-edu_manage/service"
)

var authService = new(service.AuthService)

// Login 登录
func Login(c *gin.Context) {

	//1.创建DTO
	var loginDto *dto.LoginDto

	//2.校验参数并绑定
	if err := c.ShouldBindJSON(&loginDto); err != nil {
		util.FailResToC(c, res.BadRequestFail, util.FormatMsg(err.Error()))
		return
	}

}
