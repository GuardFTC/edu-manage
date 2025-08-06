// Package config @Author:冯铁城 [17615007230@163.com] 2025-07-30 17:27:03
package config

// LogConfig 日志配置
type LogConfig struct {
	Level string `mapstructure:"level"`
	Color bool   `mapstructure:"color"`
}
