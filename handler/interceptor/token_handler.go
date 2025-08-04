// Package interceptor @Author:冯铁城 [17615007230@163.com] 2025-07-31 11:09:11
package interceptor

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"net-project-edu_manage/common/res"
	"net-project-edu_manage/common/util"
	"time"
)

// GetTokenHandler 获取Token处理器
func GetTokenHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

		//1.获取请求头
		token := c.GetHeader("token")
		if token == "" {
			util.FailResToC(c, res.UnauthorizedFail, "token is null")
			c.Abort()
			return
		}

		//2.解析JWT Token
		claims, err := util.ParseJWT(token)
		if err != nil {
			util.FailResToC(c, res.UnauthorizedFail, err.Error())
			c.Abort()
			return
		}

		//3.判定token是否为合法token
		iat := cast.ToInt64(claims["iat"])
		if time.Now().Unix() < iat {
			util.FailResToC(c, res.UnauthorizedFail, "token is invalid")
			c.Abort()
			return
		}

		//4.判定token是否过期
		exp := cast.ToInt64(claims["exp"])
		if time.Now().Unix() > exp {
			util.FailResToC(c, res.UnauthorizedFail, "token is expired")
			c.Abort()
			return
		}

		//5.TODO 后续进行权限比较

		//6.遍历claims，添加内容到上下文
		for key, value := range claims {
			util.AddKVToC(c, key, value)
		}

		//7.执行请求
		c.Next()

		//8.打印请求状态
		log.Printf("request status is %v\n", c.Writer.Status())
	}
}
