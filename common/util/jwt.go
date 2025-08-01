// Package util @Author:冯铁城 [17615007230@163.com] 2025-08-01 19:45:55
package util

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	log "github.com/sirupsen/logrus"
	"net-project-edu_manage/config/config"
	"time"
)

// JwtKey
var key = []byte(config.AppConfig.Jwt.Key)

// GenerateJWT 生成JWT Token
func GenerateJWT(username string, email string, expireHour int) (string, error) {

	//1.获取过期时间,优先顺序为参数，然后为配置文件
	if expireHour == 0 {
		expireHour = config.AppConfig.Jwt.ExpireHour
	}

	//2.创建信息声明
	claims := jwt.MapClaims{
		"username": username,
		"email":    email,
		"exp":      time.Now().Add(time.Duration(expireHour) * time.Hour).Unix(), // 过期时间
		"iat":      time.Now().Unix(),                                            // 签发时间
	}

	//3.创建 token 对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//4.签名并获取完整的编码Token
	if jwtToken, err := token.SignedString(key); err != nil {
		return "", err
	} else {
		log.Printf("jwtToken init success:%v", jwtToken)
		return jwtToken, nil
	}
}

// ParseJWT 解析JWT Token
func ParseJWT(tokenString string) (map[string]any, error) {

	//1.解析并验证 token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {

		//2.验证签名方法是否为 HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return key, nil
	})

	//3.err判空
	if err != nil {
		return nil, err
	}

	//4.提取信息声明
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	//5.默认返回
	return nil, fmt.Errorf("invalid token")
}
