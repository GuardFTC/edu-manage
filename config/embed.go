// Package config @Author:冯铁城 [17615007230@163.com] 2025-08-04 11:23:39
package config

import "embed"

//go:embed resources/*.yaml
var ConfigFiles embed.FS
