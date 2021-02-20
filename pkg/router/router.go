package router

import (
	"gin_basic/routes"

	"github.com/gin-gonic/gin"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	// r.Use(ginzap.Ginzap(time.RFC3339, true))
	// r.Use(ginzap.RecoveryWithZap(true))
	routes.GetRoutes(r)
	return r
}
