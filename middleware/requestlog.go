package middleware

import (
	"bytes"
	"gin_basic/pkg/logger"
	"io/ioutil"
	"time"

	"github.com/gin-gonic/gin"
)

// RequestInLog 请求进入日志
func RequestInLog(c *gin.Context) {
	c.Set("startExecTime", time.Now())

	bodyBytes, _ := ioutil.ReadAll(c.Request.Body)
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes)) // Write body back

	//日志
	logger.Infow("gin-request-in", "uri", c.Request.RequestURI, "method", c.Request.Method, "args", c.Request.PostForm, "body", string(bodyBytes), "ip", c.ClientIP())
}

// RequestOutLog 请求输出日志
func RequestOutLog(c *gin.Context) {
	endExecTime := time.Now()
	response, _ := c.Get("response")
	st, _ := c.Get("startExecTime")

	startExecTime, _ := st.(time.Time)

	//日志
	logger.Infow("gin-request-out", "uri", c.Request.RequestURI, "method", c.Request.Method, "args", c.Request.PostForm, "ip", c.ClientIP(), "response", response, "proc_time", endExecTime.Sub(startExecTime).Seconds())
}

// RequestLog 日志中间件调用
func RequestLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		RequestInLog(c)
		defer RequestOutLog(c)
		c.Next()
	}
}
