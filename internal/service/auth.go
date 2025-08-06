// Package service @Author:冯铁城 [17615007230@163.com] 2025-08-01 19:13:53
package service

import (
	"errors"
	"net-project-edu_manage/internal/common/constant"
	"net-project-edu_manage/internal/common/util"
	"net-project-edu_manage/internal/config"
	"net-project-edu_manage/internal/infrastructure/db/model"
	"net-project-edu_manage/internal/infrastructure/redis"
	"net-project-edu_manage/internal/model/dto"
	"net-project-edu_manage/internal/model/res"
	"net-project-edu_manage/internal/repository"
	"sync"
	"time"

	"github.com/spf13/cast"
)

// systemUserRepo 系统用户仓库
var systemUserRepo = repository.NewSystemUserRepository()

// AuthService 认证服务
type AuthService struct {
	sync sync.Mutex //预留锁 并发高时使用
}

// Login 登录
func (a *AuthService) Login(loginDto *dto.LoginDto) (*dto.LoginResultDto, error) {

	//1.根据账号查询用户信息
	systemUser, err := systemUserRepo.GetByAccount(loginDto.Account)
	if err != nil {
		return nil, err
	}

	//2.TODO 密码解密 获取密码原文

	//3.比较密码
	if err = util.VerifyPassword(systemUser.Password, loginDto.Password); err != nil {
		return nil, errors.New("password verify fail. " + res.UnProcessTag)
	}

	//4.获取token
	token, err := getToken(systemUser)
	if err != nil {
		return nil, err
	}

	//5.获取refreshToken
	refreshToken, err := getRefreshToken(systemUser)
	if err != nil {
		return nil, err
	}

	//6.封装结构体
	loginRes := &dto.LoginResultDto{
		Token:        token,
		RefreshToken: refreshToken,
	}

	//7.返回
	return loginRes, nil
}

// RefreshToken 刷新token
func (a *AuthService) RefreshToken(token string) (*dto.LoginResultDto, error) {
	return nil, nil
}

// getToken 获取token
func getToken(systemUser *model.SystemUser) (string, error) {

	//1.获取token过期时间
	exp := time.Duration(config.AppConfig.Jwt.ExpireHour) * time.Hour

	//2.获取token
	token, _ := redis.HashClient.HGet(constant.LoginTokenKey, cast.ToString(systemUser.ID))

	//3.token为空，生成token
	if token == "" {

		//4.生成token
		t, err := util.GenerateJWT(systemUser.ID, systemUser.Name, systemUser.Email, exp, false)
		if err != nil {
			return "", err
		}
		token = t

		//5.异步写入redis
		go func() {

			//6.记录用户ID-Token
			redis.HashClient.HSet(constant.LoginTokenKey, cast.ToString(systemUser.ID), token)
			redis.StringClient.Expire(constant.LoginTokenKey, exp)

			//7.记录Token-用户ID
			redis.HashClient.HSet(constant.LoginTokenMapKey, token, cast.ToString(systemUser.ID))
			redis.StringClient.Expire(constant.LoginTokenMapKey, exp)
		}()
	}

	//6.返回
	return token, nil
}

// getRefreshToken 获取refreshToken
func getRefreshToken(systemUser *model.SystemUser) (string, error) {

	//1.获取refreshToken过期时间
	exp := time.Duration(config.AppConfig.Jwt.RefreshExpireHour) * time.Hour

	//2.获取refreshToken
	refreshToken, _ := redis.HashClient.HGet(constant.LoginRefreshTokenKey, cast.ToString(systemUser.ID))

	//3.refreshToken为空，生成refreshToken
	if refreshToken == "" {

		//4.生成refreshToken
		t, err := util.GenerateJWT(systemUser.ID, systemUser.Name, systemUser.Email, exp, true)
		if err != nil {
			return "", err
		}
		refreshToken = t

		//5.异步写入redis
		go func() {
			redis.HashClient.HSet(constant.LoginRefreshTokenKey, cast.ToString(systemUser.ID), refreshToken)
			redis.StringClient.Expire(constant.LoginRefreshTokenKey, exp)
		}()
	}

	//6.返回
	return refreshToken, nil
}
