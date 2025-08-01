// @Author:冯铁城 [17615007230@163.com] 2025-08-01 20:17:29
package util

import (
	"net-project-edu_manage/config/config"
	"reflect"
	"testing"
)

func TestGenerateJWT(t *testing.T) {

	//1.配置初始化
	config.InitConfig()

	//2.参数结构体
	type args struct {
		username   string
		email      string
		expireHour int
	}

	//3.编写测试参数
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
		errMsg  string
	}{
		{
			name: "测试不设置过期时间",
			args: args{
				username:   "fengtiecheng",
				email:      "17615007230@163.com",
				expireHour: 0,
			},
			wantErr: false,
		},
		{
			name: "测试设置过期时间",
			args: args{
				username:   "fengtiecheng",
				email:      "17615007230@163.com",
				expireHour: 2,
			},
			wantErr: false,
		},
	}

	//4.循环执行测试用例
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			//1.调用GenerateJWT
			jwtToken, err := GenerateJWT(tt.args.username, tt.args.email, tt.args.expireHour)

			//2.判定是否发生异常，如果与预期不一致，则返回错误
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateJWT() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			//3.根据是否发生异常进行处理
			//如果发生异常，校验异常信息
			//如果未发生异常，则解析JWT Token
			if tt.wantErr {
				if err.Error() != tt.errMsg {
					t.Errorf("GenerateJWT() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			} else {
				claims, err := ParseJWT(jwtToken)
				if err != nil {
					t.Errorf("ParseJWT() error = %v, wantErr %v", err, tt.wantErr)
					return
				}

				//1.判定username
				username := claims["username"]
				if username != tt.args.username {
					t.Errorf("GenerateJWT() username = %v, want %v", username, tt.args.username)
					return
				}

				//2.判定email
				email := claims["email"]
				if email != tt.args.email {
					t.Errorf("GenerateJWT() email = %v, want %v", email, tt.args.email)
					return
				}

				//3.判定expireHour
				exp := claims["exp"]
				if exp != tt.args.expireHour {
					t.Errorf("GenerateJWT() exp = %v, want %v", exp, tt.args.expireHour)
					return
				}

				//4.判定iat
				iat := claims["iat"]
				if iat == nil {
					t.Errorf("GenerateJWT() iat = %v, want %v", iat, tt.args.expireHour)
					return
				}
			}
		})
	}
}

func TestParseJWT(t *testing.T) {
	type args struct {
		tokenString string
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]any
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseJWT(tt.args.tokenString)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseJWT() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseJWT() got = %v, want %v", got, tt.want)
			}
		})
	}
}
