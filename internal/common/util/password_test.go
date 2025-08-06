// @Author:冯铁城 [17615007230@163.com] 2025-07-30 14:59:18
package util

import "testing"

func TestHashPassword(t *testing.T) {
	type args struct {
		password string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
		errMsg  string
	}{
		{
			name:    "密码为空",
			args:    args{password: ""},
			wantErr: true,
			errMsg:  "密码不能为空",
		},
		{
			name:    "正常密码",
			args:    args{password: "123456"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			//1.调用HashPassword
			got, err := HashPassword(tt.args.password)

			//2.判定是否发生异常，如果与预期不一致，则返回错误
			if (err != nil) != tt.wantErr {
				t.Errorf("HashPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			//3.根据是否发生异常进行处理
			//如果发生异常，校验异常信息
			//如果未发生异常，则校验密码
			if tt.wantErr {
				if err.Error() != tt.errMsg {
					t.Errorf("HashPassword() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			} else {
				err = VerifyPassword(got, tt.args.password)
				if err != nil {
					t.Errorf("VerifyPassword() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			}
		})
	}
}

func TestVerifyPassword(t *testing.T) {
	type args struct {
		password   string
		dbPassword string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "输入密码为空",
			args:    args{password: "", dbPassword: "123456"},
			wantErr: true,
		},
		{
			name:    "数据库密码为空",
			args:    args{password: "123456", dbPassword: ""},
			wantErr: true,
		},
		{
			name:    "密码都为空",
			args:    args{password: "", dbPassword: ""},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := VerifyPassword(tt.args.password, tt.args.dbPassword); (err != nil) != tt.wantErr {
				t.Errorf("VerifyPassword() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
