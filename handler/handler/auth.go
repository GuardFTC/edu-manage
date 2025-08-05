// Package handler @Author:冯铁城 [17615007230@163.com] 2025-08-01 17:52:00
package handler

import (
	"github.com/gin-gonic/gin"
	"net-project-edu_manage/common/res"
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
		res.FailResToC(c, res.BadRequestFail, err.Error())
		return
	}

	//3.登录
	jwtToken, err := authService.Login(loginDto)
	if err != nil {
		res.FailResToCByMsg(c, err.Error())
		return
	}

	//4.返回
	res.SuccessResToC(c, res.QuerySuccess, jwtToken)
}
