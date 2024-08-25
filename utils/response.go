package utils

import "github.com/gin-gonic/gin"

func ErrorResponseGeneric(ctx *gin.Context, statusCode int, errorMsg string) {
	ctx.JSON(statusCode, gin.H{
		"code":      statusCode,
		"data":      nil,
		"error":     errorMsg,
		"requestID": ctx.GetString("requestID"),
	})
}

func SuccessResponseGeneric(ctx *gin.Context, statusCode int, data any) {
	ctx.JSON(statusCode, gin.H{
		"code":      statusCode,
		"data":      data,
		"error":     nil,
		"requestID": ctx.GetString("requestID"),
	})
}
