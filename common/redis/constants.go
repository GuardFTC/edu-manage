// Package redis @Author:冯铁城 [17615007230@163.com] 2025-08-05 15:40:52
package redis

import (
	"net-project-edu_manage/core/redis"
	"time"
)

var (
	DefaultExpire = time.Minute * 15 //默认过期时间
)

var (
	rdb = redis.Rdb //redis实例
	ctx = redis.Ctx //空白上下文
)
