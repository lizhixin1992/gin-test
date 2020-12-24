package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddSearchRoutesV1(rg *gin.RouterGroup)  {
	search := rg.Group("/search")

	//普通的路由请求
	search.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"code": "100",
			"msg":  "success",
		})
	})

	search.GET("/condition/:name", func(context *gin.Context) {
		name := context.Param("name")
		context.JSON(http.StatusOK, gin.H{
			"code": "101",
			"msg":  "success",
			"data": name,
		})
	})
}


func AddSearchRoutesV2(rg *gin.RouterGroup)  {
	search := rg.Group("/search")

	//普通的路由请求
	search.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"code": "100",
			"msg":  "success",
		})
	})

	search.GET("/condition/:name/:tag", func(context *gin.Context) {
		name := context.Param("name")
		tag := context.Param("type")
		data := name + ", " + tag
		context.JSON(http.StatusOK, gin.H{
			"code": "101",
			"msg":  "success",
			"data": data,
		})
	})
}