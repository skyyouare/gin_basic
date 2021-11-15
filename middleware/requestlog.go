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

	// 日志
	logger.Infow("gin-request-in",
		"uri", c.Request.RequestURI,
		"method", c.Request.Method,
		"args", c.Request.PostForm,
		"body", string(bodyBytes),
		"ip", c.ClientIP(),
	)
}

// RequestOutLog 请求输出日志
func RequestOutLog(c *gin.Context) {
	endExecTime := time.Now()
	response, _ := c.Get("response")
	st, _ := c.Get("startExecTime")

	startExecTime, _ := st.(time.Time)

	// 日志
	logger.Infow("gin-request-out",
		"uri", c.Request.RequestURI,
		"method", c.Request.Method,
		"args", c.Request.PostForm,
		"ip", c.ClientIP(),
		"response", response,
		"exec_time", endExecTime.Sub(startExecTime).Seconds(),
	)
}

// RequestLog 日志中间件调用
func RequestLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(c.Errors) > 0 {
			// Append error field if this is an erroneous request.
			for _, e := range c.Errors.Errors() {
				logger.Errorf(e)
			}
		} else {
			//弃用，合成一条日志
			// RequestInLog(c)
			// defer RequestOutLog(c)
			start := time.Now()
			bodyBytes, _ := ioutil.ReadAll(c.Request.Body)
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes)) // Write body back
			c.Next()
			response, _ := c.Get("response")
			end := time.Now()
			exec_time := end.Sub(start)
			logger.Infow("gin-request",
				"uri", c.Request.RequestURI,
				"method", c.Request.Method,
				"args", c.Request.PostForm,
				"body", string(bodyBytes),
				"ip", c.ClientIP(),
				"response", response,
				"exec_time", exec_time,
			)
		}
		c.Next()
	}
}
