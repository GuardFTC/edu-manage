// Package core @Author:冯铁城 [17615007230@163.com] 2025-07-31 11:20:26
package core

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net-project-edu_manage/config/config"
)

// Config 配置结构体
type Config struct {
	Database config.DatabaseConfig `mapstructure:"database"`
	Server   config.ServerConfig   `mapstructure:"server"`
	Log      config.LogConfig      `mapstructure:"log"`
}

// AppConfig 项目总配置
var AppConfig Config

// InitConfig 初始化配置
func InitConfig() {

	//1.初始化viper
	v := viper.New()

	//2.从环境变量获取环境
	env := initEnv()

	//3.设置配置文件信息
	v.SetConfigName("application-" + env)
	v.SetConfigType("yaml")
	v.AddConfigPath("config/resources")

	//4.读取配置
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("read config error: %w", err))
	}

	//5.配置会写到AppConfig
	if err := v.Unmarshal(&AppConfig); err != nil {
		panic(fmt.Errorf("unmarshal config error: %w", err))
	}

	//6.初始化日志
	initLog()
}

// initEnv 初始化环境变量
func initEnv() string {

	//1.将配置项 EDU_ENV 绑定到环境变量
	if err := viper.BindEnv("EDU_ENV"); err != nil {
		log.Fatalf("bind env error: %v", err)
	}

	//2.读取EDU_ENV，默认为dev
	env := viper.GetString("EDU_ENV")
	if env == "" {
		env = "dev"
	}

	//3.返回
	return env
}

// initLog 日志初始化
func initLog() {

	//1.获取配置
	logConfig := AppConfig.Log

	//2.设置日志格式带颜色
	log.SetFormatter(&log.TextFormatter{
		ForceColors:   logConfig.Color,
		FullTimestamp: true,
	})

	//3.设置日志级别，默认 Info 及以上输出
	level, _ := log.ParseLevel(logConfig.Level)
	log.SetLevel(level)
}
