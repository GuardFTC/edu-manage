// Package auth @Author:冯铁城 [17615007230@163.com] 2025-08-01 17:52:00
package auth

import (
	"net-project-edu_manage/internal/common/constant"
	"net-project-edu_manage/internal/common/util"
	"net-project-edu_manage/internal/infrastructure/redis"
	dtoPack "net-project-edu_manage/internal/model/dto/auth"
	"net-project-edu_manage/internal/model/res"
	"net-project-edu_manage/internal/service/auth"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

var authService = new(auth.AuthService)

// Login 登录
func Login(c *gin.Context) {

	//1.创建DTO
	var loginDto *dtoPack.LoginDto

	//2.校验参数并绑定
	if err := c.ShouldBindJSON(&loginDto); err != nil {
		res.FailResToC(c, res.BadRequestFail, err.Error())
		return
	}

	//3.登录
	loginRes, err := authService.Login(c, loginDto)
	if err != nil {
		res.FailResToCByMsg(c, err.Error())
		return
	}

	//4.返回
	res.SuccessResToC(c, res.QuerySuccess, loginRes)
}

// RefreshToken 刷新token
func RefreshToken(c *gin.Context) {

	//1.获取refreshToken
	refreshToken := c.Query("refreshToken")

	//2.判空
	claims, err := util.ParseJWT(refreshToken)
	if err != nil {
		res.FailResToC(c, res.BadRequestFail, "refreshToken is invalid")
		return
	}

	//3. 判定redis是否包含当前token,
	userId := cast.ToString(claims["id"])
	refreshTokenInRedis, err := redis.GetDefaultClient().Hash.HGet(c, constant.LoginRefreshTokenKey, userId)
	if err != nil || refreshTokenInRedis != refreshToken {
		res.FailResToC(c, res.BadRequestFail, "refreshToken is invalid")
		return
	}

	//4.刷新token
	refreshRes, err := authService.RefreshToken(c, refreshToken, claims)
	if err != nil {
		res.FailResToCByMsg(c, err.Error())
		return
	}

	//5.返回
	res.SuccessResToC(c, res.QuerySuccess, refreshRes)
}
