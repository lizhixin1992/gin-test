package routes

import (
	"github.com/gin-gonic/gin"
	commons "github.com/lizhixin1992/gin-test/commons/vo"
	"github.com/lizhixin1992/gin-test/controllers/api"
	"github.com/lizhixin1992/gin-test/models"
	"github.com/lizhixin1992/gin-test/models/condition"
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
			//todo:类似新增，修改操作对象，要向下传该对象的地址
			err := apiController.SaveSysUser(&sysUser)
			if err != nil {
				commons.SetResponseOperateFail(context)
			} else {
				commons.SetResponseSuccess(context)
			}
		} else {
			commons.SetResponseDataFail(context)
		}
	})

	group.GET("/search", func(context *gin.Context) {
		var conditions condition.SysUserCondition
		if context.ShouldBind(&conditions) == nil {
			if conditions.Ids != "" {
				split := strings.Split(conditions.Ids, ",")
				ids := make([]int64, 8)
				for i := 0; i < len(split); i++ {
					ids[i], _ = strconv.ParseInt(split[i], 10, 64)
				}

				results, err := apiController.GetListByIds(ids[:len(split)])
				if err == nil {
					commons.SetResponseSuccessData(context, results)
				} else {
					commons.SetResponseOperateFail(context)
				}
			}
		} else {
			commons.SetResponseSuccess(context)
		}
	})

}
