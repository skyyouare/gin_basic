package controller

import (
	"fmt"
	"gin_basic/middleware"
	"gin_basic/model"
	"gin_basic/pkg/gorm"

	"github.com/gin-gonic/gin"
)

// TestController 控制器
type TestController struct {
}

// TestRegister 路由注册
func TestRegister(router *gin.RouterGroup) {
	controller := new(TestController)
	router.GET("/test/test", controller.test)
}

// 测试
func (t *TestController) test(c *gin.Context) {
	// blog_auth
	var authlist []model.Auth
	gorm.Conn.Find(&authlist)
	fmt.Println(authlist)
	middleware.ResponseSuccess(c, "test")
}
