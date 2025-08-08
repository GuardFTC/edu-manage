// Package auth @Author:冯铁城 [17615007230@163.com] 2025-08-01 19:15:33
package auth

// LoginDto 登录参数
type LoginDto struct {
	Account  string `json:"account" form:"account" binding:"required,max=128"`   //账号
	Password string `json:"password" form:"password" binding:"required,max=128"` //密码
}

// LoginResultDto 登录结果
type LoginResultDto struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
}
