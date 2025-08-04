// Package vo @Author:冯铁城 [17615007230@163.com] 2025-08-04 20:26:31
package vo

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
	"time"
)

// FormatTime 自定义时间类型，支持自定义时间格式序列化
type FormatTime struct {
	time.Time        // 嵌入原始 time.Time
	Layout    string // 时间格式，如："2006-01-02 15:04:05"
}

// 默认时间格式,相当于 Java 的 yyyy-MM-dd HH:mm:ss
const defaultLayout = time.DateTime

// NewFormatTime 工厂方法，用于创建 FormatTime 实例
func NewFormatTime(t time.Time, layout string) FormatTime {
	return FormatTime{
		Time:   t,
		Layout: layout,
	}
}

// MarshalJSON 自定义序列化方法，输出指定格式时间字符串
func (t *FormatTime) MarshalJSON() ([]byte, error) {

	//1.获取格式化模版
	layout := t.Layout
	if layout == "" {
		layout = defaultLayout
	}

	//2.格式化代码
	formatted := fmt.Sprintf("\"%s\"", t.Format(layout))

	//3.返回
	return []byte(formatted), nil
}

// UnmarshalJSON 自定义反序列化方法，将字符串解析为时间
func (t *FormatTime) UnmarshalJSON(b []byte) error {

	//1.获取格式化模版
	layout := t.Layout
	if layout == "" {
		layout = defaultLayout
	}

	//2.获取时间字符串
	str := strings.Trim(string(b), `"`)

	//3.解析时间
	parsed, err := time.Parse(layout, str)
	if err != nil {
		return err
	}

	//4.赋值
	t.Time = parsed

	//5.默认返回
	return nil
}

// Scan 实现 sql.Scanner 接口（用于从 DB 读取数据）
func (t *FormatTime) Scan(value interface{}) error {

	//1.断言为 time.Time 类型
	v, ok := value.(time.Time)
	if !ok {
		return errors.New("FormatTime: cannot scan non-time value")
	}

	//2.赋值
	t.Time = v

	//3.默认返回
	return nil
}

// Value 实现 driver.Valuer 接口（用于写入 DB）
func (t *FormatTime) Value() (driver.Value, error) {
	return t.Time, nil
}
