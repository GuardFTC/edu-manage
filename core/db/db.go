// Package db @Author:冯铁城 [17615007230@163.com] 2023-09-14 09:22:28
package db

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net-project-edu_manage/core"
	"net-project-edu_manage/dao/query"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB 数据库连接
var DB *gorm.DB

// Q 数据库查询对象
var Q *query.Query

// InitDbConn 初始化数据库链接
func InitDbConn() {

	//1.获取DSN
	dsn := getDsn()

	//2.打开数据库连接
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("database connection error: %v", err)
	}

	//3.DB赋值
	DB = db

	//4.设置连接池参数
	setConnPool()

	//5.日志打印
	log.Printf("database connection success")

	//6.初始化查询对象
	Q = query.Use(DB)

	//7.日志打印
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

// getDsn 获取DSN
func getDsn() string {

	//1.读取值
	username := core.AppConfig.Database.Username
	password := core.AppConfig.Database.Password
	ip := core.AppConfig.Database.Host
	dbName := core.AppConfig.Database.DBName
	port := core.AppConfig.Database.Port
	dsnConfig := core.AppConfig.Database.Config

	//2.拼接DSN，返回
	return fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?%v", username, password, ip, port, dbName, dsnConfig)
}

// setConnPool 设置连接池参数
func setConnPool() {

	//1.获取底层sql.DB
	sqlDb, err := DB.DB()
	if err != nil {
		log.Fatalf("get sql db connection error: %v", err)
	}

	//2.设置连接池参数
	sqlDb.SetMaxOpenConns(core.AppConfig.Database.MaxOpenConns)                     // 最多20个连接
	sqlDb.SetMaxIdleConns(core.AppConfig.Database.MaxIdleConns)                     // 最多10个空闲连接
	sqlDb.SetConnMaxLifetime(core.AppConfig.Database.ConnMaxLifetime * time.Minute) // 每个连接最多用1分钟
	sqlDb.SetConnMaxIdleTime(core.AppConfig.Database.ConnMaxIdleTime * time.Second) // 空闲超过30秒就关闭
}
