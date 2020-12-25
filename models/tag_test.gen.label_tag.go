package	models	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
"time"	
)	

type _LabelTagMgr struct {
	*_BaseMgr
}

// LabelTagMgr open func
func LabelTagMgr(db *gorm.DB) *_LabelTagMgr {
	if db == nil {
		panic(fmt.Errorf("LabelTagMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_LabelTagMgr{_BaseMgr: &_BaseMgr{DB: db.Table("label_tag"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_LabelTagMgr) GetTableName() string {
	return "label_tag"
}

// Get 获取
func (obj *_LabelTagMgr) Get() (result LabelTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_LabelTagMgr) Gets() (results []*LabelTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithTagID tag_id获取 标签主键
func (obj *_LabelTagMgr) WithTagID(tagID int) Option {
	return optionFunc(func(o *options) { o.query["tag_id"] = tagID })
}

// WithTagUUID tag_uuid获取 唯一标识
func (obj *_LabelTagMgr) WithTagUUID(tagUUID string) Option {
	return optionFunc(func(o *options) { o.query["tag_uuid"] = tagUUID })
}

// WithTagName tag_name获取 标签名称
func (obj *_LabelTagMgr) WithTagName(tagName string) Option {
	return optionFunc(func(o *options) { o.query["tag_name"] = tagName })
}

// WithTagType tag_type获取 标签类型
func (obj *_LabelTagMgr) WithTagType(tagType int) Option {
	return optionFunc(func(o *options) { o.query["tag_type"] = tagType })
}

// WithTagParentID tag_parent_id获取 父级标识
func (obj *_LabelTagMgr) WithTagParentID(tagParentID string) Option {
	return optionFunc(func(o *options) { o.query["tag_parent_id"] = tagParentID })
}

// WithTagLayer tag_layer获取 层级
func (obj *_LabelTagMgr) WithTagLayer(tagLayer int) Option {
	return optionFunc(func(o *options) { o.query["tag_layer"] = tagLayer })
}

// WithTagImg tag_img获取 图片
func (obj *_LabelTagMgr) WithTagImg(tagImg string) Option {
	return optionFunc(func(o *options) { o.query["tag_img"] = tagImg })
}

// WithTagImgAliOssPath tag_img_ali_oss_path获取 图片在阿里云OSS上地址
func (obj *_LabelTagMgr) WithTagImgAliOssPath(tagImgAliOssPath string) Option {
	return optionFunc(func(o *options) { o.query["tag_img_ali_oss_path"] = tagImgAliOssPath })
}

// WithApplyStatus apply_status获取 是否启用,0:启用,1:不启用
func (obj *_LabelTagMgr) WithApplyStatus(applyStatus int) Option {
	return optionFunc(func(o *options) { o.query["apply_status"] = applyStatus })
}

// WithDeleteStatus delete_status获取 是否删除,0:未删除,1:已删除
func (obj *_LabelTagMgr) WithDeleteStatus(deleteStatus int) Option {
	return optionFunc(func(o *options) { o.query["delete_status"] = deleteStatus })
}

// WithCreateTime create_time获取 创建时间
func (obj *_LabelTagMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateTime update_time获取 修改时间
func (obj *_LabelTagMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithCreateUser create_user获取 创建人
func (obj *_LabelTagMgr) WithCreateUser(createUser string) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateUser update_user获取 修改人
func (obj *_LabelTagMgr) WithUpdateUser(updateUser string) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithOrderNum order_num获取 排序序号
func (obj *_LabelTagMgr) WithOrderNum(orderNum int) Option {
	return optionFunc(func(o *options) { o.query["order_num"] = orderNum })
}

// WithUploadUUID upload_uuid获取 
func (obj *_LabelTagMgr) WithUploadUUID(uploadUUID string) Option {
	return optionFunc(func(o *options) { o.query["upload_uuid"] = uploadUUID })
}

// WithIsHaveChildren is_have_children获取 是否可以有下级,0:可以,1:不可以
func (obj *_LabelTagMgr) WithIsHaveChildren(isHaveChildren int) Option {
	return optionFunc(func(o *options) { o.query["is_have_children"] = isHaveChildren })
}


// GetByOption 功能选项模式获取
func (obj *_LabelTagMgr) GetByOption(opts ...Option) (result LabelTag, err error) {
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
func (obj *_LabelTagMgr) GetByOptions(opts ...Option) (results []*LabelTag, err error) {
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


// GetFromTagID 通过tag_id获取内容 标签主键 
func (obj *_LabelTagMgr) GetFromTagID(tagID int) (results []*LabelTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("tag_id = ?", tagID).Find(&results).Error
	
	return
}

// GetBatchFromTagID 批量唯一主键查找 标签主键
func (obj *_LabelTagMgr) GetBatchFromTagID(tagIDs []int) (results []*LabelTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("tag_id IN (?)", tagIDs).Find(&results).Error
	
	return
}
 
// GetFromTagUUID 通过tag_uuid获取内容 唯一标识 
func (obj *_LabelTagMgr) GetFromTagUUID(tagUUID string) (results []*LabelTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("tag_uuid = ?", tagUUID).Find(&results).Error
	
	return
}

// GetBatchFromTagUUID 批量唯一主键查找 唯一标识
func (obj *_LabelTagMgr) GetBatchFromTagUUID(tagUUIDs []string) (results []*LabelTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("tag_uuid IN (?)", tagUUIDs).Find(&results).Error
	
	return
}
 
// GetFromTagName 通过tag_name获取内容 标签名称 
func (obj *_LabelTagMgr) GetFromTagName(tagName string) (results []*LabelTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("tag_name = ?", tagName).Find(&results).Error
	
	return
}

// GetBatchFromTagName 批量唯一主键查找 标签名称
func (obj *_LabelTagMgr) GetBatchFromTagName(tagNames []string) (results []*LabelTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("tag_name IN (?)", tagNames).Find(&results).Error
	
	return
}
 
// GetFromTagType 通过tag_type获取内容 标签类型 
func (obj *_LabelTagMgr) GetFromTagType(tagType int) (results []*LabelTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("tag_type = ?", tagType).Find(&results).Error
	
	return
}

// GetBatchFromTagType 批量唯一主键查找 标签类型
func (obj *_LabelTagMgr) GetBatchFromTagType(tagTypes []int) (results []*LabelTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("tag_type IN (?)", tagTypes).Find(&results).Error
	
	return
}
 
// GetFromTagParentID 通过tag_parent_id获取内容 父级标识 
func (obj *_LabelTagMgr) GetFromTagParentID(tagParentID string) (results []*LabelTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("tag_parent_id = ?", tagParentID).Find(&results).Error
	
	return
}

// GetBatchFromTagParentID 批量唯一主键查找 父级标识
func (obj *_LabelTagMgr) GetBatchFromTagParentID(tagParentIDs []string) (results []*LabelTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("tag_parent_id IN (?)", tagParentIDs).Find(&results).Error
	
	return
}
 
// GetFromTagLayer 通过tag_layer获取内容 层级 
func (obj *_LabelTagMgr) GetFromTagLayer(tagLayer int) (results []*LabelTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("tag_layer = ?", tagLayer).Find(&results).Error
	
	return
}

// GetBatchFromTagLayer 批量唯一主键查找 层级
func (obj *_LabelTagMgr) GetBatchFromTagLayer(tagLayers []int) (results []*LabelTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("tag_layer IN (?)", tagLayers).Find(&results).Error
	
	return
}
 
// GetFromTagImg 通过tag_img获取内容 图片 
func (obj *_LabelTagMgr) GetFromTagImg(tagImg string) (results []*LabelTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("tag_img = ?", tagImg).Find(&results).Error
	
	return
}

// GetBatchFromTagImg 批量唯一主键查找 图片
func (obj *_LabelTagMgr) GetBatchFromTagImg(tagImgs []string) (results []*LabelTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("tag_img IN (?)", tagImgs).Find(&results).Error
	
	return
}
 
// GetFromTagImgAliOssPath 通过tag_img_ali_oss_path获取内容 图片在阿里云OSS上地址 
func (obj *_LabelTagMgr) GetFromTagImgAliOssPath(tagImgAliOssPath string) (results []*LabelTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("tag_img_ali_oss_path = ?", tagImgAliOssPath).Find(&results).Error
	
	return
}

// GetBatchFromTagImgAliOssPath 批量唯一主键查找 图片在阿里云OSS上地址
func (obj *_LabelTagMgr) GetBatchFromTagImgAliOssPath(tagImgAliOssPaths []string) (results []*LabelTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("tag_img_ali_oss_path IN (?)", tagImgAliOssPaths).Find(&results).Error
	
	return
}
 
// GetFromApplyStatus 通过apply_status获取内容 是否启用,0:启用,1:不启用 
func (obj *_LabelTagMgr) GetFromApplyStatus(applyStatus int) (results []*LabelTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("apply_status = ?", applyStatus).Find(&results).Error
	
	return
}

// GetBatchFromApplyStatus 批量唯一主键查找 是否启用,0:启用,1:不启用
func (obj *_LabelTagMgr) GetBatchFromApplyStatus(applyStatuss []int) (results []*LabelTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("apply_status IN (?)", applyStatuss).Find(&results).Error
	
	return
}
 
// GetFromDeleteStatus 通过delete_status获取内容 是否删除,0:未删除,1:已删除 
func (obj *_LabelTagMgr) GetFromDeleteStatus(deleteStatus int) (results []*LabelTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("delete_status = ?", deleteStatus).Find(&results).Error
	
	return
}

// GetBatchFromDeleteStatus 批量唯一主键查找 是否删除,0:未删除,1:已删除
func (obj *_LabelTagMgr) GetBatchFromDeleteStatus(deleteStatuss []int) (results []*LabelTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("delete_status IN (?)", deleteStatuss).Find(&results).Error
	
	return
}
 
// GetFromCreateTime 通过create_time获取内容 创建时间 
func (obj *_LabelTagMgr) GetFromCreateTime(createTime time.Time) (results []*LabelTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_time = ?", createTime).Find(&results).Error
	
	return
}

// GetBatchFromCreateTime 批量唯一主键查找 创建时间
func (obj *_LabelTagMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*LabelTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_time IN (?)", createTimes).Find(&results).Error
	
	return
}
 
// GetFromUpdateTime 通过update_time获取内容 修改时间 
func (obj *_LabelTagMgr) GetFromUpdateTime(updateTime time.Time) (results []*LabelTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_time = ?", updateTime).Find(&results).Error
	
	return
}

// GetBatchFromUpdateTime 批量唯一主键查找 修改时间
func (obj *_LabelTagMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*LabelTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_time IN (?)", updateTimes).Find(&results).Error
	
	return
}
 
// GetFromCreateUser 通过create_user获取内容 创建人 
func (obj *_LabelTagMgr) GetFromCreateUser(createUser string) (results []*LabelTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_user = ?", createUser).Find(&results).Error
	
	return
}

// GetBatchFromCreateUser 批量唯一主键查找 创建人
func (obj *_LabelTagMgr) GetBatchFromCreateUser(createUsers []string) (results []*LabelTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_user IN (?)", createUsers).Find(&results).Error
	
	return
}
 
// GetFromUpdateUser 通过update_user获取内容 修改人 
func (obj *_LabelTagMgr) GetFromUpdateUser(updateUser string) (results []*LabelTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_user = ?", updateUser).Find(&results).Error
	
	return
}

// GetBatchFromUpdateUser 批量唯一主键查找 修改人
func (obj *_LabelTagMgr) GetBatchFromUpdateUser(updateUsers []string) (results []*LabelTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_user IN (?)", updateUsers).Find(&results).Error
	
	return
}
 
// GetFromOrderNum 通过order_num获取内容 排序序号 
func (obj *_LabelTagMgr) GetFromOrderNum(orderNum int) (results []*LabelTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("order_num = ?", orderNum).Find(&results).Error
	
	return
}

// GetBatchFromOrderNum 批量唯一主键查找 排序序号
func (obj *_LabelTagMgr) GetBatchFromOrderNum(orderNums []int) (results []*LabelTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("order_num IN (?)", orderNums).Find(&results).Error
	
	return
}
 
// GetFromUploadUUID 通过upload_uuid获取内容  
func (obj *_LabelTagMgr) GetFromUploadUUID(uploadUUID string) (results []*LabelTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("upload_uuid = ?", uploadUUID).Find(&results).Error
	
	return
}

// GetBatchFromUploadUUID 批量唯一主键查找 
func (obj *_LabelTagMgr) GetBatchFromUploadUUID(uploadUUIDs []string) (results []*LabelTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("upload_uuid IN (?)", uploadUUIDs).Find(&results).Error
	
	return
}
 
// GetFromIsHaveChildren 通过is_have_children获取内容 是否可以有下级,0:可以,1:不可以 
func (obj *_LabelTagMgr) GetFromIsHaveChildren(isHaveChildren int) (results []*LabelTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("is_have_children = ?", isHaveChildren).Find(&results).Error
	
	return
}

// GetBatchFromIsHaveChildren 批量唯一主键查找 是否可以有下级,0:可以,1:不可以
func (obj *_LabelTagMgr) GetBatchFromIsHaveChildren(isHaveChildrens []int) (results []*LabelTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("is_have_children IN (?)", isHaveChildrens).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_LabelTagMgr) FetchByPrimaryKey(tagID int ,tagUUID string ) (result LabelTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("tag_id = ? AND tag_uuid = ?", tagID , tagUUID).Find(&result).Error
	
	return
}
 

 
 // FetchIndexByIndexTagID  获取多个内容
 func (obj *_LabelTagMgr) FetchIndexByIndexTagID(tagID int ) (results []*LabelTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("tag_id = ?", tagID).Find(&results).Error
	
	return
}
 
 // FetchIndexByIndexTagUUID  获取多个内容
 func (obj *_LabelTagMgr) FetchIndexByIndexTagUUID(tagUUID string ) (results []*LabelTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("tag_uuid = ?", tagUUID).Find(&results).Error
	
	return
}
 
 // FetchIndexByIndexTagParentID  获取多个内容
 func (obj *_LabelTagMgr) FetchIndexByIndexTagParentID(tagParentID string ) (results []*LabelTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("tag_parent_id = ?", tagParentID).Find(&results).Error
	
	return
}
 

	

