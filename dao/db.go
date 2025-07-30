// Package dao @Author:冯铁城 [17615007230@163.com] 2023-09-14 09:22:28
package dao

import (
	"fmt"
	"log"
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

// InitDB 初始化数据库
func InitDB() *gorm.DB {

	//1.拼接DSN
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?%v", username, password, ip, port, database, dsnConfig)

	//2.打开数据库连接
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("database connection error")
	}

	//3.获取底层sql.DB
	sqlDB, err := db.DB()
	if err != nil {
		panic("get sql db connection error")
	}

	//4.设置连接池参数
	sqlDB.SetMaxOpenConns(20)                  // 最多20个连接
	sqlDB.SetMaxIdleConns(10)                  // 最多10个空闲连接
	sqlDB.SetConnMaxLifetime(1 * time.Minute)  // 每个连接最多用1分钟
	sqlDB.SetConnMaxIdleTime(30 * time.Second) // 空闲超过30秒就关闭

	//5.返回数据库连接
	log.Printf("database connection success")
	return db
}
