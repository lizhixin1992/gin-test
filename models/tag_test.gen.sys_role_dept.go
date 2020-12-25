package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _SysRoleDeptMgr struct {
	*_BaseMgr
}

// SysRoleDeptMgr open func
func SysRoleDeptMgr(db *gorm.DB) *_SysRoleDeptMgr {
	if db == nil {
		panic(fmt.Errorf("SysRoleDeptMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SysRoleDeptMgr{_BaseMgr: &_BaseMgr{DB: db.Table("sys_role_dept"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SysRoleDeptMgr) GetTableName() string {
	return "sys_role_dept"
}

// Get 获取
func (obj *_SysRoleDeptMgr) Get() (result SysRoleDept, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SysRoleDeptMgr) Gets() (results []*SysRoleDept, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithRoleID role_id获取 角色ID
func (obj *_SysRoleDeptMgr) WithRoleID(roleID int64) Option {
	return optionFunc(func(o *options) { o.query["role_id"] = roleID })
}

// WithDeptID dept_id获取 部门ID
func (obj *_SysRoleDeptMgr) WithDeptID(deptID int64) Option {
	return optionFunc(func(o *options) { o.query["dept_id"] = deptID })
}

// GetByOption 功能选项模式获取
func (obj *_SysRoleDeptMgr) GetByOption(opts ...Option) (result SysRoleDept, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_SysRoleDeptMgr) GetByOptions(opts ...Option) (results []*SysRoleDept, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromRoleID 通过role_id获取内容 角色ID
func (obj *_SysRoleDeptMgr) GetFromRoleID(roleID int64) (results []*SysRoleDept, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("role_id = ?", roleID).Find(&results).Error

	return
}

// GetBatchFromRoleID 批量唯一主键查找 角色ID
func (obj *_SysRoleDeptMgr) GetBatchFromRoleID(roleIDs []int64) (results []*SysRoleDept, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("role_id IN (?)", roleIDs).Find(&results).Error

	return
}

// GetFromDeptID 通过dept_id获取内容 部门ID
func (obj *_SysRoleDeptMgr) GetFromDeptID(deptID int64) (results []*SysRoleDept, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("dept_id = ?", deptID).Find(&results).Error

	return
}

// GetBatchFromDeptID 批量唯一主键查找 部门ID
func (obj *_SysRoleDeptMgr) GetBatchFromDeptID(deptIDs []int64) (results []*SysRoleDept, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("dept_id IN (?)", deptIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_SysRoleDeptMgr) FetchByPrimaryKey(roleID int64, deptID int64) (result SysRoleDept, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("role_id = ? AND dept_id = ?", roleID, deptID).Find(&result).Error

	return
}
