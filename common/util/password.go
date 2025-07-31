// @Author:冯铁城 [17615007230@163.com] 2025-07-30 14:50:05
package util

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword 获取密码
func HashPassword(password string) (string, error) {

	//1.判空处理
	if password == "" {
		return "", errors.New("密码不能为空")
	}

	//2.加密处理
	pwdHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	//3.转为字符串,返回
	return string(pwdHash), err
}

// VerifyPassword 验证密码
func VerifyPassword(dbPassword string, password string) error {

	//1.判空处理
	if password == "" || dbPassword == "" {
		return errors.New("登录密码为空或数据库无法获取到密码")
	}

	//2.校验,返回
	return bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(password))
}
