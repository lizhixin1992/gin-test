package commons

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type returnData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func SetResponse(context *gin.Context, code int, message string, data interface{}) *gin.Context {
	ret := returnData{Code: code, Message: message, Data: data}
	context.JSON(http.StatusOK, ret)
	return context
}

func SetResponseSuccessData(context *gin.Context, data interface{}) *gin.Context {
	ret := returnData{Code: 0, Message: "success", Data: data}
	context.JSON(http.StatusOK, ret)
	return context
}

func SetResponseSuccess(context *gin.Context) *gin.Context {
	ret := returnData{Code: 0, Message: "success"}
	context.JSON(http.StatusOK, ret)
	return context
}

func SetResponseOperateFail(context *gin.Context) *gin.Context {
	ret := returnData{Code: 1001, Message: "method error"}
	context.JSON(http.StatusOK, ret)
	return context
}

func SetResponseDataFail(context *gin.Context) *gin.Context {
	ret := returnData{Code: 1002, Message: "data error"}
	context.JSON(http.StatusOK, ret)
	return context
}
