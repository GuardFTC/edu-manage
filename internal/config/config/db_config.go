// Package config @Author:冯铁城 [17615007230@163.com] 2025-07-31 11:17:16
package config

import "time"

// DatabaseConfig 数据库配置结构体
type DatabaseConfig struct {

	//链接参数
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
	Config   string `mapstructure:"config"`

	//连接池参数
	MaxOpenConns    int           `mapstructure:"max_open_conns"`     // 最大连接数
	MaxIdleConns    int           `mapstructure:"max_idle_conns"`     // 最大空闲连接数
	ConnMaxLifetime time.Duration `mapstructure:"conn_max_lifetime"`  // 单连接最大生命周期
	ConnMaxIdleTime time.Duration `mapstructure:"conn_max_idle_time"` // 空闲连接最大存活时间
}
