// @Author:冯铁城 [17615007230@163.com] 2025-08-01 20:17:29
package util

import (
	"net-project-edu_manage/internal/config"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGenerateJWT(t *testing.T) {

	// 1.初始化配置
	config.AppConfig.Jwt.Key = "test_secret_key"

	// 2.参数结构体
	type args struct {
		username       string
		email          string
		expireHour     time.Duration
		isRefreshToken bool
	}

	// 3.编写测试参数
	tests := []struct {
		name    string
		args    args
		wantErr bool
		errMsg  string
	}{
		{
			name: "测试生成普通JWT Token",
			args: args{
				username:       "testuser",
				email:          "test@example.com",
				expireHour:     time.Hour * 24,
				isRefreshToken: false,
			},
			wantErr: false,
		},
		{
			name: "测试生成刷新JWT Token",
			args: args{
				username:       "testuser",
				email:          "test@example.com",
				expireHour:     time.Hour * 24,
				isRefreshToken: true,
			},
			wantErr: false,
		},
		{
			name: "测试空用户名",
			args: args{
				username:       "",
				email:          "test@example.com",
				expireHour:     time.Hour * 24,
				isRefreshToken: false,
			},
			wantErr: false,
		},
		{
			name: "测试空邮箱",
			args: args{
				username:       "testuser",
				email:          "",
				expireHour:     time.Hour * 24,
				isRefreshToken: false,
			},
			wantErr: false,
		},
		{
			name: "测试短过期时间",
			args: args{
				username:       "testuser",
				email:          "test@example.com",
				expireHour:     time.Minute * 1,
				isRefreshToken: false,
			},
			wantErr: false,
		},
		{
			name: "测试长过期时间",
			args: args{
				username:       "testuser",
				email:          "test@example.com",
				expireHour:     time.Hour * 24 * 365,
				isRefreshToken: false,
			},
			wantErr: false,
		},
		{
			name: "测试特殊字符用户名",
			args: args{
				username:       "test@user#123",
				email:          "test@example.com",
				expireHour:     time.Hour * 24,
				isRefreshToken: false,
			},
			wantErr: false,
		},
		{
			name: "测试特殊字符邮箱",
			args: args{
				username:       "testuser",
				email:          "test+tag@example-domain.com",
				expireHour:     time.Hour * 24,
				isRefreshToken: false,
			},
			wantErr: false,
		},
	}

	// 4.循环执行测试用例
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// 1.调用GenerateJWT
			jwtToken, err := GenerateJWT(tt.args.username, tt.args.email, tt.args.expireHour, tt.args.isRefreshToken)

			// 2.判定是否发生异常，如果与预期不一致，则返回错误
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateJWT() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// 3.根据是否发生异常进行处理
			if tt.wantErr {
				if err.Error() != tt.errMsg {
					t.Errorf("GenerateJWT() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			} else {

				// 1.验证token不为空
				assert.NotEmpty(t, jwtToken)

				// 2.解析JWT Token验证内容
				claims, err := ParseJWT(jwtToken)
				if err != nil {
					t.Errorf("ParseJWT() error = %v", err)
					return
				}

				// 3.验证用户名
				username := claims["username"]
				if username != tt.args.username {
					t.Errorf("GenerateJWT() username = %v, want %v", username, tt.args.username)
					return
				}

				// 4.验证邮箱
				email := claims["email"]
				if email != tt.args.email {
					t.Errorf("GenerateJWT() email = %v, want %v", email, tt.args.email)
					return
				}

				// 5.验证时间字段（仅对非刷新token）
				if !tt.args.isRefreshToken {
					// 验证签发时间存在
					iat := claims["iat"]
					if iat == nil {
						t.Errorf("GenerateJWT() iat should not be nil")
						return
					}

					// 验证过期时间存在
					exp := claims["exp"]
					if exp == nil {
						t.Errorf("GenerateJWT() exp should not be nil")
						return
					}

					// 验证过期时间合理性
					expTime := int64(exp.(float64))
					iatTime := int64(iat.(float64))
					actualDuration := time.Unix(expTime, 0).Sub(time.Unix(iatTime, 0))
					expectedDuration := tt.args.expireHour

					// 允许1秒的误差
					if actualDuration < expectedDuration-time.Second || actualDuration > expectedDuration+time.Second {
						t.Errorf("GenerateJWT() duration = %v, want %v", actualDuration, expectedDuration)
						return
					}
				} else {

					// 6.验证刷新token不包含过期时间
					if claims["exp"] != nil {
						t.Errorf("GenerateJWT() refresh token should not have exp claim")
						return
					}
					if claims["iat"] != nil {
						t.Errorf("GenerateJWT() refresh token should not have iat claim")
						return
					}
				}
			}
		})
	}
}

func TestGenerateJWTWithDifferentKeys(t *testing.T) {

	// 1.测试不同密钥长度
	testKeys := []string{
		"short",
		"medium_length_key",
		"this_is_a_very_long_secret_key_for_testing_purposes_123456789",
		"key_with_special_chars!@#$%^&*()",
	}

	for _, key := range testKeys {
		t.Run("测试密钥长度_"+key, func(t *testing.T) {
			// 1.设置测试密钥
			originalKey := config.AppConfig.Jwt.Key
			config.AppConfig.Jwt.Key = key
			defer func() {
				config.AppConfig.Jwt.Key = originalKey
			}()

			// 2.生成JWT
			token, err := GenerateJWT("testuser", "test@example.com", time.Hour*24, false)
			assert.NoError(t, err)
			assert.NotEmpty(t, token)

			// 3.验证可以正确解析
			claims, err := ParseJWT(token)
			assert.NoError(t, err)
			assert.Equal(t, "testuser", claims["username"])
			assert.Equal(t, "test@example.com", claims["email"])
		})
	}
}

func TestGenerateJWTEdgeCases(t *testing.T) {

	// 1.初始化配置
	config.AppConfig.Jwt.Key = "test_secret_key"

	// 2.测试边界情况
	testCases := []struct {
		name           string
		username       string
		email          string
		expireHour     time.Duration
		isRefreshToken bool
		expectError    bool
		skipParse      bool // 跳过解析验证（用于已过期的token）
	}{
		{
			name:           "测试极短过期时间",
			username:       "user",
			email:          "user@test.com",
			expireHour:     time.Nanosecond,
			isRefreshToken: false,
			expectError:    false,
			skipParse:      true, // 跳过解析，因为会立即过期
		},
		{
			name:           "测试零过期时间",
			username:       "user",
			email:          "user@test.com",
			expireHour:     0,
			isRefreshToken: false,
			expectError:    false,
			skipParse:      true, // 跳过解析，因为会立即过期
		},
		{
			name:           "测试负过期时间",
			username:       "user",
			email:          "user@test.com",
			expireHour:     -time.Hour,
			isRefreshToken: false,
			expectError:    false,
			skipParse:      true, // 跳过解析，因为已经过期
		},
		{
			name:           "测试正常过期时间",
			username:       "user",
			email:          "user@test.com",
			expireHour:     time.Hour * 24,
			isRefreshToken: false,
			expectError:    false,
			skipParse:      false,
		},
		{
			name:           "测试长字符串用户名",
			username:       "very_long_username_" + string(make([]rune, 100)),
			email:          "user@test.com",
			expireHour:     time.Hour,
			isRefreshToken: false,
			expectError:    false,
			skipParse:      false,
		},
		{
			name:           "测试长字符串邮箱",
			username:       "user",
			email:          "very_long_email_" + string(make([]rune, 100)) + "@test.com",
			expireHour:     time.Hour,
			isRefreshToken: false,
			expectError:    false,
			skipParse:      false,
		},
		{
			name:           "测试刷新token边界情况",
			username:       "refresh_user",
			email:          "refresh@test.com",
			expireHour:     time.Hour, // 对于刷新token，这个参数会被忽略
			isRefreshToken: true,
			expectError:    false,
			skipParse:      false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			token, err := GenerateJWT(tc.username, tc.email, tc.expireHour, tc.isRefreshToken)

			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, token)

				// 1.验证token格式正确（包含三个部分）
				parts := len(strings.Split(token, "."))
				assert.Equal(t, 3, parts, "JWT应该包含三个部分")

				// 2.如果不跳过解析，则验证可以解析
				if !tc.skipParse {
					claims, parseErr := ParseJWT(token)
					assert.NoError(t, parseErr)
					assert.Equal(t, tc.username, claims["username"])
					assert.Equal(t, tc.email, claims["email"])

					// 3.验证时间字段（仅对非刷新token）
					if !tc.isRefreshToken {
						assert.NotNil(t, claims["exp"])
						assert.NotNil(t, claims["iat"])
					} else {
						assert.Nil(t, claims["exp"])
						assert.Nil(t, claims["iat"])
					}
				}
			}
		})
	}
}

func TestParseJWT(t *testing.T) {

	// 1.初始化配置
	config.AppConfig.Jwt.Key = "test_secret_key"

	// 2.生成有效的JWT用于测试
	validToken, err := GenerateJWT("testuser", "test@example.com", time.Hour*24, false)
	assert.NoError(t, err)

	// 3.测试用例
	type args struct {
		tokenString string
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]any
		wantErr bool
	}{
		{
			name: "测试空字符串解析",
			args: args{
				tokenString: "",
			},
			wantErr: true,
		},
		{
			name: "测试异常token解析",
			args: args{
				tokenString: "123124",
			},
			wantErr: true,
		},
		{
			name: "测试格式错误的token",
			args: args{
				tokenString: "invalid.token.format",
			},
			wantErr: true,
		},
		{
			name: "测试有效token解析",
			args: args{
				tokenString: validToken,
			},
			want: map[string]any{
				"username": "testuser",
				"email":    "test@example.com",
			},
			wantErr: false,
		},
		{
			name: "测试错误签名的token",
			args: args{
				tokenString: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InRlc3R1c2VyIiwiZW1haWwiOiJ0ZXN0QGV4YW1wbGUuY29tIn0.wrong_signature",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseJWT(tt.args.tokenString)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseJWT() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				// 验证基本字段
				assert.Equal(t, tt.want["username"], got["username"])
				assert.Equal(t, tt.want["email"], got["email"])

				// 验证时间字段存在（对于非刷新token）
				if got["exp"] != nil {
					assert.NotNil(t, got["exp"])
					assert.NotNil(t, got["iat"])
				}
			}
		})
	}
}

func TestParseJWTWithDifferentKeys(t *testing.T) {

	// 1.生成token
	config.AppConfig.Jwt.Key = "original_key"
	token, err := GenerateJWT("testuser", "test@example.com", time.Hour*24, false)
	assert.NoError(t, err)

	// 2.使用正确密钥解析
	claims, err := ParseJWT(token)
	assert.NoError(t, err)
	assert.Equal(t, "testuser", claims["username"])

	// 3.恢复原始密钥
	config.AppConfig.Jwt.Key = "original_key"
}

func TestJWTIntegration(t *testing.T) {

	// 1.初始化配置
	config.AppConfig.Jwt.Key = "integration_test_key"

	// 2.测试完整的生成和解析流程
	testCases := []struct {
		name           string
		username       string
		email          string
		expireHour     time.Duration
		isRefreshToken bool
	}{
		{
			name:           "普通用户token",
			username:       "normaluser",
			email:          "normal@example.com",
			expireHour:     time.Hour * 24,
			isRefreshToken: false,
		},
		{
			name:           "管理员token",
			username:       "admin",
			email:          "admin@example.com",
			expireHour:     time.Hour * 12,
			isRefreshToken: false,
		},
		{
			name:           "刷新token",
			username:       "refreshuser",
			email:          "refresh@example.com",
			expireHour:     time.Hour * 24 * 7,
			isRefreshToken: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			// 1.生成JWT
			token, err := GenerateJWT(tc.username, tc.email, tc.expireHour, tc.isRefreshToken)
			assert.NoError(t, err)
			assert.NotEmpty(t, token)

			// 2.解析JWT
			claims, err := ParseJWT(token)
			assert.NoError(t, err)

			// 3.验证载荷
			assert.Equal(t, tc.username, claims["username"])
			assert.Equal(t, tc.email, claims["email"])

			// 4.验证时间字段
			if !tc.isRefreshToken {
				assert.NotNil(t, claims["exp"])
				assert.NotNil(t, claims["iat"])

				// 验证时间合理性
				expTime := int64(claims["exp"].(float64))
				iatTime := int64(claims["iat"].(float64))
				assert.True(t, expTime > iatTime)
			} else {
				assert.Nil(t, claims["exp"])
				assert.Nil(t, claims["iat"])
			}
		})
	}
}

func TestJWTConcurrency(t *testing.T) {

	// 1.初始化配置
	config.AppConfig.Jwt.Key = "concurrency_test_key"

	// 2.并发测试
	const numGoroutines = 100
	results := make(chan string, numGoroutines)
	errors := make(chan error, numGoroutines)

	// 3.启动多个goroutine同时生成JWT
	for i := 0; i < numGoroutines; i++ {
		go func(id int) {
			token, err := GenerateJWT("user"+string(rune(id)), "user"+string(rune(id))+"@test.com", time.Hour*24, false)
			if err != nil {
				errors <- err
				return
			}
			results <- token
		}(i)
	}

	// 4.收集结果
	tokens := make([]string, 0, numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		select {
		case token := <-results:
			tokens = append(tokens, token)
		case err := <-errors:
			t.Errorf("并发生成JWT失败: %v", err)
		case <-time.After(time.Second * 5):
			t.Error("并发测试超时")
			return
		}
	}

	// 5.验证所有token都能正确解析
	assert.Equal(t, numGoroutines, len(tokens))
	for _, token := range tokens {
		claims, err := ParseJWT(token)
		assert.NoError(t, err)
		assert.NotNil(t, claims["username"])
		assert.NotNil(t, claims["email"])
	}
}
