package	models	
import (	
"context"	
"gorm.io/gorm"	
"time"	
"fmt"	
)	

type _SysConfigMgr struct {
	*_BaseMgr
}

// SysConfigMgr open func
func SysConfigMgr(db *gorm.DB) *_SysConfigMgr {
	if db == nil {
		panic(fmt.Errorf("SysConfigMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SysConfigMgr{_BaseMgr: &_BaseMgr{DB: db.Table("sys_config"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SysConfigMgr) GetTableName() string {
	return "sys_config"
}

// Get 获取
func (obj *_SysConfigMgr) Get() (result SysConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_SysConfigMgr) Gets() (results []*SysConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithConfigID config_id获取 参数主键
func (obj *_SysConfigMgr) WithConfigID(configID int) Option {
	return optionFunc(func(o *options) { o.query["config_id"] = configID })
}

// WithConfigName config_name获取 参数名称
func (obj *_SysConfigMgr) WithConfigName(configName string) Option {
	return optionFunc(func(o *options) { o.query["config_name"] = configName })
}

// WithConfigKey config_key获取 参数键名
func (obj *_SysConfigMgr) WithConfigKey(configKey string) Option {
	return optionFunc(func(o *options) { o.query["config_key"] = configKey })
}

// WithConfigValue config_value获取 参数键值
func (obj *_SysConfigMgr) WithConfigValue(configValue string) Option {
	return optionFunc(func(o *options) { o.query["config_value"] = configValue })
}

// WithConfigType config_type获取 系统内置（Y是 N否）
func (obj *_SysConfigMgr) WithConfigType(configType string) Option {
	return optionFunc(func(o *options) { o.query["config_type"] = configType })
}

// WithCreateBy create_by获取 创建者
func (obj *_SysConfigMgr) WithCreateBy(createBy string) Option {
	return optionFunc(func(o *options) { o.query["create_by"] = createBy })
}

// WithCreateTime create_time获取 创建时间
func (obj *_SysConfigMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateBy update_by获取 更新者
func (obj *_SysConfigMgr) WithUpdateBy(updateBy string) Option {
	return optionFunc(func(o *options) { o.query["update_by"] = updateBy })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_SysConfigMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithRemark remark获取 备注
func (obj *_SysConfigMgr) WithRemark(remark string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}


// GetByOption 功能选项模式获取
func (obj *_SysConfigMgr) GetByOption(opts ...Option) (result SysConfig, err error) {
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
func (obj *_SysConfigMgr) GetByOptions(opts ...Option) (results []*SysConfig, err error) {
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


// GetFromConfigID 通过config_id获取内容 参数主键 
func (obj *_SysConfigMgr)  GetFromConfigID(configID int) (result SysConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("config_id = ?", configID).Find(&result).Error
	
	return
}

// GetBatchFromConfigID 批量唯一主键查找 参数主键
func (obj *_SysConfigMgr) GetBatchFromConfigID(configIDs []int) (results []*SysConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("config_id IN (?)", configIDs).Find(&results).Error
	
	return
}
 
// GetFromConfigName 通过config_name获取内容 参数名称 
func (obj *_SysConfigMgr) GetFromConfigName(configName string) (results []*SysConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("config_name = ?", configName).Find(&results).Error
	
	return
}

// GetBatchFromConfigName 批量唯一主键查找 参数名称
func (obj *_SysConfigMgr) GetBatchFromConfigName(configNames []string) (results []*SysConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("config_name IN (?)", configNames).Find(&results).Error
	
	return
}
 
// GetFromConfigKey 通过config_key获取内容 参数键名 
func (obj *_SysConfigMgr) GetFromConfigKey(configKey string) (results []*SysConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("config_key = ?", configKey).Find(&results).Error
	
	return
}

// GetBatchFromConfigKey 批量唯一主键查找 参数键名
func (obj *_SysConfigMgr) GetBatchFromConfigKey(configKeys []string) (results []*SysConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("config_key IN (?)", configKeys).Find(&results).Error
	
	return
}
 
// GetFromConfigValue 通过config_value获取内容 参数键值 
func (obj *_SysConfigMgr) GetFromConfigValue(configValue string) (results []*SysConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("config_value = ?", configValue).Find(&results).Error
	
	return
}

// GetBatchFromConfigValue 批量唯一主键查找 参数键值
func (obj *_SysConfigMgr) GetBatchFromConfigValue(configValues []string) (results []*SysConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("config_value IN (?)", configValues).Find(&results).Error
	
	return
}
 
// GetFromConfigType 通过config_type获取内容 系统内置（Y是 N否） 
func (obj *_SysConfigMgr) GetFromConfigType(configType string) (results []*SysConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("config_type = ?", configType).Find(&results).Error
	
	return
}

// GetBatchFromConfigType 批量唯一主键查找 系统内置（Y是 N否）
func (obj *_SysConfigMgr) GetBatchFromConfigType(configTypes []string) (results []*SysConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("config_type IN (?)", configTypes).Find(&results).Error
	
	return
}
 
// GetFromCreateBy 通过create_by获取内容 创建者 
func (obj *_SysConfigMgr) GetFromCreateBy(createBy string) (results []*SysConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_by = ?", createBy).Find(&results).Error
	
	return
}

// GetBatchFromCreateBy 批量唯一主键查找 创建者
func (obj *_SysConfigMgr) GetBatchFromCreateBy(createBys []string) (results []*SysConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_by IN (?)", createBys).Find(&results).Error
	
	return
}
 
// GetFromCreateTime 通过create_time获取内容 创建时间 
func (obj *_SysConfigMgr) GetFromCreateTime(createTime time.Time) (results []*SysConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_time = ?", createTime).Find(&results).Error
	
	return
}

// GetBatchFromCreateTime 批量唯一主键查找 创建时间
func (obj *_SysConfigMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*SysConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_time IN (?)", createTimes).Find(&results).Error
	
	return
}
 
// GetFromUpdateBy 通过update_by获取内容 更新者 
func (obj *_SysConfigMgr) GetFromUpdateBy(updateBy string) (results []*SysConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_by = ?", updateBy).Find(&results).Error
	
	return
}

// GetBatchFromUpdateBy 批量唯一主键查找 更新者
func (obj *_SysConfigMgr) GetBatchFromUpdateBy(updateBys []string) (results []*SysConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_by IN (?)", updateBys).Find(&results).Error
	
	return
}
 
// GetFromUpdateTime 通过update_time获取内容 更新时间 
func (obj *_SysConfigMgr) GetFromUpdateTime(updateTime time.Time) (results []*SysConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_time = ?", updateTime).Find(&results).Error
	
	return
}

// GetBatchFromUpdateTime 批量唯一主键查找 更新时间
func (obj *_SysConfigMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*SysConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_time IN (?)", updateTimes).Find(&results).Error
	
	return
}
 
// GetFromRemark 通过remark获取内容 备注 
func (obj *_SysConfigMgr) GetFromRemark(remark string) (results []*SysConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("remark = ?", remark).Find(&results).Error
	
	return
}

// GetBatchFromRemark 批量唯一主键查找 备注
func (obj *_SysConfigMgr) GetBatchFromRemark(remarks []string) (results []*SysConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("remark IN (?)", remarks).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_SysConfigMgr) FetchByPrimaryKey(configID int ) (result SysConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("config_id = ?", configID).Find(&result).Error
	
	return
}
 

 

	

