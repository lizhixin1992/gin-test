package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _SysDictTypeMgr struct {
	*_BaseMgr
}

// SysDictTypeMgr open func
func SysDictTypeMgr(db *gorm.DB) *_SysDictTypeMgr {
	if db == nil {
		panic(fmt.Errorf("SysDictTypeMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SysDictTypeMgr{_BaseMgr: &_BaseMgr{DB: db.Table("sys_dict_type"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SysDictTypeMgr) GetTableName() string {
	return "sys_dict_type"
}

// Get 获取
func (obj *_SysDictTypeMgr) Get() (result SysDictType, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SysDictTypeMgr) Gets() (results []*SysDictType, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithDictID dict_id获取 字典主键
func (obj *_SysDictTypeMgr) WithDictID(dictID int64) Option {
	return optionFunc(func(o *options) { o.query["dict_id"] = dictID })
}

// WithDictName dict_name获取 字典名称
func (obj *_SysDictTypeMgr) WithDictName(dictName string) Option {
	return optionFunc(func(o *options) { o.query["dict_name"] = dictName })
}

// WithDictType dict_type获取 字典类型
func (obj *_SysDictTypeMgr) WithDictType(dictType string) Option {
	return optionFunc(func(o *options) { o.query["dict_type"] = dictType })
}

// WithStatus status获取 状态（0正常 1停用）
func (obj *_SysDictTypeMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithCreateBy create_by获取 创建者
func (obj *_SysDictTypeMgr) WithCreateBy(createBy string) Option {
	return optionFunc(func(o *options) { o.query["create_by"] = createBy })
}

// WithCreateTime create_time获取 创建时间
func (obj *_SysDictTypeMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateBy update_by获取 更新者
func (obj *_SysDictTypeMgr) WithUpdateBy(updateBy string) Option {
	return optionFunc(func(o *options) { o.query["update_by"] = updateBy })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_SysDictTypeMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithRemark remark获取 备注
func (obj *_SysDictTypeMgr) WithRemark(remark string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// GetByOption 功能选项模式获取
func (obj *_SysDictTypeMgr) GetByOption(opts ...Option) (result SysDictType, err error) {
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
func (obj *_SysDictTypeMgr) GetByOptions(opts ...Option) (results []*SysDictType, err error) {
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

// GetFromDictID 通过dict_id获取内容 字典主键
func (obj *_SysDictTypeMgr) GetFromDictID(dictID int64) (result SysDictType, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("dict_id = ?", dictID).Find(&result).Error

	return
}

// GetBatchFromDictID 批量唯一主键查找 字典主键
func (obj *_SysDictTypeMgr) GetBatchFromDictID(dictIDs []int64) (results []*SysDictType, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("dict_id IN (?)", dictIDs).Find(&results).Error

	return
}

// GetFromDictName 通过dict_name获取内容 字典名称
func (obj *_SysDictTypeMgr) GetFromDictName(dictName string) (results []*SysDictType, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("dict_name = ?", dictName).Find(&results).Error

	return
}

// GetBatchFromDictName 批量唯一主键查找 字典名称
func (obj *_SysDictTypeMgr) GetBatchFromDictName(dictNames []string) (results []*SysDictType, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("dict_name IN (?)", dictNames).Find(&results).Error

	return
}

// GetFromDictType 通过dict_type获取内容 字典类型
func (obj *_SysDictTypeMgr) GetFromDictType(dictType string) (result SysDictType, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("dict_type = ?", dictType).Find(&result).Error

	return
}

// GetBatchFromDictType 批量唯一主键查找 字典类型
func (obj *_SysDictTypeMgr) GetBatchFromDictType(dictTypes []string) (results []*SysDictType, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("dict_type IN (?)", dictTypes).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 状态（0正常 1停用）
func (obj *_SysDictTypeMgr) GetFromStatus(status string) (results []*SysDictType, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量唯一主键查找 状态（0正常 1停用）
func (obj *_SysDictTypeMgr) GetBatchFromStatus(statuss []string) (results []*SysDictType, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status IN (?)", statuss).Find(&results).Error

	return
}

// GetFromCreateBy 通过create_by获取内容 创建者
func (obj *_SysDictTypeMgr) GetFromCreateBy(createBy string) (results []*SysDictType, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_by = ?", createBy).Find(&results).Error

	return
}

// GetBatchFromCreateBy 批量唯一主键查找 创建者
func (obj *_SysDictTypeMgr) GetBatchFromCreateBy(createBys []string) (results []*SysDictType, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_by IN (?)", createBys).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_SysDictTypeMgr) GetFromCreateTime(createTime time.Time) (results []*SysDictType, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_time = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量唯一主键查找 创建时间
func (obj *_SysDictTypeMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*SysDictType, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_time IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromUpdateBy 通过update_by获取内容 更新者
func (obj *_SysDictTypeMgr) GetFromUpdateBy(updateBy string) (results []*SysDictType, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_by = ?", updateBy).Find(&results).Error

	return
}

// GetBatchFromUpdateBy 批量唯一主键查找 更新者
func (obj *_SysDictTypeMgr) GetBatchFromUpdateBy(updateBys []string) (results []*SysDictType, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_by IN (?)", updateBys).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_SysDictTypeMgr) GetFromUpdateTime(updateTime time.Time) (results []*SysDictType, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_time = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量唯一主键查找 更新时间
func (obj *_SysDictTypeMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*SysDictType, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_time IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_SysDictTypeMgr) GetFromRemark(remark string) (results []*SysDictType, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("remark = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量唯一主键查找 备注
func (obj *_SysDictTypeMgr) GetBatchFromRemark(remarks []string) (results []*SysDictType, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("remark IN (?)", remarks).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_SysDictTypeMgr) FetchByPrimaryKey(dictID int64) (result SysDictType, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("dict_id = ?", dictID).Find(&result).Error

	return
}

// FetchUniqueByDictType primay or index 获取唯一内容
func (obj *_SysDictTypeMgr) FetchUniqueByDictType(dictType string) (result SysDictType, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("dict_type = ?", dictType).Find(&result).Error

	return
}
