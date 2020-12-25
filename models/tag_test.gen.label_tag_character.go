package	models	
import (	
"context"	
"gorm.io/gorm"	
"time"	
"fmt"	
)	

type _LabelTagCharacterMgr struct {
	*_BaseMgr
}

// LabelTagCharacterMgr open func
func LabelTagCharacterMgr(db *gorm.DB) *_LabelTagCharacterMgr {
	if db == nil {
		panic(fmt.Errorf("LabelTagCharacterMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_LabelTagCharacterMgr{_BaseMgr: &_BaseMgr{DB: db.Table("label_tag_character"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_LabelTagCharacterMgr) GetTableName() string {
	return "label_tag_character"
}

// Get 获取
func (obj *_LabelTagCharacterMgr) Get() (result LabelTagCharacter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_LabelTagCharacterMgr) Gets() (results []*LabelTagCharacter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithRelationUUID relation_uuid获取 唯一标识
func (obj *_LabelTagCharacterMgr) WithRelationUUID(relationUUID string) Option {
	return optionFunc(func(o *options) { o.query["relation_uuid"] = relationUUID })
}

// WithTagUUID tag_uuid获取 标签标识
func (obj *_LabelTagCharacterMgr) WithTagUUID(tagUUID string) Option {
	return optionFunc(func(o *options) { o.query["tag_uuid"] = tagUUID })
}

// WithCharacterUUID character_uuid获取 人物标识
func (obj *_LabelTagCharacterMgr) WithCharacterUUID(characterUUID string) Option {
	return optionFunc(func(o *options) { o.query["character_uuid"] = characterUUID })
}

// WithCreateTime create_time获取 创建时间
func (obj *_LabelTagCharacterMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateTime update_time获取 修改时间
func (obj *_LabelTagCharacterMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithCreateUser create_user获取 创建人
func (obj *_LabelTagCharacterMgr) WithCreateUser(createUser string) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateUser update_user获取 修改人
func (obj *_LabelTagCharacterMgr) WithUpdateUser(updateUser string) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithDeleteStatus delete_status获取 是否删除,0:未删除,1:已删除
func (obj *_LabelTagCharacterMgr) WithDeleteStatus(deleteStatus int) Option {
	return optionFunc(func(o *options) { o.query["delete_status"] = deleteStatus })
}


// GetByOption 功能选项模式获取
func (obj *_LabelTagCharacterMgr) GetByOption(opts ...Option) (result LabelTagCharacter, err error) {
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
func (obj *_LabelTagCharacterMgr) GetByOptions(opts ...Option) (results []*LabelTagCharacter, err error) {
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


// GetFromRelationUUID 通过relation_uuid获取内容 唯一标识 
func (obj *_LabelTagCharacterMgr) GetFromRelationUUID(relationUUID string) (results []*LabelTagCharacter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("relation_uuid = ?", relationUUID).Find(&results).Error
	
	return
}

// GetBatchFromRelationUUID 批量唯一主键查找 唯一标识
func (obj *_LabelTagCharacterMgr) GetBatchFromRelationUUID(relationUUIDs []string) (results []*LabelTagCharacter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("relation_uuid IN (?)", relationUUIDs).Find(&results).Error
	
	return
}
 
// GetFromTagUUID 通过tag_uuid获取内容 标签标识 
func (obj *_LabelTagCharacterMgr) GetFromTagUUID(tagUUID string) (results []*LabelTagCharacter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("tag_uuid = ?", tagUUID).Find(&results).Error
	
	return
}

// GetBatchFromTagUUID 批量唯一主键查找 标签标识
func (obj *_LabelTagCharacterMgr) GetBatchFromTagUUID(tagUUIDs []string) (results []*LabelTagCharacter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("tag_uuid IN (?)", tagUUIDs).Find(&results).Error
	
	return
}
 
// GetFromCharacterUUID 通过character_uuid获取内容 人物标识 
func (obj *_LabelTagCharacterMgr) GetFromCharacterUUID(characterUUID string) (results []*LabelTagCharacter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("character_uuid = ?", characterUUID).Find(&results).Error
	
	return
}

// GetBatchFromCharacterUUID 批量唯一主键查找 人物标识
func (obj *_LabelTagCharacterMgr) GetBatchFromCharacterUUID(characterUUIDs []string) (results []*LabelTagCharacter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("character_uuid IN (?)", characterUUIDs).Find(&results).Error
	
	return
}
 
// GetFromCreateTime 通过create_time获取内容 创建时间 
func (obj *_LabelTagCharacterMgr) GetFromCreateTime(createTime time.Time) (results []*LabelTagCharacter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_time = ?", createTime).Find(&results).Error
	
	return
}

// GetBatchFromCreateTime 批量唯一主键查找 创建时间
func (obj *_LabelTagCharacterMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*LabelTagCharacter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_time IN (?)", createTimes).Find(&results).Error
	
	return
}
 
// GetFromUpdateTime 通过update_time获取内容 修改时间 
func (obj *_LabelTagCharacterMgr) GetFromUpdateTime(updateTime time.Time) (results []*LabelTagCharacter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_time = ?", updateTime).Find(&results).Error
	
	return
}

// GetBatchFromUpdateTime 批量唯一主键查找 修改时间
func (obj *_LabelTagCharacterMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*LabelTagCharacter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_time IN (?)", updateTimes).Find(&results).Error
	
	return
}
 
// GetFromCreateUser 通过create_user获取内容 创建人 
func (obj *_LabelTagCharacterMgr) GetFromCreateUser(createUser string) (results []*LabelTagCharacter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_user = ?", createUser).Find(&results).Error
	
	return
}

// GetBatchFromCreateUser 批量唯一主键查找 创建人
func (obj *_LabelTagCharacterMgr) GetBatchFromCreateUser(createUsers []string) (results []*LabelTagCharacter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_user IN (?)", createUsers).Find(&results).Error
	
	return
}
 
// GetFromUpdateUser 通过update_user获取内容 修改人 
func (obj *_LabelTagCharacterMgr) GetFromUpdateUser(updateUser string) (results []*LabelTagCharacter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_user = ?", updateUser).Find(&results).Error
	
	return
}

// GetBatchFromUpdateUser 批量唯一主键查找 修改人
func (obj *_LabelTagCharacterMgr) GetBatchFromUpdateUser(updateUsers []string) (results []*LabelTagCharacter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_user IN (?)", updateUsers).Find(&results).Error
	
	return
}
 
// GetFromDeleteStatus 通过delete_status获取内容 是否删除,0:未删除,1:已删除 
func (obj *_LabelTagCharacterMgr) GetFromDeleteStatus(deleteStatus int) (results []*LabelTagCharacter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("delete_status = ?", deleteStatus).Find(&results).Error
	
	return
}

// GetBatchFromDeleteStatus 批量唯一主键查找 是否删除,0:未删除,1:已删除
func (obj *_LabelTagCharacterMgr) GetBatchFromDeleteStatus(deleteStatuss []int) (results []*LabelTagCharacter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("delete_status IN (?)", deleteStatuss).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_LabelTagCharacterMgr) FetchByPrimaryKey(tagUUID string ,characterUUID string ) (result LabelTagCharacter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("tag_uuid = ? AND character_uuid = ?", tagUUID , characterUUID).Find(&result).Error
	
	return
}
 

 

	

