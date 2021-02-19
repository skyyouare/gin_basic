package main

import (
	"gin_basic/pkg/db"
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
		log.Fatalln("请增加启动模式")
	}
	//根据参数选择配置文件
	switch serverMode {
	case "dev":
		// 加载调试模式配置
		log.Printf(" [INFO] 传递模式为%s，加载开发模式配置\n", serverMode)
		setting.Setup("./conf/dev/app.toml")
	case "prod":
		// 加载正式环境配置
		log.Printf(" [INFO] 传递模式为%s，加载正式环境配置\n", serverMode)
		setting.Setup("./conf/prod/app.toml")
	case "test":
		// 加载测试环境配置
		log.Printf(" [INFO] 传递模式为%s，加载测试环境配置\n", serverMode)
		setting.Setup("./conf/test/app.toml")
	default:
		log.Fatalln("启动模式错误:dev/test/prod")
	}
	//初始化db
	db.Setup()
}

func main() {
	//httpserver run
	server.HttpServRun()
	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	//signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")
	//httpserver stop
	server.HttpServStop()
}
