package dto

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResponseDTO struct {
	ErrorCode int         `json:"errorCode"`
	Result    interface{} `json:"result"`
}

func Success(c *gin.Context, result interface{}) {
	c.JSON(http.StatusOK, ResponseDTO{
		ErrorCode: 0,
		Result:    result,
	})
}

func Fail(c *gin.Context, errorcode int) {
	c.JSON(http.StatusOK, ResponseDTO{
		ErrorCode: errorcode,
	})
}
