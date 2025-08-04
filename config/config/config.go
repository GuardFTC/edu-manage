// Package core @Author:冯铁城 [17615007230@163.com] 2025-07-31 11:20:26
package config

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"strings"
)

// Config 配置结构体
type Config struct {
	Database DatabaseConfig `mapstructure:"database"`
	Server   ServerConfig   `mapstructure:"server"`
	Log      LogConfig      `mapstructure:"log"`
	Jwt      JwtConfig      `mapstructure:"jwt"`
}

// AppConfig 项目总配置
var AppConfig Config

// InitConfig 初始化配置
func InitConfig() {

	//1.初始化viper
	v := viper.New()

	//2.从环境变量获取环境
	env := initEnv()

	//3.获取项目根目录
	basePath := getBasePath(env)

	//4.获取配置文件目录
	configPath := filepath.Join(basePath, "config", "resources")

	//5.设置配置文件信息
	v.AddConfigPath(configPath)
	v.SetConfigName("application-" + env)
	v.SetConfigType("yaml")

	//6.读取配置
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("read config error: %w", err))
	}

	//7.配置会写到AppConfig
	if err := v.Unmarshal(&AppConfig); err != nil {
		panic(fmt.Errorf("unmarshal config error: %w", err))
	}

	//8.初始化日志
	initLog()

	//9.打印配置文件读取日志
	log.Printf("init config by application-%s.yml", env)
}

// InitUnitTestConfig 初始化单元测试配置
func InitUnitTestConfig() {

	//1.初始化viper
	v := viper.New()

	//2.从环境变量获取环境
	env := initEnv()

	//3.获取项目根目录
	basePath := getBasePath(env)
	basePath = strings.Replace(basePath, "\\common\\util", "", -1)

	//4.获取配置文件目录
	configPath := filepath.Join(basePath, "config", "resources")

	//5.设置配置文件信息
	v.AddConfigPath(configPath)
	v.SetConfigName("application-" + env)
	v.SetConfigType("yaml")

	//6.读取配置
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("read config error: %w", err))
	}

	//7.配置会写到AppConfig
	if err := v.Unmarshal(&AppConfig); err != nil {
		panic(fmt.Errorf("unmarshal config error: %w", err))
	}

	//8.初始化日志
	initLog()

	//9.打印配置文件读取日志
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

// getBasePath 获取项目根目录
func getBasePath(env string) string {

	//1.开发环境使用当前工作目录
	if env == "dev" {
		wd, err := os.Getwd()
		if err != nil {
			panic("无法获取当前工作目录: " + err.Error())
		}
		return wd
	}

	//2.生产环境使用可执行文件所在目录
	execPath, err := os.Executable()
	if err != nil {
		panic("无法获取可执行文件路径: " + err.Error())
	}
	return filepath.Dir(execPath)
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
