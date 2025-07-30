// @Author:冯铁城 [17615007230@163.com] 2025-07-30 15:51:47
package dao

import (
	"gorm.io/gorm"
	"testing"
)

func Test_generate1(t *testing.T) {
	InitDB()
	type args struct {
		db           *gorm.DB
		tables       []string
		outPath      string
		modelPkgPath string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "生成DAO层、Model代码",
			args: args{
				db:           DB,
				tables:       []string{"system_user"},
				outPath:      "query/system",
				modelPkgPath: "../model/system",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			generate(tt.args.db, tt.args.tables, tt.args.outPath, tt.args.modelPkgPath)
		})
	}
}
