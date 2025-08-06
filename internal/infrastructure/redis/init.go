// Package redis @Author:冯铁城 [17615007230@163.com] 2025-08-05 15:08:55
package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"net-project-edu_manage/internal/config"
)

// Rdb redis客户端
var rdb *redis.Client

// Ctx 空白上下文
var ctx = context.Background()

// InitRedis 初始化redis
func InitRedis() {

	//1.获取redis配置
	host, port, password, database := getConfig()

	//2.初始化Redis客户端
	rdb = redis.NewClient(&redis.Options{
		Addr:     host + ":" + cast.ToString(port),
		Password: password,
		DB:       database,
	})

	//3.测试链接是否建立成功
	if err := rdb.Ping(ctx).Err(); err != nil {
		log.Fatalf("redis connection error: %v", err)
	} else {
		log.Printf("redis connection success")
	}
}

// CloseRedis 关闭Redis连接
func CloseRedis() {
	if err := rdb.Close(); err != nil {
		log.Printf("redis close error: %v", err)
	} else {
		log.Println("redis connection closed success")
	}
}

// getConfig 获取redis配置
func getConfig() (string, int, string, int) {

	//1.获取host，如果配置值为空，则使用默认值
	host := config.AppConfig.Redis.Host
	if host == "" {
		host = DefaultHost
	}

	//2.获取port，如果配置值为空，则使用默认值
	port := config.AppConfig.Redis.Port
	if port == 0 {
		port = DefaultPort
	}

	//3.获取password，如果配置值为空，则使用默认值
	password := config.AppConfig.Redis.Password
	if password == "" {
		password = DefaultPassword
	}

	//4.获取database，如果配置值为空，则使用默认值
	database := config.AppConfig.Redis.Database
	if database == 0 {
		database = DefaultDatabase
	}

	//5.返回
	return host, port, password, database
}
