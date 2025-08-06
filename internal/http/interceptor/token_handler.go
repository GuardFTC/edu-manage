// Package interceptor @Author:冯铁城 [17615007230@163.com] 2025-07-31 11:09:11
package interceptor

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"net-project-edu_manage/internal/common/util"
	"net-project-edu_manage/internal/model/res"
	"time"
)

// GetTokenHandler 获取Token处理器
func GetTokenHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

		//1.校验Token,并获取Token参数
		claims, ok := checkJWTToken(c)
		if !ok {
			return
		}

		//2.TODO 后续进行权限比较

		//3.遍历claims，添加内容到上下文
		for key, value := range claims {
			util.AddKVToC(c, key, value)
		}

		//4.执行请求
		c.Next()
	}
}

// checkJWTToken 检测JWT Token
func checkJWTToken(c *gin.Context) (map[string]any, bool) {

	//1.获取JWT Token
	token := c.GetHeader("token")
	if token == "" {
		res.FailResToC(c, res.UnauthorizedFail, "token is null")
		return nil, false
	}

	//2.解析JWT Token
	claims, err := util.ParseJWT(token)
	if err != nil {
		res.FailResToC(c, res.UnauthorizedFail, err.Error())
		return nil, false
	}

	//3.判定token是否为合法token
	iat := cast.ToInt64(claims["iat"])
	if time.Now().Unix() < iat {
		res.FailResToC(c, res.UnauthorizedFail, "token is invalid")
		return nil, false
	}

	//4.判定token是否过期
	exp := cast.ToInt64(claims["exp"])
	if time.Now().Unix() > exp {
		res.FailResToC(c, res.UnauthorizedFail, "token is expired")
		return nil, false
	}

	//5.默认返回
	return claims, true
}
