// Package config @Author:冯铁城 [17615007230@163.com] 2025-07-31 11:27:19
package config

// ServerConfig 服务配置结构体
type ServerConfig struct {
	Port string `mapstructure:"port"`
}
