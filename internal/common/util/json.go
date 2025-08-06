// Package util @Author:冯铁城 [17615007230@163.com] 2025-08-06 15:37:28
package util

import "encoding/json"

// ToJSON 将任意对象转换为JSON字符串
func ToJSON(v any) (string, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// FromJSON 将JSON字符串转换为任意对象
func FromJSON[T any](data string) (T, error) {
	var obj T
	err := json.Unmarshal([]byte(data), &obj)
	return obj, err
}

// PrettyJSON 将任意对象转换为格式化JSON字符串
func PrettyJSON(v any) (string, error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return "", err
	}
	return string(b), nil
}
