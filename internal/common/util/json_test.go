// @Author:冯铁城 [17615007230@163.com] 2025-08-06 17:35:00
package util

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToJSON(t *testing.T) {

	// 1.测试用例结构体
	type args struct {
		v interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "测试字符串转JSON",
			args: args{
				v: "hello world",
			},
			want:    `"hello world"`,
			wantErr: false,
		},
		{
			name: "测试整数转JSON",
			args: args{
				v: 123,
			},
			want:    "123",
			wantErr: false,
		},
		{
			name: "测试浮点数转JSON",
			args: args{
				v: 3.14,
			},
			want:    "3.14",
			wantErr: false,
		},
		{
			name: "测试布尔值转JSON",
			args: args{
				v: true,
			},
			want:    "true",
			wantErr: false,
		},
		{
			name: "测试nil转JSON",
			args: args{
				v: nil,
			},
			want:    "null",
			wantErr: false,
		},
		{
			name: "测试空字符串转JSON",
			args: args{
				v: "",
			},
			want:    `""`,
			wantErr: false,
		},
		{
			name: "测试结构体转JSON",
			args: args{
				v: struct {
					Name string `json:"name"`
					Age  int    `json:"age"`
				}{
					Name: "张三",
					Age:  25,
				},
			},
			want:    `{"name":"张三","age":25}`,
			wantErr: false,
		},
		{
			name: "测试切片转JSON",
			args: args{
				v: []string{"apple", "banana", "cherry"},
			},
			want:    `["apple","banana","cherry"]`,
			wantErr: false,
		},
		{
			name: "测试映射转JSON",
			args: args{
				v: map[string]interface{}{
					"name": "李四",
					"age":  30,
				},
			},
			want:    `{"age":30,"name":"李四"}`,
			wantErr: false,
		},
		{
			name: "测试包含特殊字符的字符串",
			args: args{
				v: "hello\nworld\t\"test\"",
			},
			want:    `"hello\nworld\t\"test\""`,
			wantErr: false,
		},
	}

	// 2.循环执行测试用例
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// 1.调用被测试函数
			got, err := ToJSON(tt.args.v)

			// 2.验证错误情况
			if (err != nil) != tt.wantErr {
				t.Errorf("ToJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// 3.验证返回值
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestFromJSON(t *testing.T) {

	// 1.测试用例结构体
	type TestStruct struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	type args struct {
		data string
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name: "测试解析字符串",
			args: args{
				data: `"hello world"`,
			},
			want:    "hello world",
			wantErr: false,
		},
		{
			name: "测试解析整数",
			args: args{
				data: "123",
			},
			want:    float64(123), // JSON数字默认解析为float64
			wantErr: false,
		},
		{
			name: "测试解析浮点数",
			args: args{
				data: "3.14",
			},
			want:    3.14,
			wantErr: false,
		},
		{
			name: "测试解析布尔值",
			args: args{
				data: "true",
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "测试解析null",
			args: args{
				data: "null",
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "测试解析切片",
			args: args{
				data: `["apple","banana","cherry"]`,
			},
			want:    []interface{}{"apple", "banana", "cherry"},
			wantErr: false,
		},
		{
			name: "测试解析映射",
			args: args{
				data: `{"name":"张三","age":25}`,
			},
			want: map[string]interface{}{
				"name": "张三",
				"age":  float64(25), // JSON数字默认解析为float64
			},
			wantErr: false,
		},
		{
			name: "测试解析无效JSON",
			args: args{
				data: `{"name":"张三","age":}`,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "测试解析空字符串",
			args: args{
				data: "",
			},
			want:    nil,
			wantErr: true,
		},
	}

	// 2.循环执行测试用例
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// 1.调用被测试函数
			got, err := FromJSON[interface{}](tt.args.data)

			// 2.验证错误情况
			if (err != nil) != tt.wantErr {
				t.Errorf("FromJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// 3.验证返回值
			if !tt.wantErr {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestFromJSONWithSpecificTypes(t *testing.T) {

	// 1.测试解析到具体类型
	t.Run("测试解析到字符串类型", func(t *testing.T) {

		// 1.解析字符串
		result, err := FromJSON[string](`"hello world"`)
		assert.NoError(t, err)
		assert.Equal(t, "hello world", result)

		// 2.解析错误类型
		_, err = FromJSON[string]("123")
		assert.Error(t, err)
	})

	// 2.测试解析到整数类型
	t.Run("测试解析到整数类型", func(t *testing.T) {

		// 1.解析整数
		result, err := FromJSON[int]("123")
		assert.NoError(t, err)
		assert.Equal(t, 123, result)

		// 2.解析错误类型
		_, err = FromJSON[int](`"hello"`)
		assert.Error(t, err)
	})

	// 3.测试解析到结构体类型
	t.Run("测试解析到结构体类型", func(t *testing.T) {

		// 1.定义测试结构体
		type User struct {
			Name string `json:"name"`
			Age  int    `json:"age"`
		}

		// 2.解析结构体
		result, err := FromJSON[User](`{"name":"张三","age":25}`)
		assert.NoError(t, err)
		assert.Equal(t, "张三", result.Name)
		assert.Equal(t, 25, result.Age)

		// 3.解析错误格式
		_, err = FromJSON[User](`{"name":"张三","age":"invalid"}`)
		assert.Error(t, err)
	})

	// 4.测试解析到切片类型
	t.Run("测试解析到切片类型", func(t *testing.T) {

		// 1.解析字符串切片
		result, err := FromJSON[[]string](`["a","b","c"]`)
		assert.NoError(t, err)
		assert.Equal(t, []string{"a", "b", "c"}, result)

		// 2.解析整数切片
		intResult, err := FromJSON[[]int](`[1,2,3]`)
		assert.NoError(t, err)
		assert.Equal(t, []int{1, 2, 3}, intResult)
	})

	// 5.测试解析到映射类型
	t.Run("测试解析到映射类型", func(t *testing.T) {

		// 1.解析字符串映射
		result, err := FromJSON[map[string]string](`{"key1":"value1","key2":"value2"}`)
		assert.NoError(t, err)
		expected := map[string]string{
			"key1": "value1",
			"key2": "value2",
		}
		assert.Equal(t, expected, result)

		// 2.解析混合类型映射
		mixedResult, err := FromJSON[map[string]interface{}](`{"name":"张三","age":25,"active":true}`)
		assert.NoError(t, err)
		assert.Equal(t, "张三", mixedResult["name"])
		assert.Equal(t, float64(25), mixedResult["age"]) // JSON数字解析为float64
		assert.Equal(t, true, mixedResult["active"])
	})
}

func TestPrettyJSON(t *testing.T) {

	// 1.测试用例结构体
	type args struct {
		v interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		check   func(string) bool
	}{
		{
			name: "测试格式化简单对象",
			args: args{
				v: map[string]interface{}{
					"name": "张三",
					"age":  25,
				},
			},
			wantErr: false,
			check: func(result string) bool {
				return strings.Contains(result, "  ") && // 包含缩进
					strings.Contains(result, "\n") && // 包含换行
					strings.Contains(result, "张三") &&
					strings.Contains(result, "25")
			},
		},
		{
			name: "测试格式化嵌套对象",
			args: args{
				v: map[string]interface{}{
					"user": map[string]interface{}{
						"name":  "李四",
						"email": "lisi@example.com",
					},
					"active": true,
				},
			},
			wantErr: false,
			check: func(result string) bool {
				return strings.Contains(result, "  ") && // 包含缩进
					strings.Contains(result, "\n") && // 包含换行
					strings.Contains(result, "李四") &&
					strings.Contains(result, "lisi@example.com")
			},
		},
		{
			name: "测试格式化数组",
			args: args{
				v: []map[string]interface{}{
					{"id": 1, "name": "用户1"},
					{"id": 2, "name": "用户2"},
				},
			},
			wantErr: false,
			check: func(result string) bool {
				return strings.Contains(result, "  ") && // 包含缩进
					strings.Contains(result, "\n") && // 包含换行
					strings.Contains(result, "用户1") &&
					strings.Contains(result, "用户2")
			},
		},
		{
			name: "测试格式化字符串",
			args: args{
				v: "simple string",
			},
			wantErr: false,
			check: func(result string) bool {
				return result == `"simple string"`
			},
		},
		{
			name: "测试格式化数字",
			args: args{
				v: 123,
			},
			wantErr: false,
			check: func(result string) bool {
				return result == "123"
			},
		},
	}

	// 2.循环执行测试用例
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// 1.调用被测试函数
			got, err := PrettyJSON(tt.args.v)

			// 2.验证错误情况
			if (err != nil) != tt.wantErr {
				t.Errorf("PrettyJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// 3.验证返回值
			if !tt.wantErr && tt.check != nil {
				assert.True(t, tt.check(got), "PrettyJSON输出格式不符合预期: %s", got)
			}
		})
	}
}

func TestJSONEdgeCases(t *testing.T) {

	// 1.边界条件测试
	t.Run("测试边界条件", func(t *testing.T) {

		// 1.测试最大整数
		maxInt := int64(9223372036854775807)
		jsonStr, err := ToJSON(maxInt)
		assert.NoError(t, err)
		result, err := FromJSON[int64](jsonStr)
		assert.NoError(t, err)
		assert.Equal(t, maxInt, result)

		// 2.测试最小整数
		minInt := int64(-9223372036854775808)
		jsonStr, err = ToJSON(minInt)
		assert.NoError(t, err)
		result, err = FromJSON[int64](jsonStr)
		assert.NoError(t, err)
		assert.Equal(t, minInt, result)

		// 3.测试空切片
		var emptySlice []string
		jsonStr, err = ToJSON(emptySlice)
		assert.NoError(t, err)
		assert.Equal(t, "[]", jsonStr)

		// 4.测试空映射
		emptyMap := map[string]interface{}{}
		jsonStr, err = ToJSON(emptyMap)
		assert.NoError(t, err)
		assert.Equal(t, "{}", jsonStr)

		// 5.测试包含特殊字符的字符串
		specialStr := "包含\n换行\t制表符\"引号的字符串"
		jsonStr, err = ToJSON(specialStr)
		assert.NoError(t, err)
		result2, err := FromJSON[string](jsonStr)
		assert.NoError(t, err)
		assert.Equal(t, specialStr, result2)
	})
}

func TestJSONInvalidCases(t *testing.T) {

	// 1.测试无效输入
	t.Run("测试无效JSON解析", func(t *testing.T) {

		// 1.测试无效JSON格式
		invalidJSONs := []string{
			`{"name":"张三","age":}`,
			`{"name":"张三",}`,
			`{name:"张三"}`,
			`{"name":"张三""age":25}`,
			`[1,2,3,]`,
			`{"unclosed": "string}`,
		}

		for _, invalidJSON := range invalidJSONs {
			_, err := FromJSON[interface{}](invalidJSON)
			assert.Error(t, err, "应该解析失败: %s", invalidJSON)
		}

		// 2.测试空字符串
		_, err := FromJSON[interface{}]("")
		assert.Error(t, err)

		// 3.测试类型不匹配
		_, err = FromJSON[int](`"not a number"`)
		assert.Error(t, err)

		_, err = FromJSON[string]("123")
		assert.Error(t, err)
	})

	// 2.测试可能导致Marshal错误的输入
	t.Run("测试Marshal错误情况", func(t *testing.T) {

		// 1.测试函数类型（不能序列化）
		_, err := ToJSON(func() {})
		assert.Error(t, err)

		// 2.测试通道类型（不能序列化）
		_, err = ToJSON(make(chan int))
		assert.Error(t, err)

		// 3.测试包含不可序列化字段的结构体
		type InvalidStruct struct {
			Name string
			Func func()
		}
		_, err = ToJSON(InvalidStruct{Name: "test", Func: func() {}})
		assert.Error(t, err)
	})
}

func TestJSONRoundTrip(t *testing.T) {

	// 1.往返转换测试
	testCases := []interface{}{
		"测试字符串",
		123,
		3.14159,
		true,
		false,
		[]string{"a", "b", "c"},
		[]int{1, 2, 3, 4, 5},
		map[string]interface{}{
			"name":   "张三",
			"age":    25,
			"active": true,
		},
		struct {
			Name  string                 `json:"name"`
			Age   int                    `json:"age"`
			Tags  []string               `json:"tags"`
			Extra map[string]interface{} `json:"extra"`
		}{
			Name: "李四",
			Age:  30,
			Tags: []string{"开发者", "Go语言"},
			Extra: map[string]interface{}{
				"level":  "senior",
				"salary": 15000,
			},
		},
	}

	// 2.执行往返测试
	for i, testCase := range testCases {
		t.Run("往返测试_"+string(rune(i)), func(t *testing.T) {

			// 1.序列化
			jsonStr, err := ToJSON(testCase)
			assert.NoError(t, err)
			assert.NotEmpty(t, jsonStr)

			// 2.反序列化
			var result interface{}
			result, err = FromJSON[interface{}](jsonStr)
			assert.NoError(t, err)

			// 3.验证基本一致性（注意JSON数字会变成float64）
			originalJSON, _ := json.Marshal(testCase)
			resultJSON, _ := json.Marshal(result)
			assert.JSONEq(t, string(originalJSON), string(resultJSON))
		})
	}
}

func TestJSONPerformance(t *testing.T) {

	// 1.性能测试
	t.Run("测试大对象序列化性能", func(t *testing.T) {

		// 1.创建大对象
		largeObject := make(map[string]interface{})
		for i := 0; i < 1000; i++ {
			largeObject["key"+string(rune(i))] = map[string]interface{}{
				"id":    i,
				"name":  "用户" + string(rune(i)),
				"email": "user" + string(rune(i)) + "@example.com",
				"data":  []int{i, i + 1, i + 2, i + 3, i + 4},
			}
		}

		// 2.测试ToJSON性能
		jsonStr, err := ToJSON(largeObject)
		assert.NoError(t, err)
		assert.NotEmpty(t, jsonStr)

		// 3.测试PrettyJSON性能
		prettyStr, err := PrettyJSON(largeObject)
		assert.NoError(t, err)
		assert.NotEmpty(t, prettyStr)
		assert.True(t, len(prettyStr) > len(jsonStr)) // 格式化后应该更长

		// 4.测试FromJSON性能
		var result map[string]interface{}
		result, err = FromJSON[map[string]interface{}](jsonStr)
		assert.NoError(t, err)
		assert.Equal(t, len(largeObject), len(result))
	})
}

func TestJSONConcurrency(t *testing.T) {

	// 1.并发测试
	t.Run("测试并发调用", func(t *testing.T) {

		// 1.准备测试数据
		testData := map[string]interface{}{
			"name":  "并发测试",
			"count": 12345,
			"data":  []string{"a", "b", "c", "d", "e"},
		}

		// 2.并发执行
		const numGoroutines = 100
		results := make(chan string, numGoroutines)
		errors := make(chan error, numGoroutines)

		for i := 0; i < numGoroutines; i++ {
			go func() {
				// 测试ToJSON
				jsonStr, err := ToJSON(testData)
				if err != nil {
					errors <- err
					return
				}

				// 测试FromJSON
				var result map[string]interface{}
				result, err = FromJSON[map[string]interface{}](jsonStr)
				if err != nil {
					errors <- err
					return
				}

				// 测试PrettyJSON
				_, err = PrettyJSON(result)
				if err != nil {
					errors <- err
					return
				}

				results <- jsonStr
			}()
		}

		// 3.收集结果
		successCount := 0
		for i := 0; i < numGoroutines; i++ {
			select {
			case result := <-results:
				assert.NotEmpty(t, result)
				successCount++
			case err := <-errors:
				t.Errorf("并发测试失败: %v", err)
			}
		}

		// 4.验证所有调用都成功
		assert.Equal(t, numGoroutines, successCount)
	})
}

func TestJSONIntegration(t *testing.T) {

	// 1.集成测试：模拟实际使用场景
	t.Run("测试API数据处理场景", func(t *testing.T) {

		// 1.模拟API请求数据
		type APIRequest struct {
			Method  string                 `json:"method"`
			URL     string                 `json:"url"`
			Headers map[string]string      `json:"headers"`
			Body    map[string]interface{} `json:"body"`
		}

		request := APIRequest{
			Method: "POST",
			URL:    "/api/users",
			Headers: map[string]string{
				"Content-Type":  "application/json",
				"Authorization": "Bearer token123",
			},
			Body: map[string]interface{}{
				"name":   "新用户",
				"email":  "newuser@example.com",
				"age":    28,
				"active": true,
				"tags":   []string{"新用户", "测试"},
				"metadata": map[string]interface{}{
					"source":    "web",
					"timestamp": 1625097600,
				},
			},
		}

		// 2.序列化请求
		requestJSON, err := ToJSON(request)
		assert.NoError(t, err)
		assert.NotEmpty(t, requestJSON)

		// 3.格式化输出（用于日志）
		prettyJSON, err := PrettyJSON(request)
		assert.NoError(t, err)
		assert.Contains(t, prettyJSON, "新用户")
		assert.Contains(t, prettyJSON, "  ") // 包含缩进

		// 4.反序列化请求
		var parsedRequest APIRequest
		parsedRequest, err = FromJSON[APIRequest](requestJSON)
		assert.NoError(t, err)
		assert.Equal(t, request.Method, parsedRequest.Method)
		assert.Equal(t, request.URL, parsedRequest.URL)
		assert.Equal(t, len(request.Headers), len(parsedRequest.Headers))

		// 5.验证嵌套数据
		assert.Equal(t, "新用户", parsedRequest.Body["name"])
		assert.Equal(t, "newuser@example.com", parsedRequest.Body["email"])
		assert.Equal(t, float64(28), parsedRequest.Body["age"]) // JSON数字解析为float64
		assert.Equal(t, true, parsedRequest.Body["active"])
	})

	// 2.测试配置文件处理场景
	t.Run("测试配置文件处理场景", func(t *testing.T) {

		// 1.模拟配置数据
		type DatabaseConfig struct {
			Host     string `json:"host"`
			Port     int    `json:"port"`
			Username string `json:"username"`
			Password string `json:"password"`
			Database string `json:"database"`
		}

		type AppConfig struct {
			AppName  string          `json:"app_name"`
			Debug    bool            `json:"debug"`
			Database DatabaseConfig  `json:"database"`
			Features map[string]bool `json:"features"`
		}

		config := AppConfig{
			AppName: "教育管理系统",
			Debug:   true,
			Database: DatabaseConfig{
				Host:     "localhost",
				Port:     5432,
				Username: "admin",
				Password: "password123",
				Database: "edu_manage",
			},
			Features: map[string]bool{
				"user_management":   true,
				"course_management": true,
				"grade_management":  false,
			},
		}

		// 2.序列化配置
		configJSON, err := ToJSON(config)
		assert.NoError(t, err)
		assert.Contains(t, configJSON, "教育管理系统")

		// 3.格式化配置（用于配置文件）
		prettyConfig, err := PrettyJSON(config)
		assert.NoError(t, err)
		assert.Contains(t, prettyConfig, "  ") // 包含缩进
		assert.Contains(t, prettyConfig, "\n") // 包含换行

		// 4.反序列化配置
		var parsedConfig AppConfig
		parsedConfig, err = FromJSON[AppConfig](configJSON)
		assert.NoError(t, err)
		assert.Equal(t, config.AppName, parsedConfig.AppName)
		assert.Equal(t, config.Debug, parsedConfig.Debug)
		assert.Equal(t, config.Database.Host, parsedConfig.Database.Host)
		assert.Equal(t, config.Database.Port, parsedConfig.Database.Port)
	})
}
