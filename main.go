package main

import (
	"gin_basic/pkg/app"
	"gin_basic/pkg/gorm2"
	"gin_basic/pkg/logger"
	"gin_basic/pkg/server"
	"gin_basic/pkg/setting"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	// 启动模式设置等
	app.Setup()
	// 配置设置
	setting.Setup(app.ConfPath)
	// 初始化日志
	logger.Setup()
	logger.Infof("传递模式为%s，加载%s配置", app.ServerMode, app.SetupType)
	// 初始化db
	// db.Setup()
	// 初始化gorm
	gorm2.Setup()
}

func main() {
	// 关闭db，redis连接等
	defer func() {
		logger.Infof("数据库连接关闭")
		// db.Close()
		gorm2.Close()
	}()
	// 启动http服务
	server.HTTPServRun()
	// 优雅关闭服务
	quit := make(chan os.Signal, 1)
	// signal输入信号转发到c syscall.SIGKILL,
	signal.Notify(quit, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Errorf("Shutting down server...")
	// 关闭http服务
	server.HTTPServStop()
}
