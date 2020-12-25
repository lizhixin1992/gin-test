package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _LabelCharacterMgr struct {
	*_BaseMgr
}

// LabelCharacterMgr open func
func LabelCharacterMgr(db *gorm.DB) *_LabelCharacterMgr {
	if db == nil {
		panic(fmt.Errorf("LabelCharacterMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_LabelCharacterMgr{_BaseMgr: &_BaseMgr{DB: db.Table("label_character"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_LabelCharacterMgr) GetTableName() string {
	return "label_character"
}

// Get 获取
func (obj *_LabelCharacterMgr) Get() (result LabelCharacter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_LabelCharacterMgr) Gets() (results []*LabelCharacter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithCharacterID character_id获取 人物主键
func (obj *_LabelCharacterMgr) WithCharacterID(characterID int) Option {
	return optionFunc(func(o *options) { o.query["character_id"] = characterID })
}

// WithCharacterUUID character_uuid获取 唯一标识
func (obj *_LabelCharacterMgr) WithCharacterUUID(characterUUID string) Option {
	return optionFunc(func(o *options) { o.query["character_uuid"] = characterUUID })
}

// WithCharacterName character_name获取 人物名称
func (obj *_LabelCharacterMgr) WithCharacterName(characterName string) Option {
	return optionFunc(func(o *options) { o.query["character_name"] = characterName })
}

// WithCharacterNamePinyinFirst character_name_pinyin_first获取 人物名称首字母
func (obj *_LabelCharacterMgr) WithCharacterNamePinyinFirst(characterNamePinyinFirst string) Option {
	return optionFunc(func(o *options) { o.query["character_name_pinyin_first"] = characterNamePinyinFirst })
}

// WithCharacterType character_type获取 人物类型
func (obj *_LabelCharacterMgr) WithCharacterType(characterType int) Option {
	return optionFunc(func(o *options) { o.query["character_type"] = characterType })
}

// WithCharacterImg character_img获取 图片
func (obj *_LabelCharacterMgr) WithCharacterImg(characterImg string) Option {
	return optionFunc(func(o *options) { o.query["character_img"] = characterImg })
}

// WithCharacterImgAliOssPath character_img_ali_oss_path获取 图片在阿里云OSS上地址
func (obj *_LabelCharacterMgr) WithCharacterImgAliOssPath(characterImgAliOssPath string) Option {
	return optionFunc(func(o *options) { o.query["character_img_ali_oss_path"] = characterImgAliOssPath })
}

// WithCharacterSex character_sex获取 性别,0:男,1:女
func (obj *_LabelCharacterMgr) WithCharacterSex(characterSex int) Option {
	return optionFunc(func(o *options) { o.query["character_sex"] = characterSex })
}

// WithCharacterBirthday character_birthday获取 出生日期
func (obj *_LabelCharacterMgr) WithCharacterBirthday(characterBirthday time.Time) Option {
	return optionFunc(func(o *options) { o.query["character_birthday"] = characterBirthday })
}

// WithCharacterCountry character_country获取 国家
func (obj *_LabelCharacterMgr) WithCharacterCountry(characterCountry string) Option {
	return optionFunc(func(o *options) { o.query["character_country"] = characterCountry })
}

// WithApplyStatus apply_status获取 是否启用,0:启用,1:不启用
func (obj *_LabelCharacterMgr) WithApplyStatus(applyStatus int) Option {
	return optionFunc(func(o *options) { o.query["apply_status"] = applyStatus })
}

// WithDeleteStatus delete_status获取 是否删除,0:未删除,1:已删除
func (obj *_LabelCharacterMgr) WithDeleteStatus(deleteStatus int) Option {
	return optionFunc(func(o *options) { o.query["delete_status"] = deleteStatus })
}

// WithCreateTime create_time获取 创建时间
func (obj *_LabelCharacterMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateTime update_time获取 修改时间
func (obj *_LabelCharacterMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithCreateUser create_user获取 创建人
func (obj *_LabelCharacterMgr) WithCreateUser(createUser string) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateUser update_user获取 修改人
func (obj *_LabelCharacterMgr) WithUpdateUser(updateUser string) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithUploadUUID upload_uuid获取
func (obj *_LabelCharacterMgr) WithUploadUUID(uploadUUID string) Option {
	return optionFunc(func(o *options) { o.query["upload_uuid"] = uploadUUID })
}

// GetByOption 功能选项模式获取
func (obj *_LabelCharacterMgr) GetByOption(opts ...Option) (result LabelCharacter, err error) {
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
func (obj *_LabelCharacterMgr) GetByOptions(opts ...Option) (results []*LabelCharacter, err error) {
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

// GetFromCharacterID 通过character_id获取内容 人物主键
func (obj *_LabelCharacterMgr) GetFromCharacterID(characterID int) (results []*LabelCharacter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("character_id = ?", characterID).Find(&results).Error

	return
}

// GetBatchFromCharacterID 批量唯一主键查找 人物主键
func (obj *_LabelCharacterMgr) GetBatchFromCharacterID(characterIDs []int) (results []*LabelCharacter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("character_id IN (?)", characterIDs).Find(&results).Error

	return
}

// GetFromCharacterUUID 通过character_uuid获取内容 唯一标识
func (obj *_LabelCharacterMgr) GetFromCharacterUUID(characterUUID string) (results []*LabelCharacter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("character_uuid = ?", characterUUID).Find(&results).Error

	return
}

// GetBatchFromCharacterUUID 批量唯一主键查找 唯一标识
func (obj *_LabelCharacterMgr) GetBatchFromCharacterUUID(characterUUIDs []string) (results []*LabelCharacter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("character_uuid IN (?)", characterUUIDs).Find(&results).Error

	return
}

// GetFromCharacterName 通过character_name获取内容 人物名称
func (obj *_LabelCharacterMgr) GetFromCharacterName(characterName string) (results []*LabelCharacter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("character_name = ?", characterName).Find(&results).Error

	return
}

// GetBatchFromCharacterName 批量唯一主键查找 人物名称
func (obj *_LabelCharacterMgr) GetBatchFromCharacterName(characterNames []string) (results []*LabelCharacter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("character_name IN (?)", characterNames).Find(&results).Error

	return
}

// GetFromCharacterNamePinyinFirst 通过character_name_pinyin_first获取内容 人物名称首字母
func (obj *_LabelCharacterMgr) GetFromCharacterNamePinyinFirst(characterNamePinyinFirst string) (results []*LabelCharacter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("character_name_pinyin_first = ?", characterNamePinyinFirst).Find(&results).Error

	return
}

// GetBatchFromCharacterNamePinyinFirst 批量唯一主键查找 人物名称首字母
func (obj *_LabelCharacterMgr) GetBatchFromCharacterNamePinyinFirst(characterNamePinyinFirsts []string) (results []*LabelCharacter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("character_name_pinyin_first IN (?)", characterNamePinyinFirsts).Find(&results).Error

	return
}

// GetFromCharacterType 通过character_type获取内容 人物类型
func (obj *_LabelCharacterMgr) GetFromCharacterType(characterType int) (results []*LabelCharacter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("character_type = ?", characterType).Find(&results).Error

	return
}

// GetBatchFromCharacterType 批量唯一主键查找 人物类型
func (obj *_LabelCharacterMgr) GetBatchFromCharacterType(characterTypes []int) (results []*LabelCharacter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("character_type IN (?)", characterTypes).Find(&results).Error

	return
}

// GetFromCharacterImg 通过character_img获取内容 图片
func (obj *_LabelCharacterMgr) GetFromCharacterImg(characterImg string) (results []*LabelCharacter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("character_img = ?", characterImg).Find(&results).Error

	return
}

// GetBatchFromCharacterImg 批量唯一主键查找 图片
func (obj *_LabelCharacterMgr) GetBatchFromCharacterImg(characterImgs []string) (results []*LabelCharacter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("character_img IN (?)", characterImgs).Find(&results).Error

	return
}

// GetFromCharacterImgAliOssPath 通过character_img_ali_oss_path获取内容 图片在阿里云OSS上地址
func (obj *_LabelCharacterMgr) GetFromCharacterImgAliOssPath(characterImgAliOssPath string) (results []*LabelCharacter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("character_img_ali_oss_path = ?", characterImgAliOssPath).Find(&results).Error

	return
}

// GetBatchFromCharacterImgAliOssPath 批量唯一主键查找 图片在阿里云OSS上地址
func (obj *_LabelCharacterMgr) GetBatchFromCharacterImgAliOssPath(characterImgAliOssPaths []string) (results []*LabelCharacter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("character_img_ali_oss_path IN (?)", characterImgAliOssPaths).Find(&results).Error

	return
}

// GetFromCharacterSex 通过character_sex获取内容 性别,0:男,1:女
func (obj *_LabelCharacterMgr) GetFromCharacterSex(characterSex int) (results []*LabelCharacter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("character_sex = ?", characterSex).Find(&results).Error

	return
}

// GetBatchFromCharacterSex 批量唯一主键查找 性别,0:男,1:女
func (obj *_LabelCharacterMgr) GetBatchFromCharacterSex(characterSexs []int) (results []*LabelCharacter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("character_sex IN (?)", characterSexs).Find(&results).Error

	return
}

// GetFromCharacterBirthday 通过character_birthday获取内容 出生日期
func (obj *_LabelCharacterMgr) GetFromCharacterBirthday(characterBirthday time.Time) (results []*LabelCharacter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("character_birthday = ?", characterBirthday).Find(&results).Error

	return
}

// GetBatchFromCharacterBirthday 批量唯一主键查找 出生日期
func (obj *_LabelCharacterMgr) GetBatchFromCharacterBirthday(characterBirthdays []time.Time) (results []*LabelCharacter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("character_birthday IN (?)", characterBirthdays).Find(&results).Error

	return
}

// GetFromCharacterCountry 通过character_country获取内容 国家
func (obj *_LabelCharacterMgr) GetFromCharacterCountry(characterCountry string) (results []*LabelCharacter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("character_country = ?", characterCountry).Find(&results).Error

	return
}

// GetBatchFromCharacterCountry 批量唯一主键查找 国家
func (obj *_LabelCharacterMgr) GetBatchFromCharacterCountry(characterCountrys []string) (results []*LabelCharacter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("character_country IN (?)", characterCountrys).Find(&results).Error

	return
}

// GetFromApplyStatus 通过apply_status获取内容 是否启用,0:启用,1:不启用
func (obj *_LabelCharacterMgr) GetFromApplyStatus(applyStatus int) (results []*LabelCharacter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("apply_status = ?", applyStatus).Find(&results).Error

	return
}

// GetBatchFromApplyStatus 批量唯一主键查找 是否启用,0:启用,1:不启用
func (obj *_LabelCharacterMgr) GetBatchFromApplyStatus(applyStatuss []int) (results []*LabelCharacter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("apply_status IN (?)", applyStatuss).Find(&results).Error

	return
}

// GetFromDeleteStatus 通过delete_status获取内容 是否删除,0:未删除,1:已删除
func (obj *_LabelCharacterMgr) GetFromDeleteStatus(deleteStatus int) (results []*LabelCharacter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("delete_status = ?", deleteStatus).Find(&results).Error

	return
}

// GetBatchFromDeleteStatus 批量唯一主键查找 是否删除,0:未删除,1:已删除
func (obj *_LabelCharacterMgr) GetBatchFromDeleteStatus(deleteStatuss []int) (results []*LabelCharacter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("delete_status IN (?)", deleteStatuss).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_LabelCharacterMgr) GetFromCreateTime(createTime time.Time) (results []*LabelCharacter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_time = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量唯一主键查找 创建时间
func (obj *_LabelCharacterMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*LabelCharacter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_time IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 修改时间
func (obj *_LabelCharacterMgr) GetFromUpdateTime(updateTime time.Time) (results []*LabelCharacter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_time = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量唯一主键查找 修改时间
func (obj *_LabelCharacterMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*LabelCharacter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_time IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建人
func (obj *_LabelCharacterMgr) GetFromCreateUser(createUser string) (results []*LabelCharacter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_user = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量唯一主键查找 创建人
func (obj *_LabelCharacterMgr) GetBatchFromCreateUser(createUsers []string) (results []*LabelCharacter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_user IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 修改人
func (obj *_LabelCharacterMgr) GetFromUpdateUser(updateUser string) (results []*LabelCharacter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_user = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量唯一主键查找 修改人
func (obj *_LabelCharacterMgr) GetBatchFromUpdateUser(updateUsers []string) (results []*LabelCharacter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_user IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromUploadUUID 通过upload_uuid获取内容
func (obj *_LabelCharacterMgr) GetFromUploadUUID(uploadUUID string) (results []*LabelCharacter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("upload_uuid = ?", uploadUUID).Find(&results).Error

	return
}

// GetBatchFromUploadUUID 批量唯一主键查找
func (obj *_LabelCharacterMgr) GetBatchFromUploadUUID(uploadUUIDs []string) (results []*LabelCharacter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("upload_uuid IN (?)", uploadUUIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_LabelCharacterMgr) FetchByPrimaryKey(characterID int, characterUUID string) (result LabelCharacter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("character_id = ? AND character_uuid = ?", characterID, characterUUID).Find(&result).Error

	return
}
