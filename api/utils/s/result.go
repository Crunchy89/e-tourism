package s

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func AbortWithStatus(c *gin.Context, status int, err error) {
	result := NewResultError(err)
	result.Code = status
	println(result.ErrorMessage)
	c.AbortWithStatusJSON(status, result)
}

func AbortWithMessageStatus(c *gin.Context, status int, message string) {
	AbortWithStatus(c, status, errors.New(message))
}

func Auto(c *gin.Context, data interface{}, err error) {
	if err != nil {
		println(err.Error())
		Abort(c, err)
	} else {
		ResultWithData(c, data)
	}
}

func Abort(c *gin.Context, err error) {
	result := NewResultError(err)
	println(result.ErrorMessage)
	c.AbortWithStatusJSON(result.Code, result)
}

func AbortWithMessage(c *gin.Context, message string) {
	result := NewResultError(errors.New(message))
	println(result.ErrorMessage)
	c.AbortWithStatusJSON(result.Code, result)
}

func ResultWithData(c *gin.Context, data interface{}) {
	result := NewResultData(data)
	c.AbortWithStatusJSON(result.Code, result)
}
