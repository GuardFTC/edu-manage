// @Author:冯铁城 [17615007230@163.com] 2025-07-30 15:51:47
package db

import (
	"gorm.io/gorm"
	"net-project-edu_manage/internal/config"
	"testing"
)

func Test_generate1(t *testing.T) {

	//1.初始化配置
	config.InitConfig()

	//2.初始化DB
	InitDbConn()

	//3.确保最终关闭数据库链接
	defer CloseDbConn()

	//4.定义参数结构体
	type args struct {
		db           *gorm.DB
		tables       []string
		outPath      string
		modelPkgPath string
	}

	//5.定义测试用例
	tests := []struct {
		name string
		args args
	}{
		{
			name: "生成DAO层、Model代码",
			args: args{
				db: DB,
				tables: []string{
					"system_user",
					"academic_year",
					"grade",
					"grade_year",
				},
				outPath:      "query",
				modelPkgPath: "model",
			},
		},
	}

	//6.执行测试（也就是生成代码）
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			generate(tt.args.db, tt.args.tables, tt.args.outPath, tt.args.modelPkgPath)
		})
	}
}
