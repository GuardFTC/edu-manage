// Package config @Author:冯铁城 [17615007230@163.com] 2025-08-04 11:24:08
package config

import (
	"bytes"
	"fmt"
	"net-project-edu_manage/internal/config/config"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Config 配置结构体
type Config struct {
	DataBaseSource config.DataBaseSourceConfig `mapstructure:"database"`
	Server         config.ServerConfig         `mapstructure:"server"`
	Log            config.LogConfig            `mapstructure:"log"`
	Jwt            config.JwtConfig            `mapstructure:"jwt"`
	Redis          config.RedisConfig          `mapstructure:"redis"`
}

// AppConfig 项目总配置
var AppConfig Config

// InitConfig 从 embed.FS 读取配置，并映射到结构体
func InitConfig() {

	//1.从环境变量获取环境
	env := initEnv()

	//2.通过 embedded 读取配置文件内容
	fileName := fmt.Sprintf("resources/application-%s.yaml", env)
	data, err := ConfigFiles.ReadFile(fileName)
	if err != nil {
		panic(fmt.Errorf("read embedded config error: %w", err))
	}

	//3.初始化viper,设置配置文件信息
	v := viper.New()
	v.SetConfigType("yaml")

	//4.读取配置
	if err = v.ReadConfig(bytes.NewBuffer(data)); err != nil {
		panic(fmt.Errorf("read config error: %w", err))
	}

	//5.配置回写到AppConfig
	if err = v.Unmarshal(&AppConfig); err != nil {
		panic(fmt.Errorf("unmarshal config error: %w", err))
	}

	//6.初始化日志
	initLog()

	//7.打印配置文件读取日志
	log.Printf("init config by application-%s.yml", env)
}

// initEnv 初始化环境变量
func initEnv() string {

	//1.启用自动绑定
	viper.AutomaticEnv()

	//2.将配置项 EDU_ENV 绑定到环境变量
	if err := viper.BindEnv("EDU_ENV"); err != nil {
		log.Fatalf("bind env error: %v", err)
	}

	//3.读取EDU_ENV，默认为dev
	env := viper.GetString("EDU_ENV")
	if env == "" {
		env = "dev"
	}

	//4.返回
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
