package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResponseCode int

const (
	SuccessCode ResponseCode = 1000
	ErrorCode   ResponseCode = 1001
)

type successResponse struct {
	Code ResponseCode
	Data interface{}
}

type errorResponse struct {
	Code ResponseCode
	Msg  string
}

func ResponseError(c *gin.Context, code ResponseCode, err error) {
	fmt.Print(err)
	resp := &errorResponse{Code: code, Msg: err.Error()}
	c.JSON(http.StatusOK, resp)
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	resp := &successResponse{Code: SuccessCode, Data: data}
	c.JSON(http.StatusOK, resp)
}
