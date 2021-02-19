package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"gin_basic/pkg/router"
	"gin_basic/pkg/setting"
	"log"
	"net/http"
	"time"
)

var (
	serv *http.Server
)

func HttpServRun() {
	gin.SetMode(setting.AppSetting.DebugMode)
	r := router.InitRouter()
	serv = &http.Server{
		Addr:    setting.ServerSetting.HttpPort,
		Handler: r,
	}
	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		log.Printf(" [INFO] HttpServerRun:%s\n", setting.AppSetting.DebugMode)
		if err := serv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf(" [ERROR] HttpServerRun:%s err:%v\n", setting.ServerSetting.HttpPort, err)
		}
	}()
}

func HttpServStop() {
	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := serv.Shutdown(ctx); err != nil {
		log.Fatalf(" [ERROR] HttpServerStop err:%v\n", err)
	}

	log.Printf(" [INFO] HttpServerStop stopped\n")
}
