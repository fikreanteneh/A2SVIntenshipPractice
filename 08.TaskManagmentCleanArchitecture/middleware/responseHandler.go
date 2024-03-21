package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Response struct {
	StatusCode  int         `json:"statusCode"`
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Error     string      `json:"error,omitempty"`
}

func SuccessResponseHandler(c *gin.Context, statusCode int, message string, data interface{} ) {
	fmt.Println("SuccessResponseHandler ======= ", data)
	response := Response{
		StatusCode: statusCode,
		Success: true,
		Message: message,
		Data: data,
	}
	c.JSON(statusCode, response)
}