package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lizhixin1992/gin-test/models"
	"log"
	"net/http"
)

func AddUserRoutesV1(rg *gin.RouterGroup) {
	users := rg.Group("/user")

	//普通的路由请求
	users.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"code": "0",
			"msg":  "success",
		})
	})

	users.GET("/list", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"code": "2",
			"msg":  "success",
		})
	})

	users.GET("/condition/:id", func(context *gin.Context) {
		id := context.Param("id")
		context.JSON(http.StatusOK, gin.H{
			"code": "3",
			"msg":  "success",
			"data": id,
		})
	})

	// 此 handler 将匹配 /user/1221/ 和 /user/1221/send
	// 如果没有其他路由匹配 /user/1221，它将重定向到 /user/1221/
	users.GET("/condition/:id/*action", func(context *gin.Context) {
		name := context.Param("name")
		action := context.Param("action")
		data := name + "is" + action
		context.JSON(http.StatusOK, gin.H{
			"code": "4",
			"msg":  "success",
			"data": data,
		})
	})

	users.POST("/search", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"code": "10",
			"msg":  "success",
		})
	})

	//xxxxxx?a=xxxx&b=xxxxx
	users.GET("/save", func(context *gin.Context) {
		var user models.User
		if context.ShouldBind(&user) == nil {
			log.Println("====== &= ======")
			log.Println(user.Name)
			log.Println(user.Address)
		}
		context.String(http.StatusOK, "Success")
	})

	//Content-Type:application/json
	users.POST("/save/json", func(context *gin.Context) {
		var user models.User
		if context.ShouldBindJSON(&user) == nil {
			log.Println("====== JSON ======")
			log.Println(user.Name)
			log.Println(user.Address)
		}
		context.String(http.StatusOK, "Success")
	})

}