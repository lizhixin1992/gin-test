package api

import (
	"errors"
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

func (c *SysUserController) GetSysUserById(id string) (result models.SysUser, err error) {
	int64, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return models.SysUser{}, errors.New("id is error")
	}
	result, err1 := c.Service.GetFromUserID(int64)
	if err1 != nil || result.UserID == 0 {
		return models.SysUser{}, errors.New("not have data")
	}
	return result, nil
}

func (c *SysUserController) SaveSysUser(sysUser models.SysUser) (err error) {
	return c.Service.Save(sysUser).Error
}

func (c *SysUserController) GetListByIds(userIDs []int64) (results []*models.SysUser, err error) {
	return c.Service.GetBatchFromUserID(userIDs)
}
