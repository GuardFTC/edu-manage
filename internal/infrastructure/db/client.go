// Package db @Author:冯铁城 [17615007230@163.com] 2025-08-07 15:25:18
package db

import (
	"database/sql"
	"fmt"
	"net-project-edu_manage/internal/config/config"
	"net-project-edu_manage/internal/infrastructure/db/query"
	"time"

	"github.com/spf13/cast"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Client 数据库客户端
type Client struct {
	db    *gorm.DB
	sqlDb *sql.DB
	q     *query.Query
}

// NewClient 创建一个新的数据库客户端实例
func newClient(dbConfig *config.DatabaseConfig) (*Client, error) {

	//1.获取主数据源DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", dbConfig.Username, dbConfig.Password, dbConfig.Host, cast.ToInt(dbConfig.Port), dbConfig.DBName, dbConfig.Config)

	//2.打开数据库连接
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	//3.获取底层sql.DB
	sqlDb, err := db.DB()
	if err != nil {
		return nil, err
	}

	//4.设置连接池参数
	sqlDb.SetMaxOpenConns(dbConfig.MaxOpenConns)                     // 最多20个连接
	sqlDb.SetMaxIdleConns(dbConfig.MaxIdleConns)                     // 最多10个空闲连接
	sqlDb.SetConnMaxLifetime(dbConfig.ConnMaxLifetime * time.Minute) // 每个连接最多用1分钟
	sqlDb.SetConnMaxIdleTime(dbConfig.ConnMaxIdleTime * time.Second) // 空闲超过30秒就关闭

	//5.测试连接
	if err = sqlDb.Ping(); err != nil {
		return nil, err
	}

	//6.生成gen查询对象
	q := query.Use(db)

	//7.创建客户端
	dbClient := &Client{
		db:    db,
		q:     q,
		sqlDb: sqlDb,
	}

	//8.返回
	return dbClient, nil
}

// Close 关闭数据库连接
func (c *Client) Close() error {
	if c.sqlDb != nil {
		return c.sqlDb.Close()
	}
	return nil
}

// Ping 测试数据库连接
func (c *Client) Ping() error {
	return c.sqlDb.Ping()
}

// GetQuery 获取查询对象
func (c *Client) GetQuery() *query.Query {
	return c.q
}

// GetDB 获取数据库连接
func (c *Client) GetDB() *gorm.DB {
	return c.db
}

// GetRawClient 获取原生数据库连接
func (c *Client) GetRawClient() *sql.DB {
	return c.sqlDb
}
