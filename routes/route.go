package routes

import (
	"github.com/gin-gonic/gin"
	"gin_basic/controller"
	"net/http"
)

// getRoutes will create our routes of our entire application
// this way every group of routes can be defined in their own file
// so this one won't be so messy
func GetRoutes(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome Gin Server14")
	})
	v1 := router.Group("/v1")

	v1.Use(
	//middleware.RecoveryMiddleware(),
	)
	{
		controller.ReportRegister(v1)
	}
}
