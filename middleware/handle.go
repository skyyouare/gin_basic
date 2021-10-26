package middleware

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func HandleNotFound(c *gin.Context) {
	ResponseError(c, NotFoundCode, errors.New("page not found: "+c.Request.RequestURI))
}
