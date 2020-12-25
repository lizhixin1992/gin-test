package api

import (
	"github.com/lizhixin1992/gin-test/database"
	"github.com/lizhixin1992/gin-test/models"
	"strconv"
)

type SysUserController struct {
	Service *models.SysUserMgr
}

func NewSysUserController() *SysUserController {
	return &SysUserController{
		Service: models.NewSysUserMgr(database.InstanceMaster()),
	}
}

func (c *SysUserController) GetSysUserById(id string) (result models.SysUser) {
	int64, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return models.SysUser{}
	}
	sysUser, _ := c.Service.GetFromUserID(int64)
	return sysUser
}
