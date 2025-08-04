// Package validate @Author:冯铁城 [17615007230@163.com] 2025-08-04 17:35:56
package validate

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// InitValidate 初始化校验器
func InitValidate() {

	//1.获取校验器
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok {
		panic("validator.Validate 获取失败")
	}

	//2.注册自定义时间校验器
	if err := v.RegisterValidation("datetime", validateDatetime); err != nil {
		panic("datetime 校验器注册失败")
	}
}
