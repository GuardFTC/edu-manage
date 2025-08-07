// Package redis @Author:冯铁城 [17615007230@163.com] 2025-08-05 15:08:55
package redis

import (
	con "net-project-edu_manage/internal/config"
	"net-project-edu_manage/internal/config/config"

	log "github.com/sirupsen/logrus"
)

// clients redis客户端Map
var clients map[string]*Client

// GetDefaultClient 获取默认数据源
func GetDefaultClient() *Client {
	return GetClient(con.AppConfig.Redis.Default)
}

// GetClient 获取数据源
func GetClient(rsName string) *Client {
	return clients[rsName]
}

// InitClient 初始化redis客户端
func InitClient(rsConfig *config.RedisSourceConfig) {

	//1.遍历Redis数据源配置
	for rsName, redisConfig := range rsConfig.Sources {

		//2.初始化Redis客户端
		rsClient, err := newClient(&redisConfig)
		if err != nil {
			log.Fatalf("redis-%v connection error: %v", rsName, err)
		} else {
			log.Printf("redis-%v connection success", rsName)
		}

		//3.存入map
		clients[rsName] = rsClient
	}
}

// CloseClient 关闭redis客户端
func CloseClient() {

	//1.遍历客户端集合
	for rsName, redisClient := range clients {

		//2.关闭数据源
		if err := redisClient.Close(); err != nil {
			log.Errorf("redis-%v connection closed error: %v", rsName, err)
		} else {
			log.Printf("redis-%v connection closed success", rsName)
		}
	}
}
