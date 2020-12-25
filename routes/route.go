package routes

import (
	"github.com/gin-gonic/gin"
)

func InitRoute() *gin.Engine {
	//router := gin.Default()
	////普通的路由请求
	//router.GET("/", func(context *gin.Context) {
	//	context.JSON(http.StatusOK, gin.H{
	//		"code": "0",
	//		"msg":  "success",
	//	})
	//})
	//
	////路由重定向
	//router.GET("/user/id", func(context *gin.Context) {
	//	context.Request.URL.Path = "/user/list"
	//	router.HandleContext(context)
	//})
	//router.GET("/user/list", func(context *gin.Context) {
	//	context.JSON(http.StatusOK, gin.H{
	//		"code": "2",
	//		"msg":  "success",
	//	})
	//})
	//
	//// 此 handler 将匹配 /user/1221 但不会匹配 /user/ 或者 /user
	//router.GET("/user/condition/:id", func(context *gin.Context) {
	//	id := context.Param("id")
	//	context.JSON(http.StatusOK, gin.H{
	//		"code": "3",
	//		"msg":  "success",
	//		"data": id,
	//	})
	//})
	//
	//// 此 handler 将匹配 /user/1221/ 和 /user/1221/send
	//// 如果没有其他路由匹配 /user/1221，它将重定向到 /user/1221/
	//router.GET("/user/condition/:id/*action", func(context *gin.Context) {
	//	name := context.Param("name")
	//	action := context.Param("action")
	//	data := name + "is" + action
	//	context.JSON(http.StatusOK, gin.H{
	//		"code": "4",
	//		"msg":  "success",
	//		"data": data,
	//	})
	//})
	//
	//router.POST("/user/search", func(context *gin.Context) {
	//	context.JSON(http.StatusOK, gin.H{
	//		"code": "10",
	//		"msg":  "success",
	//	})
	//})
	//return router

	router := gin.Default()
	v1 := router.Group("/v1")
	AddUserRoutesV1(v1)
	AddSearchRoutesV1(v1)
	AddSysUserV1(v1)

	v2 := router.Group("/v2")
	AddSearchRoutesV2(v2)

	return router
}
