package middleware

import (
	"errors"
	"fmt"
	"gin_basic/pkg/logger"
	"gin_basic/pkg/setting"
	"net/http/httputil"
	"runtime/debug"
	"time"

	"github.com/gin-gonic/gin"
)

// RecoveryMiddleware 捕获所有panic，并且返回错误信息
func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				logger.Errorw(
					"recovery-middleware",
					"time", time.Now(),
					"error", err,
					"request", string(httpRequest),
					"stack", string(debug.Stack()),
				)
				// 接口返回
				if setting.AppSetting.DebugMode != "debug" {
					ResponseError(c, 1004, errors.New("内部错误"))
				} else {
					ResponseError(c, 1004, errors.New(fmt.Sprint(err)))
				}
				return
			}
		}()
		c.Next()
	}
}
