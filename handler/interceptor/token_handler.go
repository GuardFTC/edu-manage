// Package interceptor @Author:冯铁城 [17615007230@163.com] 2025-07-31 11:09:11
package interceptor

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// GetTokenHandler 获取Token处理器
func GetTokenHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

		//1.获取请求头
		token := c.GetHeader("token")
		if token == "" {
			//TODO 后续进行Token验证
		}

		//2.执行请求
		c.Next()

		//3.打印请求状态
		log.Printf("request status is %v\n", c.Writer.Status())
	}
}
