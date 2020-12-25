package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lizhixin1992/gin-test/controllers/api"
	"net/http"
)

func AddSysUserV1(rg *gin.RouterGroup) {
	apiController := api.NewSysUserController()
	group := rg.Group("/sys/user")
	group.GET("/:id", func(context *gin.Context) {
		id := context.Param("id")
		result, err := apiController.GetSysUserById(id)
		if err != nil {
			context.JSON(http.StatusOK, gin.H{
				"code": "1",
				"msg":  "success",
				"data": "",
			})
		}else {
			context.JSON(http.StatusOK, gin.H{
				"code": "0",
				"msg":  "success",
				"data": result,
			})
		}
	})
}
