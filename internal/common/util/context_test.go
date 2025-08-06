// @Author:冯铁城 [17615007230@163.com] 2025-08-06 17:32:00
package util

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetUsernameFromC(t *testing.T) {

	// 1.初始化Gin测试模式
	gin.SetMode(gin.TestMode)

	// 2.测试用例结构体
	tests := []struct {
		name      string
		setupCtx  func() *gin.Context
		want      string
		wantExist bool
	}{
		{
			name: "测试获取有效用户名",
			setupCtx: func() *gin.Context {
				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				c.Set("username", "testuser")
				return c
			},
			want:      "testuser",
			wantExist: true,
		},
		{
			name: "测试获取空用户名",
			setupCtx: func() *gin.Context {
				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				c.Set("username", "")
				return c
			},
			want:      "",
			wantExist: true,
		},
		{
			name: "测试获取特殊字符用户名",
			setupCtx: func() *gin.Context {
				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				c.Set("username", "user@test#123")
				return c
			},
			want:      "user@test#123",
			wantExist: true,
		},
		{
			name: "测试获取长用户名",
			setupCtx: func() *gin.Context {
				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				longUsername := "very_long_username_" + string(make([]rune, 100))
				c.Set("username", longUsername)
				return c
			},
			want:      "very_long_username_" + string(make([]rune, 100)),
			wantExist: true,
		},
		{
			name: "测试用户名不存在",
			setupCtx: func() *gin.Context {
				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				return c
			},
			want:      "",
			wantExist: false,
		},
		{
			name: "测试用户名类型错误",
			setupCtx: func() *gin.Context {
				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				c.Set("username", 123)
				return c
			},
			want:      "",
			wantExist: false,
		},
	}

	// 3.循环执行测试用例
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// 1.设置测试上下文
			c := tt.setupCtx()

			// 2.调用被测试函数
			got, exist := GetUsernameFromC(c)

			// 3.验证结果
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantExist, exist)
		})
	}
}

func TestGetEmailFromC(t *testing.T) {

	// 1.初始化Gin测试模式
	gin.SetMode(gin.TestMode)

	// 2.测试用例结构体
	tests := []struct {
		name      string
		setupCtx  func() *gin.Context
		want      string
		wantExist bool
	}{
		{
			name: "测试获取有效邮箱",
			setupCtx: func() *gin.Context {
				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				c.Set("email", "test@example.com")
				return c
			},
			want:      "test@example.com",
			wantExist: true,
		},
		{
			name: "测试获取空邮箱",
			setupCtx: func() *gin.Context {
				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				c.Set("email", "")
				return c
			},
			want:      "",
			wantExist: true,
		},
		{
			name: "测试获取复杂邮箱",
			setupCtx: func() *gin.Context {
				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				c.Set("email", "user+tag@sub.example-domain.com")
				return c
			},
			want:      "user+tag@sub.example-domain.com",
			wantExist: true,
		},
		{
			name: "测试邮箱不存在",
			setupCtx: func() *gin.Context {
				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				return c
			},
			want:      "",
			wantExist: false,
		},
		{
			name: "测试邮箱类型错误",
			setupCtx: func() *gin.Context {
				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				c.Set("email", 123.45)
				return c
			},
			want:      "",
			wantExist: false,
		},
	}

	// 3.循环执行测试用例
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// 1.设置测试上下文
			c := tt.setupCtx()

			// 2.调用被测试函数
			got, exist := GetEmailFromC(c)

			// 3.验证结果
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantExist, exist)
		})
	}
}

func TestGetStringFromC(t *testing.T) {

	// 1.初始化Gin测试模式
	gin.SetMode(gin.TestMode)

	// 2.测试用例结构体
	tests := []struct {
		name      string
		key       string
		setupCtx  func() *gin.Context
		want      string
		wantExist bool
	}{
		{
			name: "测试获取有效字符串",
			key:  "testKey",
			setupCtx: func() *gin.Context {
				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				c.Set("testKey", "testValue")
				return c
			},
			want:      "testValue",
			wantExist: true,
		},
		{
			name: "测试获取空字符串",
			key:  "emptyKey",
			setupCtx: func() *gin.Context {
				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				c.Set("emptyKey", "")
				return c
			},
			want:      "",
			wantExist: true,
		},
		{
			name: "测试键不存在",
			key:  "nonExistKey",
			setupCtx: func() *gin.Context {
				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				return c
			},
			want:      "",
			wantExist: false,
		},
		{
			name: "测试值类型错误",
			key:  "wrongTypeKey",
			setupCtx: func() *gin.Context {
				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				c.Set("wrongTypeKey", 123)
				return c
			},
			want:      "",
			wantExist: false,
		},
	}

	// 3.循环执行测试用例
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// 1.设置测试上下文
			c := tt.setupCtx()

			// 2.调用被测试函数
			got, exist := GetStringFromC(c, tt.key)

			// 3.验证结果
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantExist, exist)
		})
	}
}

func TestGetBoolFromC(t *testing.T) {

	// 1.初始化Gin测试模式
	gin.SetMode(gin.TestMode)

	// 2.测试用例结构体
	tests := []struct {
		name      string
		key       string
		setupCtx  func() *gin.Context
		want      bool
		wantExist bool
	}{
		{
			name: "测试获取true值",
			key:  "boolKey",
			setupCtx: func() *gin.Context {
				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				c.Set("boolKey", true)
				return c
			},
			want:      true,
			wantExist: true,
		},
		{
			name: "测试获取false值",
			key:  "boolKey",
			setupCtx: func() *gin.Context {
				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				c.Set("boolKey", false)
				return c
			},
			want:      false,
			wantExist: true,
		},
		{
			name: "测试键不存在",
			key:  "nonExistKey",
			setupCtx: func() *gin.Context {
				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				return c
			},
			want:      false,
			wantExist: false,
		},
		{
			name: "测试值类型错误",
			key:  "wrongTypeKey",
			setupCtx: func() *gin.Context {
				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				c.Set("wrongTypeKey", "true")
				return c
			},
			want:      false,
			wantExist: false,
		},
	}

	// 3.循环执行测试用例
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// 1.设置测试上下文
			c := tt.setupCtx()

			// 2.调用被测试函数
			got, exist := GetBoolFromC(c, tt.key)

			// 3.验证结果
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantExist, exist)
		})
	}
}

func TestGetIntFromC(t *testing.T) {

	// 1.初始化Gin测试模式
	gin.SetMode(gin.TestMode)

	// 2.测试用例结构体
	tests := []struct {
		name      string
		key       string
		setupCtx  func() *gin.Context
		want      int
		wantExist bool
	}{
		{
			name: "测试获取正整数",
			key:  "intKey",
			setupCtx: func() *gin.Context {
				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				c.Set("intKey", 123)
				return c
			},
			want:      123,
			wantExist: true,
		},
		{
			name: "测试获取零值",
			key:  "intKey",
			setupCtx: func() *gin.Context {
				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				c.Set("intKey", 0)
				return c
			},
			want:      0,
			wantExist: true,
		},
		{
			name: "测试获取负整数",
			key:  "intKey",
			setupCtx: func() *gin.Context {
				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				c.Set("intKey", -456)
				return c
			},
			want:      -456,
			wantExist: true,
		},
		{
			name: "测试键不存在",
			key:  "nonExistKey",
			setupCtx: func() *gin.Context {
				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				return c
			},
			want:      0,
			wantExist: false,
		},
		{
			name: "测试值类型错误",
			key:  "wrongTypeKey",
			setupCtx: func() *gin.Context {
				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				c.Set("wrongTypeKey", "123")
				return c
			},
			want:      0,
			wantExist: false,
		},
	}

	// 3.循环执行测试用例
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// 1.设置测试上下文
			c := tt.setupCtx()

			// 2.调用被测试函数
			got, exist := GetIntFromC(c, tt.key)

			// 3.验证结果
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantExist, exist)
		})
	}
}

func TestGetInt64FromC(t *testing.T) {

	// 1.初始化Gin测试模式
	gin.SetMode(gin.TestMode)

	// 2.测试用例结构体
	tests := []struct {
		name      string
		key       string
		setupCtx  func() *gin.Context
		want      int64
		wantExist bool
	}{
		{
			name: "测试获取正整数",
			key:  "int64Key",
			setupCtx: func() *gin.Context {
				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				c.Set("int64Key", int64(123456789))
				return c
			},
			want:      123456789,
			wantExist: true,
		},
		{
			name: "测试获取最大值",
			key:  "int64Key",
			setupCtx: func() *gin.Context {
				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				c.Set("int64Key", int64(9223372036854775807))
				return c
			},
			want:      9223372036854775807,
			wantExist: true,
		},
		{
			name: "测试获取最小值",
			key:  "int64Key",
			setupCtx: func() *gin.Context {
				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				c.Set("int64Key", int64(-9223372036854775808))
				return c
			},
			want:      -9223372036854775808,
			wantExist: true,
		},
		{
			name: "测试键不存在",
			key:  "nonExistKey",
			setupCtx: func() *gin.Context {
				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				return c
			},
			want:      0,
			wantExist: false,
		},
		{
			name: "测试值类型错误(int)",
			key:  "wrongTypeKey",
			setupCtx: func() *gin.Context {
				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				c.Set("wrongTypeKey", 123)
				return c
			},
			want:      0,
			wantExist: false,
		},
	}

	// 3.循环执行测试用例
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// 1.设置测试上下文
			c := tt.setupCtx()

			// 2.调用被测试函数
			got, exist := GetInt64FromC(c, tt.key)

			// 3.验证结果
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantExist, exist)
		})
	}
}

func TestGetFloat64FromC(t *testing.T) {

	// 1.初始化Gin测试模式
	gin.SetMode(gin.TestMode)

	// 2.测试用例结构体
	tests := []struct {
		name      string
		key       string
		setupCtx  func() *gin.Context
		want      float64
		wantExist bool
	}{
		{
			name: "测试获取正浮点数",
			key:  "floatKey",
			setupCtx: func() *gin.Context {
				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				c.Set("floatKey", 123.456)
				return c
			},
			want:      123.456,
			wantExist: true,
		},
		{
			name: "测试获取零值",
			key:  "floatKey",
			setupCtx: func() *gin.Context {
				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				c.Set("floatKey", 0.0)
				return c
			},
			want:      0.0,
			wantExist: true,
		},
		{
			name: "测试获取负浮点数",
			key:  "floatKey",
			setupCtx: func() *gin.Context {
				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				c.Set("floatKey", -789.123)
				return c
			},
			want:      -789.123,
			wantExist: true,
		},
		{
			name: "测试键不存在",
			key:  "nonExistKey",
			setupCtx: func() *gin.Context {
				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				return c
			},
			want:      0.0,
			wantExist: false,
		},
		{
			name: "测试值类型错误",
			key:  "wrongTypeKey",
			setupCtx: func() *gin.Context {
				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				c.Set("wrongTypeKey", "123.456")
				return c
			},
			want:      0.0,
			wantExist: false,
		},
	}

	// 3.循环执行测试用例
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// 1.设置测试上下文
			c := tt.setupCtx()

			// 2.调用被测试函数
			got, exist := GetFloat64FromC(c, tt.key)

			// 3.验证结果
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantExist, exist)
		})
	}
}

func TestGetSliceFromC(t *testing.T) {

	// 1.初始化Gin测试模式
	gin.SetMode(gin.TestMode)

	// 2.测试用例结构体
	tests := []struct {
		name      string
		key       string
		setupCtx  func() *gin.Context
		want      []string
		wantExist bool
	}{
		{
			name: "测试获取字符串切片",
			key:  "sliceKey",
			setupCtx: func() *gin.Context {
				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				c.Set("sliceKey", []string{"a", "b", "c"})
				return c
			},
			want:      []string{"a", "b", "c"},
			wantExist: true,
		},
		{
			name: "测试获取空切片",
			key:  "sliceKey",
			setupCtx: func() *gin.Context {
				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				c.Set("sliceKey", []string{})
				return c
			},
			want:      []string{},
			wantExist: true,
		},
		{
			name: "测试键不存在",
			key:  "nonExistKey",
			setupCtx: func() *gin.Context {
				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				return c
			},
			want:      nil,
			wantExist: false,
		},
		{
			name: "测试值类型错误",
			key:  "wrongTypeKey",
			setupCtx: func() *gin.Context {
				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				c.Set("wrongTypeKey", "not a slice")
				return c
			},
			want:      nil,
			wantExist: false,
		},
	}

	// 3.循环执行测试用例
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// 1.设置测试上下文
			c := tt.setupCtx()

			// 2.调用被测试函数
			got, exist := GetSliceFromC[string](c, tt.key)

			// 3.验证结果
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantExist, exist)
		})
	}
}

func TestGetMapFromC(t *testing.T) {

	// 1.初始化Gin测试模式
	gin.SetMode(gin.TestMode)

	// 2.测试用例结构体
	tests := []struct {
		name      string
		key       string
		setupCtx  func() *gin.Context
		want      map[string]int
		wantExist bool
	}{
		{
			name: "测试获取字符串到整数的映射",
			key:  "mapKey",
			setupCtx: func() *gin.Context {
				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				c.Set("mapKey", map[string]int{"a": 1, "b": 2})
				return c
			},
			want:      map[string]int{"a": 1, "b": 2},
			wantExist: true,
		},
		{
			name: "测试获取空映射",
			key:  "mapKey",
			setupCtx: func() *gin.Context {
				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				c.Set("mapKey", map[string]int{})
				return c
			},
			want:      map[string]int{},
			wantExist: true,
		},
		{
			name: "测试键不存在",
			key:  "nonExistKey",
			setupCtx: func() *gin.Context {
				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				return c
			},
			want:      nil,
			wantExist: false,
		},
		{
			name: "测试值类型错误",
			key:  "wrongTypeKey",
			setupCtx: func() *gin.Context {
				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				c.Set("wrongTypeKey", "not a map")
				return c
			},
			want:      nil,
			wantExist: false,
		},
	}

	// 3.循环执行测试用例
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// 1.设置测试上下文
			c := tt.setupCtx()

			// 2.调用被测试函数
			got, exist := GetMapFromC[string, int](c, tt.key)

			// 3.验证结果
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantExist, exist)
		})
	}
}

func TestGetValueFromC(t *testing.T) {

	// 1.初始化Gin测试模式
	gin.SetMode(gin.TestMode)

	// 2.测试不同类型的值
	t.Run("测试获取不同类型的值", func(t *testing.T) {

		// 1.创建测试上下文
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		// 2.设置不同类型的值
		c.Set("stringVal", "test")
		c.Set("intVal", 123)
		c.Set("boolVal", true)
		c.Set("floatVal", 3.14)

		// 3.测试获取字符串
		strVal, exist := GetValueFromC[string](c, "stringVal")
		assert.True(t, exist)
		assert.Equal(t, "test", strVal)

		// 4.测试获取整数
		intVal, exist := GetValueFromC[int](c, "intVal")
		assert.True(t, exist)
		assert.Equal(t, 123, intVal)

		// 5.测试获取布尔值
		boolVal, exist := GetValueFromC[bool](c, "boolVal")
		assert.True(t, exist)
		assert.Equal(t, true, boolVal)

		// 6.测试获取浮点数
		floatVal, exist := GetValueFromC[float64](c, "floatVal")
		assert.True(t, exist)
		assert.Equal(t, 3.14, floatVal)

		// 7.测试获取不存在的键
		_, exist = GetValueFromC[string](c, "nonExist")
		assert.False(t, exist)

		// 8.测试类型转换失败
		_, exist = GetValueFromC[int](c, "stringVal")
		assert.False(t, exist)
	})
}

func TestAddKVToC(t *testing.T) {

	// 1.初始化Gin测试模式
	gin.SetMode(gin.TestMode)

	// 2.测试添加键值对
	t.Run("测试添加不同类型的键值对", func(t *testing.T) {

		// 1.创建测试上下文
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		// 2.添加不同类型的值
		AddKVToC(c, "stringKey", "stringValue")
		AddKVToC(c, "intKey", 456)
		AddKVToC(c, "boolKey", false)
		AddKVToC(c, "sliceKey", []string{"x", "y", "z"})

		// 3.验证添加的值
		strVal, exist := GetStringFromC(c, "stringKey")
		assert.True(t, exist)
		assert.Equal(t, "stringValue", strVal)

		intVal, exist := GetIntFromC(c, "intKey")
		assert.True(t, exist)
		assert.Equal(t, 456, intVal)

		boolVal, exist := GetBoolFromC(c, "boolKey")
		assert.True(t, exist)
		assert.Equal(t, false, boolVal)

		sliceVal, exist := GetSliceFromC[string](c, "sliceKey")
		assert.True(t, exist)
		assert.Equal(t, []string{"x", "y", "z"}, sliceVal)
	})

	// 3.测试覆盖已存在的键
	t.Run("测试覆盖已存在的键", func(t *testing.T) {

		// 1.创建测试上下文
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		// 2.添加初始值
		AddKVToC(c, "testKey", "initialValue")
		val, exist := GetStringFromC(c, "testKey")
		assert.True(t, exist)
		assert.Equal(t, "initialValue", val)

		// 3.覆盖值
		AddKVToC(c, "testKey", "newValue")
		val, exist = GetStringFromC(c, "testKey")
		assert.True(t, exist)
		assert.Equal(t, "newValue", val)
	})
}

func TestContextIntegration(t *testing.T) {

	// 1.初始化Gin测试模式
	gin.SetMode(gin.TestMode)

	// 2.集成测试：模拟完整的上下文操作流程
	t.Run("测试完整的上下文操作流程", func(t *testing.T) {

		// 1.创建测试上下文
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		// 2.模拟中间件设置用户信息
		AddKVToC(c, "username", "testuser")
		AddKVToC(c, "email", "test@example.com")
		AddKVToC(c, "userId", int64(12345))
		AddKVToC(c, "isAdmin", true)
		AddKVToC(c, "permissions", []string{"read", "write", "delete"})

		// 3.验证获取用户信息
		username, exist := GetUsernameFromC(c)
		assert.True(t, exist)
		assert.Equal(t, "testuser", username)

		email, exist := GetEmailFromC(c)
		assert.True(t, exist)
		assert.Equal(t, "test@example.com", email)

		userId, exist := GetInt64FromC(c, "userId")
		assert.True(t, exist)
		assert.Equal(t, int64(12345), userId)

		isAdmin, exist := GetBoolFromC(c, "isAdmin")
		assert.True(t, exist)
		assert.Equal(t, true, isAdmin)

		permissions, exist := GetSliceFromC[string](c, "permissions")
		assert.True(t, exist)
		assert.Equal(t, []string{"read", "write", "delete"}, permissions)
	})
}

func TestContextEdgeCases(t *testing.T) {

	// 1.初始化Gin测试模式
	gin.SetMode(gin.TestMode)

	// 2.测试边界情况
	testCases := []struct {
		name        string
		key         string
		value       interface{}
		expectExist bool
	}{
		{
			name:        "测试nil值",
			key:         "nilKey",
			value:       nil,
			expectExist: true,
		},
		{
			name:        "测试空字符串键",
			key:         "",
			value:       "emptyKeyValue",
			expectExist: true,
		},
		{
			name:        "测试特殊字符键",
			key:         "key!@#$%^&*()",
			value:       "specialKeyValue",
			expectExist: true,
		},
		{
			name:        "测试长键名",
			key:         string(make([]rune, 1000)),
			value:       "longKeyValue",
			expectExist: true,
		},
		{
			name:        "测试结构体值",
			key:         "structKey",
			value:       struct{ Name string }{Name: "test"},
			expectExist: true,
		},
		{
			name:        "测试接口值",
			key:         "interfaceKey",
			value:       interface{}("interfaceValue"),
			expectExist: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			// 1.创建测试上下文
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			// 2.添加测试值
			AddKVToC(c, tc.key, tc.value)

			// 3.验证值是否存在
			val, exist := c.Get(tc.key)
			assert.Equal(t, tc.expectExist, exist)
			if tc.expectExist {
				assert.Equal(t, tc.value, val)
			}
		})
	}
}

func TestContextConcurrency(t *testing.T) {

	// 1.初始化Gin测试模式
	gin.SetMode(gin.TestMode)

	// 2.并发测试
	t.Run("测试并发访问Context", func(t *testing.T) {

		// 1.创建测试上下文
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		// 2.并发添加和读取数据
		const numGoroutines = 100
		done := make(chan bool, numGoroutines)

		// 3.启动并发测试
		for i := 0; i < numGoroutines; i++ {
			go func(id int) {
				defer func() { done <- true }()

				// 添加数据
				key := "key" + string(rune(id))
				value := "value" + string(rune(id))
				AddKVToC(c, key, value)

				// 读取数据
				retrievedValue, exist := GetStringFromC(c, key)
				assert.True(t, exist)
				assert.Equal(t, value, retrievedValue)
			}(i)
		}

		// 4.等待所有goroutine完成
		for i := 0; i < numGoroutines; i++ {
			<-done
		}
	})
}

func TestContextPerformance(t *testing.T) {

	// 1.初始化Gin测试模式
	gin.SetMode(gin.TestMode)

	// 2.性能测试
	t.Run("测试Context性能", func(t *testing.T) {

		// 1.创建测试上下文
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		// 2.预设数据
		AddKVToC(c, "username", "testuser")
		AddKVToC(c, "email", "test@example.com")
		AddKVToC(c, "userId", int64(12345))

		// 3.执行大量读取操作
		const iterations = 10000
		for i := 0; i < iterations; i++ {
			username, exist := GetUsernameFromC(c)
			assert.True(t, exist)
			assert.Equal(t, "testuser", username)

			email, exist := GetEmailFromC(c)
			assert.True(t, exist)
			assert.Equal(t, "test@example.com", email)

			userId, exist := GetInt64FromC(c, "userId")
			assert.True(t, exist)
			assert.Equal(t, int64(12345), userId)
		}
	})
}

func TestGenericFunctions(t *testing.T) {

	// 1.初始化Gin测试模式
	gin.SetMode(gin.TestMode)

	// 2.测试泛型函数
	t.Run("测试泛型函数的类型安全", func(t *testing.T) {

		// 1.创建测试上下文
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		// 2.测试不同类型的泛型函数
		// 测试整数切片
		intSlice := []int{1, 2, 3, 4, 5}
		AddKVToC(c, "intSlice", intSlice)
		retrievedIntSlice, exist := GetSliceFromC[int](c, "intSlice")
		assert.True(t, exist)
		assert.Equal(t, intSlice, retrievedIntSlice)

		// 测试字符串到布尔值的映射
		boolMap := map[string]bool{"active": true, "deleted": false}
		AddKVToC(c, "boolMap", boolMap)
		retrievedBoolMap, exist := GetMapFromC[string, bool](c, "boolMap")
		assert.True(t, exist)
		assert.Equal(t, boolMap, retrievedBoolMap)

		// 测试自定义结构体切片
		type User struct {
			ID   int
			Name string
		}
		userSlice := []User{{ID: 1, Name: "Alice"}, {ID: 2, Name: "Bob"}}
		AddKVToC(c, "userSlice", userSlice)
		retrievedUserSlice, exist := GetSliceFromC[User](c, "userSlice")
		assert.True(t, exist)
		assert.Equal(t, userSlice, retrievedUserSlice)

		// 测试整数到字符串的映射
		intStringMap := map[int]string{1: "one", 2: "two", 3: "three"}
		AddKVToC(c, "intStringMap", intStringMap)
		retrievedIntStringMap, exist := GetMapFromC[int, string](c, "intStringMap")
		assert.True(t, exist)
		assert.Equal(t, intStringMap, retrievedIntStringMap)
	})

	// 3.测试泛型函数的类型转换失败
	t.Run("测试泛型函数类型转换失败", func(t *testing.T) {

		// 1.创建测试上下文
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		// 2.设置错误类型的值
		AddKVToC(c, "wrongType", "not a slice")

		// 3.尝试获取切片类型，应该失败
		_, exist := GetSliceFromC[int](c, "wrongType")
		assert.False(t, exist)

		// 4.尝试获取映射类型，应该失败
		_, exist = GetMapFromC[string, int](c, "wrongType")
		assert.False(t, exist)
	})
}
