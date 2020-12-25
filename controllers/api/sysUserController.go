package api

import (
	"github.com/lizhixin1992/gin-test/database"
	"github.com/lizhixin1992/gin-test/models"
	"strconv"
)

func GetSysUserById(id string) (result models.SysUser) {
	sysUserMgr := models.SysUserMgr(database.InstanceMaster())
	int64, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return models.SysUser{}
	}
	sysUser, _ := sysUserMgr.GetFromUserID(int64)
	return sysUser
}
