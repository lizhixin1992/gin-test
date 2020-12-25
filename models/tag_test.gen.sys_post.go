package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _SysPostMgr struct {
	*_BaseMgr
}

// SysPostMgr open func
func SysPostMgr(db *gorm.DB) *_SysPostMgr {
	if db == nil {
		panic(fmt.Errorf("SysPostMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SysPostMgr{_BaseMgr: &_BaseMgr{DB: db.Table("sys_post"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SysPostMgr) GetTableName() string {
	return "sys_post"
}

// Get 获取
func (obj *_SysPostMgr) Get() (result SysPost, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SysPostMgr) Gets() (results []*SysPost, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithPostID post_id获取 岗位ID
func (obj *_SysPostMgr) WithPostID(postID int64) Option {
	return optionFunc(func(o *options) { o.query["post_id"] = postID })
}

// WithPostCode post_code获取 岗位编码
func (obj *_SysPostMgr) WithPostCode(postCode string) Option {
	return optionFunc(func(o *options) { o.query["post_code"] = postCode })
}

// WithPostName post_name获取 岗位名称
func (obj *_SysPostMgr) WithPostName(postName string) Option {
	return optionFunc(func(o *options) { o.query["post_name"] = postName })
}

// WithPostSort post_sort获取 显示顺序
func (obj *_SysPostMgr) WithPostSort(postSort int) Option {
	return optionFunc(func(o *options) { o.query["post_sort"] = postSort })
}

// WithStatus status获取 状态（0正常 1停用）
func (obj *_SysPostMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithCreateBy create_by获取 创建者
func (obj *_SysPostMgr) WithCreateBy(createBy string) Option {
	return optionFunc(func(o *options) { o.query["create_by"] = createBy })
}

// WithCreateTime create_time获取 创建时间
func (obj *_SysPostMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateBy update_by获取 更新者
func (obj *_SysPostMgr) WithUpdateBy(updateBy string) Option {
	return optionFunc(func(o *options) { o.query["update_by"] = updateBy })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_SysPostMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithRemark remark获取 备注
func (obj *_SysPostMgr) WithRemark(remark string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// GetByOption 功能选项模式获取
func (obj *_SysPostMgr) GetByOption(opts ...Option) (result SysPost, err error) {
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
func (obj *_SysPostMgr) GetByOptions(opts ...Option) (results []*SysPost, err error) {
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

// GetFromPostID 通过post_id获取内容 岗位ID
func (obj *_SysPostMgr) GetFromPostID(postID int64) (result SysPost, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("post_id = ?", postID).Find(&result).Error

	return
}

// GetBatchFromPostID 批量唯一主键查找 岗位ID
func (obj *_SysPostMgr) GetBatchFromPostID(postIDs []int64) (results []*SysPost, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("post_id IN (?)", postIDs).Find(&results).Error

	return
}

// GetFromPostCode 通过post_code获取内容 岗位编码
func (obj *_SysPostMgr) GetFromPostCode(postCode string) (results []*SysPost, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("post_code = ?", postCode).Find(&results).Error

	return
}

// GetBatchFromPostCode 批量唯一主键查找 岗位编码
func (obj *_SysPostMgr) GetBatchFromPostCode(postCodes []string) (results []*SysPost, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("post_code IN (?)", postCodes).Find(&results).Error

	return
}

// GetFromPostName 通过post_name获取内容 岗位名称
func (obj *_SysPostMgr) GetFromPostName(postName string) (results []*SysPost, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("post_name = ?", postName).Find(&results).Error

	return
}

// GetBatchFromPostName 批量唯一主键查找 岗位名称
func (obj *_SysPostMgr) GetBatchFromPostName(postNames []string) (results []*SysPost, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("post_name IN (?)", postNames).Find(&results).Error

	return
}

// GetFromPostSort 通过post_sort获取内容 显示顺序
func (obj *_SysPostMgr) GetFromPostSort(postSort int) (results []*SysPost, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("post_sort = ?", postSort).Find(&results).Error

	return
}

// GetBatchFromPostSort 批量唯一主键查找 显示顺序
func (obj *_SysPostMgr) GetBatchFromPostSort(postSorts []int) (results []*SysPost, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("post_sort IN (?)", postSorts).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 状态（0正常 1停用）
func (obj *_SysPostMgr) GetFromStatus(status string) (results []*SysPost, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量唯一主键查找 状态（0正常 1停用）
func (obj *_SysPostMgr) GetBatchFromStatus(statuss []string) (results []*SysPost, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status IN (?)", statuss).Find(&results).Error

	return
}

// GetFromCreateBy 通过create_by获取内容 创建者
func (obj *_SysPostMgr) GetFromCreateBy(createBy string) (results []*SysPost, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_by = ?", createBy).Find(&results).Error

	return
}

// GetBatchFromCreateBy 批量唯一主键查找 创建者
func (obj *_SysPostMgr) GetBatchFromCreateBy(createBys []string) (results []*SysPost, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_by IN (?)", createBys).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_SysPostMgr) GetFromCreateTime(createTime time.Time) (results []*SysPost, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_time = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量唯一主键查找 创建时间
func (obj *_SysPostMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*SysPost, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_time IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromUpdateBy 通过update_by获取内容 更新者
func (obj *_SysPostMgr) GetFromUpdateBy(updateBy string) (results []*SysPost, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_by = ?", updateBy).Find(&results).Error

	return
}

// GetBatchFromUpdateBy 批量唯一主键查找 更新者
func (obj *_SysPostMgr) GetBatchFromUpdateBy(updateBys []string) (results []*SysPost, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_by IN (?)", updateBys).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_SysPostMgr) GetFromUpdateTime(updateTime time.Time) (results []*SysPost, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_time = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量唯一主键查找 更新时间
func (obj *_SysPostMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*SysPost, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_time IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_SysPostMgr) GetFromRemark(remark string) (results []*SysPost, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("remark = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量唯一主键查找 备注
func (obj *_SysPostMgr) GetBatchFromRemark(remarks []string) (results []*SysPost, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("remark IN (?)", remarks).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_SysPostMgr) FetchByPrimaryKey(postID int64) (result SysPost, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("post_id = ?", postID).Find(&result).Error

	return
}
