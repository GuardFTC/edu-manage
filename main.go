// @Author:冯铁城 [17615007230@163.com] 2025-07-30 14:46:08
package main

import (
	"log"
	"net-project-edu_manage/core/server"
	"net-project-edu_manage/dao"
)

func main() {

	//1.初始化DB
	dao.InitDB()

	//2.确保最终关闭数据库链接
	defer func() {
		db, _ := dao.DB.DB()
		_ = db.Close()
		log.Printf("database connection closed")
	}()

	//3.启动服务器
	server.StartServer()
}
