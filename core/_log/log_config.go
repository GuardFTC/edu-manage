// @Author:冯铁城 [17615007230@163.com] 2025-07-30 17:27:03
package _log

import (
	log "github.com/sirupsen/logrus"
)

// InitLogger 日志初始化
func InitLogger() {

	//1.设置日志格式为带颜色的 TextFormatter
	log.SetFormatter(&log.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})

	//2.设置日志级别，默认 Info 及以上输出
	log.SetLevel(log.InfoLevel)
}
