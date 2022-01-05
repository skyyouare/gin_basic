package router

import (
	"gin_basic/middleware"
	"gin_basic/routes"

	"github.com/gin-gonic/gin"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	// r.Use(gin.Recovery())
	// ginzap弃用，使用middleware.RequestLog() middleware.RecoveryMiddleware()代替
	// r.Use(ginzap.Ginzap(time.RFC3339, true))
	// r.Use(ginzap.RecoveryWithZap(true))
	r.NoRoute(middleware.HandleNotFound)
	r.NoMethod(middleware.HandleNotFound)
	r.Use(middleware.RequestLog())
	r.Use(middleware.RecoveryMiddleware())
	r.StaticFile("/favicon.ico", "./favicon.ico")
	routes.GetRoutes(r)
	return r
}
