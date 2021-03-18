package app

import (
	"log"
	"os"
)

var ServerMode string // 用来获取最终需要的os.Args[1]
var SetupType string
var ConfPath string

// Setup 设置配置
func Setup() {
	for k, v := range os.Args {
		if k == 1 { // 假设需要获取os.Args[k], k = 1
			ServerMode = v
		}
	}
	if ServerMode == "" { // 不为空则表示os.Arg[1]存在
		log.Fatalf("请增加启动模式")
	}
	// 初始化配置文件
	switch ServerMode {
	case "dev":
		// 加载调试模式配置
		SetupType = "开发模式"
		ConfPath = "./conf/dev/app.toml"
	case "prod":
		// 加载正式环境配置
		SetupType = "正式环境"
		ConfPath = "./conf/prod/app.toml"
	case "test":
		// 加载测试环境配置
		SetupType = "测试环境"
		ConfPath = "./conf/test/app.toml"
	default:
		log.Println(" 启动模式错误:dev/test/prod", ServerMode)
	}
}
