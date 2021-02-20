package router

import (
	"gin_basic/pkg/ginzap"
	"gin_basic/routes"
	"time"

	"github.com/gin-gonic/gin"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(ginzap.Ginzap(time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(true))
	routes.GetRoutes(r)
	return r
}
