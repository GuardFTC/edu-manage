// @Author:冯铁城 [17615007230@163.com] 2025-07-30 14:46:08
package main

import (
	"net-project-edu_manage/config"
	"net-project-edu_manage/core/db"
	"net-project-edu_manage/core/redis"
	"net-project-edu_manage/core/server"
)

func main() {

	//1.初始化配置（使用嵌入式配置文件）
	config.InitConfig()

	//2.初始化DB
	db.InitDbConn()

	//3.确保最终关闭数据库链接
	defer db.CloseDbConn()

	//4.初始化Redis
	redis.InitRedis()

	//5.确保最终关闭Redis链接
	defer redis.CloseRedis()

	//6.启动服务器
	server.StartServer()
}
