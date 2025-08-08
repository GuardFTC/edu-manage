// Package validate @Author:冯铁城 [17615007230@163.com] 2025-08-04 17:36:50
package validate

import (
	"time"

	"github.com/go-playground/validator/v10"
)

// validateDatetime 验证日期时间格式
func validateDatetime(fl validator.FieldLevel) bool {
	dateStr := fl.Field().String()
	format := fl.Param()
	_, err := time.Parse(format, dateStr)
	return err == nil
}
