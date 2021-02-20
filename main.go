package main

import (
	"gin_basic/pkg/logger"
	"gin_basic/pkg/server"
	"gin_basic/pkg/setting"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	var serverMode string // 用来获取最终需要的os.Args[1]
	for k, v := range os.Args {
		if k == 1 { // 假设需要获取os.Args[k], k = 1
			serverMode = v
		}
	}
	if serverMode == "" { // 不为空则表示os.Arg[1]存在
		log.Fatalf("请增加启动模式")
	}
	//初始化配置文件
	var setupType string
	switch serverMode {
	case "dev":
		// 加载调试模式配置
		setupType = "开发模式"
		setting.Setup("./conf/dev/app.toml")
	case "prod":
		// 加载正式环境配置
		setupType = "正式环境"
		setting.Setup("./conf/prod/app.toml")
	case "test":
		// 加载测试环境配置
		setupType = "测试环境"
		setting.Setup("./conf/test/app.toml")
	default:
		log.Println(" 启动模式错误:dev/test/prod", serverMode)
	}
	//初始化日志
	logger.Setup()
	logger.Infof("传递模式为%s，加载%s配置", serverMode, setupType)
	//初始化db
	// db.Setup()
}

func main() {
	//httpserver run
	server.HTTPServRun()
	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	//signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Infof("Shutting down server...")
	//httpserver stop
	server.HTTPServStop()
}
