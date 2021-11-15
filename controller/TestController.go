package controller

import (
	"fmt"
	"gin_basic/middleware"
	"gin_basic/pkg/gorm"
	"time"

	"github.com/gin-gonic/gin"
)

var db = gorm.Conn

// TestController 控制器
type TestController struct {
}

// TestRegister 路由注册
func TestRegister(router *gin.RouterGroup) {
	controller := new(TestController)
	router.GET("/test/test", controller.test)
}

type Auth struct {
	ID        uint      `gorm:"column:id;primary_key"` //primary_key:设置主键
	Username  string    `gorm:"column:username;type:varchar(100)"`
	Password  string    `gorm:"column:password;type:varchar(100)"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

// 测试
func (t *TestController) test(c *gin.Context) {
	// middleware.ResponseError(c, middleware.ErrorCode, errors.New("测试"))
	// blog_auth
	var auth Auth
	db.First(&auth)
	fmt.Println(auth, auth.ID, auth.Username, auth.Password)
	middleware.ResponseSuccess(c, "test")
}
