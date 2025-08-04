// @Author:冯铁城 [17615007230@163.com] 2025-08-04 14:15:22
package util

import "github.com/gin-gonic/gin"

// GetUsernameFromC 获取上下文中的用户名
func GetUsernameFromC(c *gin.Context) (string, bool) {
	return GetStringFromC(c, "username")
}

// GetEmailFromC 获取上下文中的邮箱
func GetEmailFromC(c *gin.Context) (string, bool) {
	return GetStringFromC(c, "email")
}

// GetStringFromC 获取上下文中的字符串
func GetStringFromC(c *gin.Context, key string) (string, bool) {
	return GetValueFromC[string](c, key)
}

// GetBoolFromC 获取上下文中的布尔值
func GetBoolFromC(c *gin.Context, key string) (bool, bool) {
	return GetValueFromC[bool](c, key)
}

// GetIntFromC 获取上下文中的整数
func GetIntFromC(c *gin.Context, key string) (int, bool) {
	return GetValueFromC[int](c, key)
}

// GetInt64FromC 获取上下文中的64位整数
func GetInt64FromC(c *gin.Context, key string) (int64, bool) {
	return GetValueFromC[int64](c, key)
}

// GetFloat64FromC 获取上下文中的浮点数
func GetFloat64FromC(c *gin.Context, key string) (float64, bool) {
	return GetValueFromC[float64](c, key)
}

// GetSliceFromC 获取上下文中的切片
func GetSliceFromC[T any](c *gin.Context, key string) ([]T, bool) {
	return GetValueFromC[[]T](c, key)
}

// GetMapFromC 获取上下文中的映射
func GetMapFromC[K comparable, V any](c *gin.Context, key string) (map[K]V, bool) {
	return GetValueFromC[map[K]V](c, key)
}

// GetValueFromC 获取上下文中的键值对
func GetValueFromC[T any](c *gin.Context, key string) (T, bool) {

	//1.判定值是否存在
	val, exists := c.Get(key)
	if !exists {
		var zero T
		return zero, false
	}

	//2.类型转换
	result, ok := val.(T)
	if !ok {
		var zero T
		return zero, false
	}

	//3.返回转换后的值
	return result, true
}

// AddKVToC 添加键值对到上下文
func AddKVToC(c *gin.Context, key string, value interface{}) {
	c.Set(key, value)
}
