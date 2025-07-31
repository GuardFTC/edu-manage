// @Author:冯铁城 [17615007230@163.com] 2025-07-30 14:46:08
package main

import (
	"net-project-edu_manage/config"
	"net-project-edu_manage/core/db"
	"net-project-edu_manage/core/server"
)

func main() {

	//1.初始化日志
	config.InitLogger()

	//2.初始化DB
	db.InitDbConn()

	//3.确保最终关闭数据库链接
	defer db.CloseDbConn()

	//4.启动服务器
	server.StartServer()
}
