// Package service @Author:冯铁城 [17615007230@163.com] 2025-08-01 19:13:53
package service

import (
	"errors"
	"net-project-edu_manage/common/res"
	"net-project-edu_manage/common/util"
	"net-project-edu_manage/core/db"
	"net-project-edu_manage/model/dto"
	"sync"
)

// AuthService 认证服务
type AuthService struct {
	sync sync.Mutex //预留锁 并发高时使用
}

// Login 登录
func (a *AuthService) Login(loginDto *dto.LoginDto) (string, error) {

	//1.根据账号查询用户信息
	systemUser, err := db.Q.SystemUser.
		Where(db.Q.SystemUser.Name.Eq(loginDto.Account)).
		Or(db.Q.SystemUser.Email.Eq(loginDto.Account)).
		First()
	if err != nil {
		return "", err
	}

	//2.TODO 密码解密 获取密码原文

	//3.比较密码
	if err = util.VerifyPassword(systemUser.Password, loginDto.Password); err != nil {
		return "", errors.New("password verify fail. " + res.UnProcessTag)
	}

	//4.登录成功，生成JWT Token，返回
	return util.GenerateJWT(systemUser.Name, systemUser.Email, 0)
}
