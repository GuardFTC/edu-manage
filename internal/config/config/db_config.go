// Package config @Author:冯铁城 [17615007230@163.com] 2025-07-31 11:17:16
package config

import "time"

// DataBaseSourceConfig 数据库数据源配置结构体
type DataBaseSourceConfig struct {
	Master DatabaseConfig `mapstructure:"master"`
	Slave1 DatabaseConfig `mapstructure:"slave1"`
}

// DatabaseConfig 数据库配置结构体
type DatabaseConfig struct {
	Host            string        `mapstructure:"host"`
	Port            string        `mapstructure:"port"`
	Username        string        `mapstructure:"username"`
	Password        string        `mapstructure:"password"`
	DBName          string        `mapstructure:"dbname"`
	Config          string        `mapstructure:"config"`
	MaxIdleConns    int           `mapstructure:"max_idle_conns"`
	MaxOpenConns    int           `mapstructure:"max_open_conns"`
	ConnMaxLifetime time.Duration `mapstructure:"conn_max_lifetime"`
	ConnMaxIdleTime time.Duration `mapstructure:"conn_max_idle_time"`
}
