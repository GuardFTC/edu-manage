// Package main @Author:冯铁城 [17615007230@163.com] 2025-08-08 10:28:09
package main

import (
	"net-project-edu_manage/internal/config"
	"net-project-edu_manage/internal/infrastructure/db"

	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {

	//1.初始化配置
	config.InitConfig()

	//2.初始化DB
	db.InitDbConn(&config.AppConfig.DataBase)

	//3.确保最终关闭数据库链接
	defer db.CloseDbConn()

	//4.执行主数据源代码生成
	generate(
		db.GetDefaultDataSource().GetDB(),
		[]string{
			"system_user",
			"academic_year",
			"grade",
			"grade_year",
			"class",
		},
		"internal/infrastructure/db/master/query",
		"model",
	)

	//5.执行从数据源代码生成
	generate(
		db.GetDataSource("slave1").GetDB(),
		[]string{
			"system_user",
		},
		"internal/infrastructure/db/slave1/query",
		"model",
	)
}

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
