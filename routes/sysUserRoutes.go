package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lizhixin1992/gin-test/controllers/api"
	"net/http"
)

func AddSysUserV1(rg *gin.RouterGroup) {
	group := rg.Group("/sys/user")

	group.GET("/:id", func(context *gin.Context) {
		id := context.Param("id")
		result := api.GetSysUserById(id)
		context.JSON(http.StatusOK, gin.H{
			"code": "3",
			"msg":  "success",
			"data": result,
		})
	})
}
