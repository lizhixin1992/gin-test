package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _SysUserDictDataMgr struct {
	*_BaseMgr
}

// SysUserDictDataMgr open func
func SysUserDictDataMgr(db *gorm.DB) *_SysUserDictDataMgr {
	if db == nil {
		panic(fmt.Errorf("SysUserDictDataMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SysUserDictDataMgr{_BaseMgr: &_BaseMgr{DB: db.Table("sys_user_dict_data"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SysUserDictDataMgr) GetTableName() string {
	return "sys_user_dict_data"
}

// Get 获取
func (obj *_SysUserDictDataMgr) Get() (result SysUserDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SysUserDictDataMgr) Gets() (results []*SysUserDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithUserID user_id获取 用户ID
func (obj *_SysUserDictDataMgr) WithUserID(userID int64) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithDictCode dict_code获取 字典数据表ID
func (obj *_SysUserDictDataMgr) WithDictCode(dictCode int64) Option {
	return optionFunc(func(o *options) { o.query["dict_code"] = dictCode })
}

// GetByOption 功能选项模式获取
func (obj *_SysUserDictDataMgr) GetByOption(opts ...Option) (result SysUserDictData, err error) {
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
func (obj *_SysUserDictDataMgr) GetByOptions(opts ...Option) (results []*SysUserDictData, err error) {
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

// GetFromUserID 通过user_id获取内容 用户ID
func (obj *_SysUserDictDataMgr) GetFromUserID(userID int64) (results []*SysUserDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id = ?", userID).Find(&results).Error

	return
}

// GetBatchFromUserID 批量唯一主键查找 用户ID
func (obj *_SysUserDictDataMgr) GetBatchFromUserID(userIDs []int64) (results []*SysUserDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromDictCode 通过dict_code获取内容 字典数据表ID
func (obj *_SysUserDictDataMgr) GetFromDictCode(dictCode int64) (results []*SysUserDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("dict_code = ?", dictCode).Find(&results).Error

	return
}

// GetBatchFromDictCode 批量唯一主键查找 字典数据表ID
func (obj *_SysUserDictDataMgr) GetBatchFromDictCode(dictCodes []int64) (results []*SysUserDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("dict_code IN (?)", dictCodes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_SysUserDictDataMgr) FetchByPrimaryKey(userID int64, dictCode int64) (result SysUserDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id = ? AND dict_code = ?", userID, dictCode).Find(&result).Error

	return
}
