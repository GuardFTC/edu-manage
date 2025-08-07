// Package db @Author:冯铁城 [17615007230@163.com] 2023-09-14 09:22:28
package db

import (
	"net-project-edu_manage/internal/config/config"

	log "github.com/sirupsen/logrus"
)

var (
	Master *client
	Slave1 *client
)

// InitDbConn 初始化数据库链接
func InitDbConn(dsConfig *config.DataBaseSourceConfig) {

	//1.初始化主数据源
	master, err := newClient(&dsConfig.Master)
	if err != nil {
		log.Fatalf("database-%v connection error: %v", "master", err)
	} else {
		log.Printf("database-%v connection success", "master")
	}
	Master = master

	//2.初始化从数据源
	slave1, err := newClient(&dsConfig.Slave1)
	if err != nil {
		log.Fatalf("database-%v connection error: %v", "slave1", err)
	} else {
		log.Printf("database-%v connection success", "slave1")
	}
	Slave1 = slave1
}

// CloseDbConn 关闭数据库连接
func CloseDbConn() {

	//1.关闭主数据源
	if err := Master.Close(); err != nil {
		log.Errorf("database-%v connection closed error: %v", "master", err)
	} else {
		log.Printf("database-%v connection closed", "master")
	}

	//2.关闭从数据源
	if err := Slave1.Close(); err == nil {
		log.Errorf("database-%v connection closed error: %v", "slave1", err)
	} else {
		log.Printf("database-%v connection closed", "slave1")
	}
}
