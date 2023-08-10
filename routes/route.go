package routes

import (
	"gin_basic/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetRoutes will create our routes of our entire application
// this way every group of routes can be defined in their own file
// so this one won't be so messy
func GetRoutes(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome Gin Server1234")
	})
	router.GET("/panic", func(c *gin.Context) {
		panic("An unexpected error happen!")
	})
	v1 := router.Group("/v1")

	v1.Use(
	// 弃用 在InitRouter中使用中间件
	// middleware.RecoveryMiddleware(),
	// middleware.RequestLog(),
	)
	{
		controller.ReportRegister(v1)
		controller.TestRegister(v1)
	}
}
