// Package server @Author:冯铁城 [17615007230@163.com] 2025-07-30 16:20:53
package server

import (
	"context"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// router 定义核心路由
var router *gin.Engine

// StartServer 启动并优雅关闭服务器
func StartServer() {

	//1.创建 Gin 实例
	router = gin.New()

	//2.使用全局异常处理器
	router.Use(errorHandler())

	//3.使用 Logger 和 Recovery 中间件
	router.Use(gin.Logger(), gin.Recovery())

	//4.初始化路由组
	initRouter()

	//5.创建http.Server实例
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	//6.创建协程启动 HTTP 服务
	go func() {
		log.Printf("server start success，listen addr：%s", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server start error: %v", err)
		}
	}()

	//7.等待中断信号并优雅关闭服务器
	waitForShutdown(srv)
}

// waitForShutdown 优雅关闭服务器
func waitForShutdown(srv *http.Server) {

	//1.创建一个缓冲区为1的channel，用于接收 OS 信号（中断/终止）
	quit := make(chan os.Signal, 1)

	//2.注册要监听的信号：SIGINT（Ctrl+C）和 SIGTERM（容器/系统终止）
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	//3.阻塞等待接收信号（接收到信号后继续执行）
	<-quit
	log.Println("receive stop sign, closing server...")

	//4.创建一个带有超时的上下文，最多等待5秒处理完未完成的请求
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//5.执行优雅关闭操作：拒绝新连接，等待现有请求完成或超时
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("server close fail: %v", err)
	}

	//6.所有资源释放后，输出日志
	log.Println("server close success")
}
