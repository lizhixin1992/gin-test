package commons

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type returnData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SetResponse(context *gin.Context, code int, message string, data interface{}) *gin.Context {
	context.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  message,
		"data": data,
	})
	return context
}

func SetResponseSuccessData(context *gin.Context, data interface{}) *gin.Context {
	context.JSON(http.StatusOK, gin.H{
		"code": "0",
		"msg":  "success",
		"data": data,
	})
	return context
}

func SetResponseSuccess(context *gin.Context) *gin.Context {
	context.JSON(http.StatusOK, gin.H{
		"code": "0",
		"msg":  "success",
	})
	return context
}

func SetResponseOperateFail(context *gin.Context) *gin.Context {
	context.JSON(http.StatusOK, gin.H{
		"code": "1001",
		"msg":  "method error",
	})
	return context
}

func SetResponseDataFail(context *gin.Context) *gin.Context {
	context.JSON(http.StatusOK, gin.H{
		"code": "1002",
		"msg":  "data error",
	})
	return context
}
