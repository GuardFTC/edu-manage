// @Author:冯铁城 [17615007230@163.com] 2025-07-30 14:46:08
package main

import (
	"net-project-edu_manage/internal/config"
	"net-project-edu_manage/internal/http/router"
	"net-project-edu_manage/internal/infrastructure/db"
	"net-project-edu_manage/internal/infrastructure/redis"
	"net-project-edu_manage/internal/server"
)

func main() {

	//1.初始化配置（使用嵌入式配置文件）
	config.InitConfig()

	//2.初始化DB,确保最终关闭数据库链接
	db.InitDbConn()
	defer db.CloseDbConn()

	//3.初始化Redis,确保最终关闭Redis链接
	redis.InitClient(&config.AppConfig.Redis)
	defer redis.CloseClient()

	//4.初始化路由
	router.InitRouter()

	//5.启动服务器
	server.StartServer()
}
