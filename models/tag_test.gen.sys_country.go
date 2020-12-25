package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _SysCountryMgr struct {
	*_BaseMgr
}

// SysCountryMgr open func
func SysCountryMgr(db *gorm.DB) *_SysCountryMgr {
	if db == nil {
		panic(fmt.Errorf("SysCountryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SysCountryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("sys_country"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SysCountryMgr) GetTableName() string {
	return "sys_country"
}

// Get 获取
func (obj *_SysCountryMgr) Get() (result SysCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SysCountryMgr) Gets() (results []*SysCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithCountryID country_id获取 主键
func (obj *_SysCountryMgr) WithCountryID(countryID int) Option {
	return optionFunc(func(o *options) { o.query["country_id"] = countryID })
}

// WithCountryNameChinese country_name_chinese获取 国家中文名称
func (obj *_SysCountryMgr) WithCountryNameChinese(countryNameChinese string) Option {
	return optionFunc(func(o *options) { o.query["country_name_chinese"] = countryNameChinese })
}

// WithCountryCode country_code获取 国家代码
func (obj *_SysCountryMgr) WithCountryCode(countryCode string) Option {
	return optionFunc(func(o *options) { o.query["country_code"] = countryCode })
}

// WithCountryNameEnglish country_name_english获取 国家英文名称
func (obj *_SysCountryMgr) WithCountryNameEnglish(countryNameEnglish string) Option {
	return optionFunc(func(o *options) { o.query["country_name_english"] = countryNameEnglish })
}

// WithCountryNationalFlag country_national_flag获取 国旗
func (obj *_SysCountryMgr) WithCountryNationalFlag(countryNationalFlag string) Option {
	return optionFunc(func(o *options) { o.query["country_national_flag"] = countryNationalFlag })
}

// WithCountryUUID country_uuid获取 唯一标识
func (obj *_SysCountryMgr) WithCountryUUID(countryUUID string) Option {
	return optionFunc(func(o *options) { o.query["country_uuid"] = countryUUID })
}

// WithApplyStatus apply_status获取 是否启用,0:启用,1:不启用
func (obj *_SysCountryMgr) WithApplyStatus(applyStatus int) Option {
	return optionFunc(func(o *options) { o.query["apply_status"] = applyStatus })
}

// WithDeleteStatus delete_status获取 是否删除,0:未删除,1:已删除
func (obj *_SysCountryMgr) WithDeleteStatus(deleteStatus int) Option {
	return optionFunc(func(o *options) { o.query["delete_status"] = deleteStatus })
}

// WithCreateTime create_time获取 创建时间
func (obj *_SysCountryMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateTime update_time获取 修改时间
func (obj *_SysCountryMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithCreateUser create_user获取 创建人
func (obj *_SysCountryMgr) WithCreateUser(createUser string) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateUser update_user获取 修改人
func (obj *_SysCountryMgr) WithUpdateUser(updateUser string) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithCountryType country_type获取 类型,0:国家;1:地区
func (obj *_SysCountryMgr) WithCountryType(countryType int) Option {
	return optionFunc(func(o *options) { o.query["country_type"] = countryType })
}

// WithBelongContinent belong_continent获取 所属大洲
func (obj *_SysCountryMgr) WithBelongContinent(belongContinent string) Option {
	return optionFunc(func(o *options) { o.query["belong_continent"] = belongContinent })
}

// GetByOption 功能选项模式获取
func (obj *_SysCountryMgr) GetByOption(opts ...Option) (result SysCountry, err error) {
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
func (obj *_SysCountryMgr) GetByOptions(opts ...Option) (results []*SysCountry, err error) {
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

// GetFromCountryID 通过country_id获取内容 主键
func (obj *_SysCountryMgr) GetFromCountryID(countryID int) (result SysCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("country_id = ?", countryID).Find(&result).Error

	return
}

// GetBatchFromCountryID 批量唯一主键查找 主键
func (obj *_SysCountryMgr) GetBatchFromCountryID(countryIDs []int) (results []*SysCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("country_id IN (?)", countryIDs).Find(&results).Error

	return
}

// GetFromCountryNameChinese 通过country_name_chinese获取内容 国家中文名称
func (obj *_SysCountryMgr) GetFromCountryNameChinese(countryNameChinese string) (results []*SysCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("country_name_chinese = ?", countryNameChinese).Find(&results).Error

	return
}

// GetBatchFromCountryNameChinese 批量唯一主键查找 国家中文名称
func (obj *_SysCountryMgr) GetBatchFromCountryNameChinese(countryNameChineses []string) (results []*SysCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("country_name_chinese IN (?)", countryNameChineses).Find(&results).Error

	return
}

// GetFromCountryCode 通过country_code获取内容 国家代码
func (obj *_SysCountryMgr) GetFromCountryCode(countryCode string) (results []*SysCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("country_code = ?", countryCode).Find(&results).Error

	return
}

// GetBatchFromCountryCode 批量唯一主键查找 国家代码
func (obj *_SysCountryMgr) GetBatchFromCountryCode(countryCodes []string) (results []*SysCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("country_code IN (?)", countryCodes).Find(&results).Error

	return
}

// GetFromCountryNameEnglish 通过country_name_english获取内容 国家英文名称
func (obj *_SysCountryMgr) GetFromCountryNameEnglish(countryNameEnglish string) (results []*SysCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("country_name_english = ?", countryNameEnglish).Find(&results).Error

	return
}

// GetBatchFromCountryNameEnglish 批量唯一主键查找 国家英文名称
func (obj *_SysCountryMgr) GetBatchFromCountryNameEnglish(countryNameEnglishs []string) (results []*SysCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("country_name_english IN (?)", countryNameEnglishs).Find(&results).Error

	return
}

// GetFromCountryNationalFlag 通过country_national_flag获取内容 国旗
func (obj *_SysCountryMgr) GetFromCountryNationalFlag(countryNationalFlag string) (results []*SysCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("country_national_flag = ?", countryNationalFlag).Find(&results).Error

	return
}

// GetBatchFromCountryNationalFlag 批量唯一主键查找 国旗
func (obj *_SysCountryMgr) GetBatchFromCountryNationalFlag(countryNationalFlags []string) (results []*SysCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("country_national_flag IN (?)", countryNationalFlags).Find(&results).Error

	return
}

// GetFromCountryUUID 通过country_uuid获取内容 唯一标识
func (obj *_SysCountryMgr) GetFromCountryUUID(countryUUID string) (results []*SysCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("country_uuid = ?", countryUUID).Find(&results).Error

	return
}

// GetBatchFromCountryUUID 批量唯一主键查找 唯一标识
func (obj *_SysCountryMgr) GetBatchFromCountryUUID(countryUUIDs []string) (results []*SysCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("country_uuid IN (?)", countryUUIDs).Find(&results).Error

	return
}

// GetFromApplyStatus 通过apply_status获取内容 是否启用,0:启用,1:不启用
func (obj *_SysCountryMgr) GetFromApplyStatus(applyStatus int) (results []*SysCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("apply_status = ?", applyStatus).Find(&results).Error

	return
}

// GetBatchFromApplyStatus 批量唯一主键查找 是否启用,0:启用,1:不启用
func (obj *_SysCountryMgr) GetBatchFromApplyStatus(applyStatuss []int) (results []*SysCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("apply_status IN (?)", applyStatuss).Find(&results).Error

	return
}

// GetFromDeleteStatus 通过delete_status获取内容 是否删除,0:未删除,1:已删除
func (obj *_SysCountryMgr) GetFromDeleteStatus(deleteStatus int) (results []*SysCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("delete_status = ?", deleteStatus).Find(&results).Error

	return
}

// GetBatchFromDeleteStatus 批量唯一主键查找 是否删除,0:未删除,1:已删除
func (obj *_SysCountryMgr) GetBatchFromDeleteStatus(deleteStatuss []int) (results []*SysCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("delete_status IN (?)", deleteStatuss).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_SysCountryMgr) GetFromCreateTime(createTime time.Time) (results []*SysCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_time = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量唯一主键查找 创建时间
func (obj *_SysCountryMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*SysCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_time IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 修改时间
func (obj *_SysCountryMgr) GetFromUpdateTime(updateTime time.Time) (results []*SysCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_time = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量唯一主键查找 修改时间
func (obj *_SysCountryMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*SysCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_time IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建人
func (obj *_SysCountryMgr) GetFromCreateUser(createUser string) (results []*SysCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_user = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量唯一主键查找 创建人
func (obj *_SysCountryMgr) GetBatchFromCreateUser(createUsers []string) (results []*SysCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_user IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 修改人
func (obj *_SysCountryMgr) GetFromUpdateUser(updateUser string) (results []*SysCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_user = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量唯一主键查找 修改人
func (obj *_SysCountryMgr) GetBatchFromUpdateUser(updateUsers []string) (results []*SysCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_user IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromCountryType 通过country_type获取内容 类型,0:国家;1:地区
func (obj *_SysCountryMgr) GetFromCountryType(countryType int) (results []*SysCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("country_type = ?", countryType).Find(&results).Error

	return
}

// GetBatchFromCountryType 批量唯一主键查找 类型,0:国家;1:地区
func (obj *_SysCountryMgr) GetBatchFromCountryType(countryTypes []int) (results []*SysCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("country_type IN (?)", countryTypes).Find(&results).Error

	return
}

// GetFromBelongContinent 通过belong_continent获取内容 所属大洲
func (obj *_SysCountryMgr) GetFromBelongContinent(belongContinent string) (results []*SysCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("belong_continent = ?", belongContinent).Find(&results).Error

	return
}

// GetBatchFromBelongContinent 批量唯一主键查找 所属大洲
func (obj *_SysCountryMgr) GetBatchFromBelongContinent(belongContinents []string) (results []*SysCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("belong_continent IN (?)", belongContinents).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_SysCountryMgr) FetchByPrimaryKey(countryID int) (result SysCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("country_id = ?", countryID).Find(&result).Error

	return
}
