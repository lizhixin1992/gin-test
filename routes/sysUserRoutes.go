package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lizhixin1992/gin-test/controllers/api"
	"github.com/lizhixin1992/gin-test/models"
	"github.com/lizhixin1992/gin-test/models/condition"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func AddSysUserV1(rg *gin.RouterGroup) {
	apiController := api.NewSysUserController()
	group := rg.Group("/sys/user")

	//group.GET("/:id", func(context *gin.Context) {
	//	id := context.Param("id")
	//	result, err := apiController.GetSysUserById(id)
	//	if err != nil {
	//		context.JSON(http.StatusOK, gin.H{
	//			"code": "1",
	//			"msg":  "success",
	//			"data": "",
	//		})
	//	} else {
	//		context.JSON(http.StatusOK, gin.H{
	//			"code": "0",
	//			"msg":  "success",
	//			"data": result,
	//		})
	//	}
	//})

	group.POST("/save", func(context *gin.Context) {
		var sysUser models.SysUser
		if context.ShouldBindJSON(&sysUser) == nil {
			sysUser.LoginDate = time.Now()
			sysUser.CreateTime = time.Now()
			sysUser.UpdateTime = time.Now()
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
		}else {
			context.JSON(http.StatusOK, gin.H{
				"code": "1002",
				"msg":  "error",
			})
		}
	})


	group.GET("/search", func(context *gin.Context) {
		var conditions condition.SysUserCondition
		if context.ShouldBind(&conditions) == nil{
			if conditions.Ids != ""{
				split := strings.Split(conditions.Ids, ",")
				ids := make([]int64, 8)
				for i := 0; i < len(split); i++ {
					ids[i],_ = strconv.ParseInt(split[i],10,64)
				}
				
				results, err := apiController.GetListByIds(ids[:len(split)])
				if err == nil {
					context.JSON(http.StatusOK, gin.H{
						"code": "0",
						"msg":  "success",
						"data": results,
					})
				}else {
					context.JSON(http.StatusOK, gin.H{
						"code": "1001",
						"msg":  "error",
					})
				}
			}
		}else {
			context.JSON(http.StatusOK, gin.H{
				"code": "1002",
				"msg":  "error",
			})
		}
	})

}
