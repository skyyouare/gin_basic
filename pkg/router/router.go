package router

import (
	"github.com/gin-gonic/gin"
	"gin_basic/routes"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	routes.GetRoutes(r)
	return r
}
