package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _SysUserPostMgr struct {
	*_BaseMgr
}

// SysUserPostMgr open func
func SysUserPostMgr(db *gorm.DB) *_SysUserPostMgr {
	if db == nil {
		panic(fmt.Errorf("SysUserPostMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SysUserPostMgr{_BaseMgr: &_BaseMgr{DB: db.Table("sys_user_post"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SysUserPostMgr) GetTableName() string {
	return "sys_user_post"
}

// Get 获取
func (obj *_SysUserPostMgr) Get() (result SysUserPost, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SysUserPostMgr) Gets() (results []*SysUserPost, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithUserID user_id获取 用户ID
func (obj *_SysUserPostMgr) WithUserID(userID int64) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithPostID post_id获取 岗位ID
func (obj *_SysUserPostMgr) WithPostID(postID int64) Option {
	return optionFunc(func(o *options) { o.query["post_id"] = postID })
}

// GetByOption 功能选项模式获取
func (obj *_SysUserPostMgr) GetByOption(opts ...Option) (result SysUserPost, err error) {
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
func (obj *_SysUserPostMgr) GetByOptions(opts ...Option) (results []*SysUserPost, err error) {
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
func (obj *_SysUserPostMgr) GetFromUserID(userID int64) (results []*SysUserPost, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id = ?", userID).Find(&results).Error

	return
}

// GetBatchFromUserID 批量唯一主键查找 用户ID
func (obj *_SysUserPostMgr) GetBatchFromUserID(userIDs []int64) (results []*SysUserPost, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromPostID 通过post_id获取内容 岗位ID
func (obj *_SysUserPostMgr) GetFromPostID(postID int64) (results []*SysUserPost, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("post_id = ?", postID).Find(&results).Error

	return
}

// GetBatchFromPostID 批量唯一主键查找 岗位ID
func (obj *_SysUserPostMgr) GetBatchFromPostID(postIDs []int64) (results []*SysUserPost, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("post_id IN (?)", postIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_SysUserPostMgr) FetchByPrimaryKey(userID int64, postID int64) (result SysUserPost, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id = ? AND post_id = ?", userID, postID).Find(&result).Error

	return
}
