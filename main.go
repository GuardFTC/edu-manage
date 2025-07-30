// @Author:冯铁城 [17615007230@163.com] 2025-07-30 14:46:08
package main

import (
	log "github.com/sirupsen/logrus"
	"net-project-edu_manage/core/_log"
	"net-project-edu_manage/core/server"
	"net-project-edu_manage/dao"
)

func main() {

	//1.初始化日志
	core.InitLogger()

	//2.初始化DB
	dao.InitDB()

	//3.确保最终关闭数据库链接
	defer func() {
		db, _ := dao.DB.DB()
		_ = db.Close()
		log.Printf("database connection closed")
	}()

	//4.初始化Query
	dao.InitQuery()

	//5.启动服务器
	server.StartServer()
}
