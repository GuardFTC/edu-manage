// Package dao @Author:冯铁城 [17615007230@163.com] 2025-07-30 15:36:17
package dao

import (
	"gorm.io/gen"
	"gorm.io/gorm"
)

// Generate 生成代码
func generate(db *gorm.DB, tables []string, outPath string, modelPkgPath string) {

	//1.创建自动生成器
	generator := gen.NewGenerator(gen.Config{
		OutPath:      outPath,
		ModelPkgPath: modelPkgPath,
		Mode:         gen.WithDefaultQuery | gen.WithQueryInterface | gen.WithoutContext,
	})

	//2.使用Gorm数据源
	generator.UseDB(db)

	//3.循环拼接models
	var models []any
	for _, table := range tables {
		models = append(models, generator.GenerateModel(table))
	}

	//4.为入参表自动生成相关代码
	generator.ApplyBasic(models...)

	//5.执行
	generator.Execute()
}
