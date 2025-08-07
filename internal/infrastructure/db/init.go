// Package db @Author:冯铁城 [17615007230@163.com] 2023-09-14 09:22:28
package db

import (
	con "net-project-edu_manage/internal/config"
	"net-project-edu_manage/internal/config/config"

	log "github.com/sirupsen/logrus"
)

// Client 数据库客户端
var clients map[string]*Client

// GetDefaultDataSource 获取默认数据源
func GetDefaultDataSource() *Client {
	return GetDataSource(con.AppConfig.DataBaseSource.Default)
}

// GetDataSource 获取数据源
func GetDataSource(dsName string) *Client {
	return clients[dsName]
}

// InitDbConn 初始化数据库链接
func InitDbConn(dsConfig *config.DataBaseSourceConfig) {

	//1.初始化map
	clients = make(map[string]*Client, len(dsConfig.Sources))

	//2.遍历数据源
	for dsName, dbConfig := range dsConfig.Sources {

		//3.获取客户端
		dbClient, err := newClient(&dbConfig)

		//4.异常不为空打印异常，否则存入Map
		if err != nil {
			log.Fatalf("database-%v connection error: %v", dsName, err)
		} else {
			clients[dsName] = dbClient
			log.Printf("database-%v connection success", dsName)
		}
	}
}

// CloseDbConn 关闭数据库连接
func CloseDbConn() {

	//1.遍历数据源
	for dsName, dbClient := range clients {

		//2.循环关闭数据源
		if err := dbClient.Close(); err != nil {
			log.Errorf("database-%v connection closed error: %v", dsName, err)
		} else {
			log.Printf("database-%v connection closed", dsName)
		}
	}
}
