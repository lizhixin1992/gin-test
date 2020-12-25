package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lizhixin1992/gin-test/controllers/api"
	"github.com/lizhixin1992/gin-test/models"
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
		} else {
			context.JSON(http.StatusOK, gin.H{
				"code": "0",
				"msg":  "success",
				"data": result,
			})
		}
	})

	group.POST("/save", func(context *gin.Context) {
		var sysUser models.SysUser
		if context.ShouldBindJSON(&sysUser) == nil {
			context.JSON(http.StatusOK, gin.H{
				"code": "1001",
				"msg":  "error",
			})
		} else {
			err := apiController.SaveSysUser(sysUser)
			if err != nil {
				context.JSON(http.StatusOK, gin.H{
					"code": "1001",
					"msg":  "error",
				})
			} else {
				context.JSON(http.StatusOK, gin.H{
					"code": "0",
					"msg":  "success",
				})
			}
		}
	})

}
