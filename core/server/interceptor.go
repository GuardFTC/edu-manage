// Package server @Author:冯铁城 [17615007230@163.com] 2025-07-30 16:20:53
package server

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"net-project-edu_manage/common/res"
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

// errorHandler 错误处理器
func errorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

		//1.defer and recover 捕捉异常
		defer func() {

			//2.如果存在异常
			if err := recover(); err != nil {

				//3.打印异常日志
				log.Errorf("Panic recovered: %v", err)

				//4.返回统一 JSON 错误响应
				fail := res.ServerFail(cast.ToString(err))
				c.JSON(fail.Code, fail.ToJson())

				//5.中断后续处理
				c.Abort()
			}
		}()

		//6.继续处理请求
		c.Next()
	}
}
