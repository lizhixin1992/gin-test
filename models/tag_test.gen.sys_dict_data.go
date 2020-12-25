package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _SysDictDataMgr struct {
	*_BaseMgr
}

// SysDictDataMgr open func
func SysDictDataMgr(db *gorm.DB) *_SysDictDataMgr {
	if db == nil {
		panic(fmt.Errorf("SysDictDataMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SysDictDataMgr{_BaseMgr: &_BaseMgr{DB: db.Table("sys_dict_data"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SysDictDataMgr) GetTableName() string {
	return "sys_dict_data"
}

// Get 获取
func (obj *_SysDictDataMgr) Get() (result SysDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SysDictDataMgr) Gets() (results []*SysDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithDictCode dict_code获取 字典编码
func (obj *_SysDictDataMgr) WithDictCode(dictCode int64) Option {
	return optionFunc(func(o *options) { o.query["dict_code"] = dictCode })
}

// WithDictSort dict_sort获取 字典排序
func (obj *_SysDictDataMgr) WithDictSort(dictSort int) Option {
	return optionFunc(func(o *options) { o.query["dict_sort"] = dictSort })
}

// WithDictLabel dict_label获取 字典标签
func (obj *_SysDictDataMgr) WithDictLabel(dictLabel string) Option {
	return optionFunc(func(o *options) { o.query["dict_label"] = dictLabel })
}

// WithDictValue dict_value获取 字典键值
func (obj *_SysDictDataMgr) WithDictValue(dictValue string) Option {
	return optionFunc(func(o *options) { o.query["dict_value"] = dictValue })
}

// WithDictType dict_type获取 字典类型
func (obj *_SysDictDataMgr) WithDictType(dictType string) Option {
	return optionFunc(func(o *options) { o.query["dict_type"] = dictType })
}

// WithCSSClass css_class获取 样式属性（其他样式扩展）
func (obj *_SysDictDataMgr) WithCSSClass(cssClass string) Option {
	return optionFunc(func(o *options) { o.query["css_class"] = cssClass })
}

// WithListClass list_class获取 表格回显样式
func (obj *_SysDictDataMgr) WithListClass(listClass string) Option {
	return optionFunc(func(o *options) { o.query["list_class"] = listClass })
}

// WithIsDefault is_default获取 是否默认（Y是 N否）
func (obj *_SysDictDataMgr) WithIsDefault(isDefault string) Option {
	return optionFunc(func(o *options) { o.query["is_default"] = isDefault })
}

// WithStatus status获取 状态（0正常 1停用）
func (obj *_SysDictDataMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithCreateBy create_by获取 创建者
func (obj *_SysDictDataMgr) WithCreateBy(createBy string) Option {
	return optionFunc(func(o *options) { o.query["create_by"] = createBy })
}

// WithCreateTime create_time获取 创建时间
func (obj *_SysDictDataMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateBy update_by获取 更新者
func (obj *_SysDictDataMgr) WithUpdateBy(updateBy string) Option {
	return optionFunc(func(o *options) { o.query["update_by"] = updateBy })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_SysDictDataMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithRemark remark获取 备注
func (obj *_SysDictDataMgr) WithRemark(remark string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// GetByOption 功能选项模式获取
func (obj *_SysDictDataMgr) GetByOption(opts ...Option) (result SysDictData, err error) {
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
func (obj *_SysDictDataMgr) GetByOptions(opts ...Option) (results []*SysDictData, err error) {
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

// GetFromDictCode 通过dict_code获取内容 字典编码
func (obj *_SysDictDataMgr) GetFromDictCode(dictCode int64) (result SysDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("dict_code = ?", dictCode).Find(&result).Error

	return
}

// GetBatchFromDictCode 批量唯一主键查找 字典编码
func (obj *_SysDictDataMgr) GetBatchFromDictCode(dictCodes []int64) (results []*SysDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("dict_code IN (?)", dictCodes).Find(&results).Error

	return
}

// GetFromDictSort 通过dict_sort获取内容 字典排序
func (obj *_SysDictDataMgr) GetFromDictSort(dictSort int) (results []*SysDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("dict_sort = ?", dictSort).Find(&results).Error

	return
}

// GetBatchFromDictSort 批量唯一主键查找 字典排序
func (obj *_SysDictDataMgr) GetBatchFromDictSort(dictSorts []int) (results []*SysDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("dict_sort IN (?)", dictSorts).Find(&results).Error

	return
}

// GetFromDictLabel 通过dict_label获取内容 字典标签
func (obj *_SysDictDataMgr) GetFromDictLabel(dictLabel string) (results []*SysDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("dict_label = ?", dictLabel).Find(&results).Error

	return
}

// GetBatchFromDictLabel 批量唯一主键查找 字典标签
func (obj *_SysDictDataMgr) GetBatchFromDictLabel(dictLabels []string) (results []*SysDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("dict_label IN (?)", dictLabels).Find(&results).Error

	return
}

// GetFromDictValue 通过dict_value获取内容 字典键值
func (obj *_SysDictDataMgr) GetFromDictValue(dictValue string) (results []*SysDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("dict_value = ?", dictValue).Find(&results).Error

	return
}

// GetBatchFromDictValue 批量唯一主键查找 字典键值
func (obj *_SysDictDataMgr) GetBatchFromDictValue(dictValues []string) (results []*SysDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("dict_value IN (?)", dictValues).Find(&results).Error

	return
}

// GetFromDictType 通过dict_type获取内容 字典类型
func (obj *_SysDictDataMgr) GetFromDictType(dictType string) (results []*SysDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("dict_type = ?", dictType).Find(&results).Error

	return
}

// GetBatchFromDictType 批量唯一主键查找 字典类型
func (obj *_SysDictDataMgr) GetBatchFromDictType(dictTypes []string) (results []*SysDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("dict_type IN (?)", dictTypes).Find(&results).Error

	return
}

// GetFromCSSClass 通过css_class获取内容 样式属性（其他样式扩展）
func (obj *_SysDictDataMgr) GetFromCSSClass(cssClass string) (results []*SysDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("css_class = ?", cssClass).Find(&results).Error

	return
}

// GetBatchFromCSSClass 批量唯一主键查找 样式属性（其他样式扩展）
func (obj *_SysDictDataMgr) GetBatchFromCSSClass(cssClasss []string) (results []*SysDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("css_class IN (?)", cssClasss).Find(&results).Error

	return
}

// GetFromListClass 通过list_class获取内容 表格回显样式
func (obj *_SysDictDataMgr) GetFromListClass(listClass string) (results []*SysDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("list_class = ?", listClass).Find(&results).Error

	return
}

// GetBatchFromListClass 批量唯一主键查找 表格回显样式
func (obj *_SysDictDataMgr) GetBatchFromListClass(listClasss []string) (results []*SysDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("list_class IN (?)", listClasss).Find(&results).Error

	return
}

// GetFromIsDefault 通过is_default获取内容 是否默认（Y是 N否）
func (obj *_SysDictDataMgr) GetFromIsDefault(isDefault string) (results []*SysDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("is_default = ?", isDefault).Find(&results).Error

	return
}

// GetBatchFromIsDefault 批量唯一主键查找 是否默认（Y是 N否）
func (obj *_SysDictDataMgr) GetBatchFromIsDefault(isDefaults []string) (results []*SysDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("is_default IN (?)", isDefaults).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 状态（0正常 1停用）
func (obj *_SysDictDataMgr) GetFromStatus(status string) (results []*SysDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量唯一主键查找 状态（0正常 1停用）
func (obj *_SysDictDataMgr) GetBatchFromStatus(statuss []string) (results []*SysDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status IN (?)", statuss).Find(&results).Error

	return
}

// GetFromCreateBy 通过create_by获取内容 创建者
func (obj *_SysDictDataMgr) GetFromCreateBy(createBy string) (results []*SysDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_by = ?", createBy).Find(&results).Error

	return
}

// GetBatchFromCreateBy 批量唯一主键查找 创建者
func (obj *_SysDictDataMgr) GetBatchFromCreateBy(createBys []string) (results []*SysDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_by IN (?)", createBys).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_SysDictDataMgr) GetFromCreateTime(createTime time.Time) (results []*SysDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_time = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量唯一主键查找 创建时间
func (obj *_SysDictDataMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*SysDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_time IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromUpdateBy 通过update_by获取内容 更新者
func (obj *_SysDictDataMgr) GetFromUpdateBy(updateBy string) (results []*SysDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_by = ?", updateBy).Find(&results).Error

	return
}

// GetBatchFromUpdateBy 批量唯一主键查找 更新者
func (obj *_SysDictDataMgr) GetBatchFromUpdateBy(updateBys []string) (results []*SysDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_by IN (?)", updateBys).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_SysDictDataMgr) GetFromUpdateTime(updateTime time.Time) (results []*SysDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_time = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量唯一主键查找 更新时间
func (obj *_SysDictDataMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*SysDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_time IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_SysDictDataMgr) GetFromRemark(remark string) (results []*SysDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("remark = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量唯一主键查找 备注
func (obj *_SysDictDataMgr) GetBatchFromRemark(remarks []string) (results []*SysDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("remark IN (?)", remarks).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_SysDictDataMgr) FetchByPrimaryKey(dictCode int64) (result SysDictData, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("dict_code = ?", dictCode).Find(&result).Error

	return
}
