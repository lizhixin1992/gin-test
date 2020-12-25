package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _LabelUploadRelMgr struct {
	*_BaseMgr
}

// LabelUploadRelMgr open func
func LabelUploadRelMgr(db *gorm.DB) *_LabelUploadRelMgr {
	if db == nil {
		panic(fmt.Errorf("LabelUploadRelMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_LabelUploadRelMgr{_BaseMgr: &_BaseMgr{DB: db.Table("label_upload_rel"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_LabelUploadRelMgr) GetTableName() string {
	return "label_upload_rel"
}

// Get 获取
func (obj *_LabelUploadRelMgr) Get() (result LabelUploadRel, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_LabelUploadRelMgr) Gets() (results []*LabelUploadRel, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithUploadUUID upload_uuid获取 该条上传的唯一标识
func (obj *_LabelUploadRelMgr) WithUploadUUID(uploadUUID string) Option {
	return optionFunc(func(o *options) { o.query["upload_uuid"] = uploadUUID })
}

// WithUploadType upload_type获取 上传类型(labelTag,labelCharacter)
func (obj *_LabelUploadRelMgr) WithUploadType(uploadType string) Option {
	return optionFunc(func(o *options) { o.query["upload_type"] = uploadType })
}

// WithInfoID info_id获取 消息ID，用于异步回调时通知调用者的唯一性标识
func (obj *_LabelUploadRelMgr) WithInfoID(infoID string) Option {
	return optionFunc(func(o *options) { o.query["info_id"] = infoID })
}

// WithLabelUUID label_uuid获取 关联数据的唯一标识
func (obj *_LabelUploadRelMgr) WithLabelUUID(labelUUID string) Option {
	return optionFunc(func(o *options) { o.query["label_uuid"] = labelUUID })
}

// WithUploadPath upload_path获取 上传地址
func (obj *_LabelUploadRelMgr) WithUploadPath(uploadPath string) Option {
	return optionFunc(func(o *options) { o.query["upload_path"] = uploadPath })
}

// WithCreateTime create_time获取 创建时间
func (obj *_LabelUploadRelMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateTime update_time获取 修改时间
func (obj *_LabelUploadRelMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithStatus status获取 状态:0:未成功;1:成功
func (obj *_LabelUploadRelMgr) WithStatus(status int) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithUploadUser upload_user获取 上传人
func (obj *_LabelUploadRelMgr) WithUploadUser(uploadUser string) Option {
	return optionFunc(func(o *options) { o.query["upload_user"] = uploadUser })
}

// GetByOption 功能选项模式获取
func (obj *_LabelUploadRelMgr) GetByOption(opts ...Option) (result LabelUploadRel, err error) {
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
func (obj *_LabelUploadRelMgr) GetByOptions(opts ...Option) (results []*LabelUploadRel, err error) {
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

// GetFromUploadUUID 通过upload_uuid获取内容 该条上传的唯一标识
func (obj *_LabelUploadRelMgr) GetFromUploadUUID(uploadUUID string) (result LabelUploadRel, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("upload_uuid = ?", uploadUUID).Find(&result).Error

	return
}

// GetBatchFromUploadUUID 批量唯一主键查找 该条上传的唯一标识
func (obj *_LabelUploadRelMgr) GetBatchFromUploadUUID(uploadUUIDs []string) (results []*LabelUploadRel, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("upload_uuid IN (?)", uploadUUIDs).Find(&results).Error

	return
}

// GetFromUploadType 通过upload_type获取内容 上传类型(labelTag,labelCharacter)
func (obj *_LabelUploadRelMgr) GetFromUploadType(uploadType string) (results []*LabelUploadRel, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("upload_type = ?", uploadType).Find(&results).Error

	return
}

// GetBatchFromUploadType 批量唯一主键查找 上传类型(labelTag,labelCharacter)
func (obj *_LabelUploadRelMgr) GetBatchFromUploadType(uploadTypes []string) (results []*LabelUploadRel, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("upload_type IN (?)", uploadTypes).Find(&results).Error

	return
}

// GetFromInfoID 通过info_id获取内容 消息ID，用于异步回调时通知调用者的唯一性标识
func (obj *_LabelUploadRelMgr) GetFromInfoID(infoID string) (results []*LabelUploadRel, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("info_id = ?", infoID).Find(&results).Error

	return
}

// GetBatchFromInfoID 批量唯一主键查找 消息ID，用于异步回调时通知调用者的唯一性标识
func (obj *_LabelUploadRelMgr) GetBatchFromInfoID(infoIDs []string) (results []*LabelUploadRel, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("info_id IN (?)", infoIDs).Find(&results).Error

	return
}

// GetFromLabelUUID 通过label_uuid获取内容 关联数据的唯一标识
func (obj *_LabelUploadRelMgr) GetFromLabelUUID(labelUUID string) (results []*LabelUploadRel, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("label_uuid = ?", labelUUID).Find(&results).Error

	return
}

// GetBatchFromLabelUUID 批量唯一主键查找 关联数据的唯一标识
func (obj *_LabelUploadRelMgr) GetBatchFromLabelUUID(labelUUIDs []string) (results []*LabelUploadRel, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("label_uuid IN (?)", labelUUIDs).Find(&results).Error

	return
}

// GetFromUploadPath 通过upload_path获取内容 上传地址
func (obj *_LabelUploadRelMgr) GetFromUploadPath(uploadPath string) (results []*LabelUploadRel, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("upload_path = ?", uploadPath).Find(&results).Error

	return
}

// GetBatchFromUploadPath 批量唯一主键查找 上传地址
func (obj *_LabelUploadRelMgr) GetBatchFromUploadPath(uploadPaths []string) (results []*LabelUploadRel, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("upload_path IN (?)", uploadPaths).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_LabelUploadRelMgr) GetFromCreateTime(createTime time.Time) (results []*LabelUploadRel, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_time = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量唯一主键查找 创建时间
func (obj *_LabelUploadRelMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*LabelUploadRel, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_time IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 修改时间
func (obj *_LabelUploadRelMgr) GetFromUpdateTime(updateTime time.Time) (results []*LabelUploadRel, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_time = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量唯一主键查找 修改时间
func (obj *_LabelUploadRelMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*LabelUploadRel, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_time IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 状态:0:未成功;1:成功
func (obj *_LabelUploadRelMgr) GetFromStatus(status int) (results []*LabelUploadRel, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量唯一主键查找 状态:0:未成功;1:成功
func (obj *_LabelUploadRelMgr) GetBatchFromStatus(statuss []int) (results []*LabelUploadRel, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status IN (?)", statuss).Find(&results).Error

	return
}

// GetFromUploadUser 通过upload_user获取内容 上传人
func (obj *_LabelUploadRelMgr) GetFromUploadUser(uploadUser string) (results []*LabelUploadRel, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("upload_user = ?", uploadUser).Find(&results).Error

	return
}

// GetBatchFromUploadUser 批量唯一主键查找 上传人
func (obj *_LabelUploadRelMgr) GetBatchFromUploadUser(uploadUsers []string) (results []*LabelUploadRel, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("upload_user IN (?)", uploadUsers).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_LabelUploadRelMgr) FetchByPrimaryKey(uploadUUID string) (result LabelUploadRel, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("upload_uuid = ?", uploadUUID).Find(&result).Error

	return
}
