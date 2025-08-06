// Package redis @Author:冯铁城 [17615007230@163.com] 2025-08-05 15:40:52
package redis

import (
	"time"
)

var (
	DefaultExpire   = time.Minute * 15 //默认过期时间
	DefaultHost     = "127.0.0.1"      //默认redis地址
	DefaultPort     = 6379             //默认redis端口
	DefaultPassword = ""               //默认redis密码
	DefaultDatabase = 0                //默认redis数据库
)
