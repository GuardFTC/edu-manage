// Package db @Author:冯铁城 [17615007230@163.com] 2023-09-14 09:22:28
package db

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net-project-edu_manage/dao/query"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 数据库连接参数
var (
	ip        = "127.0.0.1"
	port      = "3306"
	username  = "root"
	password  = "root"
	database  = "edu_test"
	dsnConfig = "charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai"
)

// DB 数据库连接
var DB *gorm.DB

// Q 数据库查询对象
var Q *query.Query

// InitDbConn 初始化数据库链接
func InitDbConn() {

	//1.拼接DSN
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?%v", username, password, ip, port, database, dsnConfig)

	//2.打开数据库连接
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("database connection error: %v", err)
	}

	//3.DB赋值
	DB = db

	//4.获取底层sql.DB
	sqlDb, err := DB.DB()
	if err != nil {
		log.Fatalf("get sql db connection error: %v", err)
	}

	//5.设置连接池参数
	sqlDb.SetMaxOpenConns(20)                  // 最多20个连接
	sqlDb.SetMaxIdleConns(10)                  // 最多10个空闲连接
	sqlDb.SetConnMaxLifetime(1 * time.Minute)  // 每个连接最多用1分钟
	sqlDb.SetConnMaxIdleTime(30 * time.Second) // 空闲超过30秒就关闭

	//6.日志打印
	log.Printf("database connection success")

	//7.初始化查询对象
	Q = query.Use(DB)

	//8.日志打印
	log.Printf("database query init success")
}

// CloseDbConn 关闭数据库连接
func CloseDbConn() {

	//1.获取底层sql.DB
	sqlDb, err := DB.DB()
	if err != nil {
		log.Fatalf("get sql db connection error: %v", err)
	}

	//2.关闭数据库链接
	if err = sqlDb.Close(); err == nil {
		log.Printf("database connection closed success")
	} else {
		log.Fatalf("database connection closed error")
	}
}
