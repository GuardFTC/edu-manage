// Package config @Author:冯铁城 [17615007230@163.com] 2025-08-05 20:02:26
package config

// RedisConfig Redis配置
type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	Database int    `mapstructure:"database"`
}
