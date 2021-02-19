package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//ResponseCode 状态码
type ResponseCode int

const (
	//SuccessCode 成功状态码
	SuccessCode ResponseCode = 1000
	//ErrorCode 失败状态码
	ErrorCode ResponseCode = 1001
)

type successResponse struct {
	Code ResponseCode
	Data interface{}
}

type errorResponse struct {
	Code ResponseCode
	Msg  string
}

//ResponseError 成功状态返回方法
func ResponseError(c *gin.Context, code ResponseCode, err error) {
	fmt.Print(err)
	resp := &errorResponse{Code: code, Msg: err.Error()}
	c.JSON(http.StatusOK, resp)
}

//ResponseSuccess 失败状态返回方法
func ResponseSuccess(c *gin.Context, data interface{}) {
	resp := &successResponse{Code: SuccessCode, Data: data}
	c.JSON(http.StatusOK, resp)
}
