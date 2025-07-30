// Package core @Author:冯铁城 [17615007230@163.com] 2025-07-30 16:20:53
package server

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// interceptor 拦截器
func interceptor() gin.HandlerFunc {
	return func(c *gin.Context) {

		//1.获取请求头
		token := c.GetHeader("token")
		if token == "" {
			//TODO 后续进行Token验证
		}

		//2.执行请求
		c.Next()

		//3.打印请求状态
		status := c.Writer.Status()
		log.Printf("request status is %v\n", status)
	}
}
