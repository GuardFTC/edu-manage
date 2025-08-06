// Package config @Author:冯铁城 [17615007230@163.com] 2025-07-31 11:27:19
package config

// JwtConfig jwt配置
type JwtConfig struct {
	Key        string `mapstructure:"key"`
	ExpireHour int    `mapstructure:"expire_hour"`
}
