package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _SysLogininforMgr struct {
	*_BaseMgr
}

// SysLogininforMgr open func
func SysLogininforMgr(db *gorm.DB) *_SysLogininforMgr {
	if db == nil {
		panic(fmt.Errorf("SysLogininforMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SysLogininforMgr{_BaseMgr: &_BaseMgr{DB: db.Table("sys_logininfor"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SysLogininforMgr) GetTableName() string {
	return "sys_logininfor"
}

// Get 获取
func (obj *_SysLogininforMgr) Get() (result SysLogininfor, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SysLogininforMgr) Gets() (results []*SysLogininfor, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithInfoID info_id获取 访问ID
func (obj *_SysLogininforMgr) WithInfoID(infoID int64) Option {
	return optionFunc(func(o *options) { o.query["info_id"] = infoID })
}

// WithLoginName login_name获取 登录账号
func (obj *_SysLogininforMgr) WithLoginName(loginName string) Option {
	return optionFunc(func(o *options) { o.query["login_name"] = loginName })
}

// WithIPaddr ipaddr获取 登录IP地址
func (obj *_SysLogininforMgr) WithIPaddr(ipaddr string) Option {
	return optionFunc(func(o *options) { o.query["ipaddr"] = ipaddr })
}

// WithLoginLocation login_location获取 登录地点
func (obj *_SysLogininforMgr) WithLoginLocation(loginLocation string) Option {
	return optionFunc(func(o *options) { o.query["login_location"] = loginLocation })
}

// WithBrowser browser获取 浏览器类型
func (obj *_SysLogininforMgr) WithBrowser(browser string) Option {
	return optionFunc(func(o *options) { o.query["browser"] = browser })
}

// WithOs os获取 操作系统
func (obj *_SysLogininforMgr) WithOs(os string) Option {
	return optionFunc(func(o *options) { o.query["os"] = os })
}

// WithStatus status获取 登录状态（0成功 1失败）
func (obj *_SysLogininforMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithMsg msg获取 提示消息
func (obj *_SysLogininforMgr) WithMsg(msg string) Option {
	return optionFunc(func(o *options) { o.query["msg"] = msg })
}

// WithLoginTime login_time获取 访问时间
func (obj *_SysLogininforMgr) WithLoginTime(loginTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["login_time"] = loginTime })
}

// GetByOption 功能选项模式获取
func (obj *_SysLogininforMgr) GetByOption(opts ...Option) (result SysLogininfor, err error) {
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
func (obj *_SysLogininforMgr) GetByOptions(opts ...Option) (results []*SysLogininfor, err error) {
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

// GetFromInfoID 通过info_id获取内容 访问ID
func (obj *_SysLogininforMgr) GetFromInfoID(infoID int64) (result SysLogininfor, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("info_id = ?", infoID).Find(&result).Error

	return
}

// GetBatchFromInfoID 批量唯一主键查找 访问ID
func (obj *_SysLogininforMgr) GetBatchFromInfoID(infoIDs []int64) (results []*SysLogininfor, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("info_id IN (?)", infoIDs).Find(&results).Error

	return
}

// GetFromLoginName 通过login_name获取内容 登录账号
func (obj *_SysLogininforMgr) GetFromLoginName(loginName string) (results []*SysLogininfor, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("login_name = ?", loginName).Find(&results).Error

	return
}

// GetBatchFromLoginName 批量唯一主键查找 登录账号
func (obj *_SysLogininforMgr) GetBatchFromLoginName(loginNames []string) (results []*SysLogininfor, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("login_name IN (?)", loginNames).Find(&results).Error

	return
}

// GetFromIPaddr 通过ipaddr获取内容 登录IP地址
func (obj *_SysLogininforMgr) GetFromIPaddr(ipaddr string) (results []*SysLogininfor, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ipaddr = ?", ipaddr).Find(&results).Error

	return
}

// GetBatchFromIPaddr 批量唯一主键查找 登录IP地址
func (obj *_SysLogininforMgr) GetBatchFromIPaddr(ipaddrs []string) (results []*SysLogininfor, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ipaddr IN (?)", ipaddrs).Find(&results).Error

	return
}

// GetFromLoginLocation 通过login_location获取内容 登录地点
func (obj *_SysLogininforMgr) GetFromLoginLocation(loginLocation string) (results []*SysLogininfor, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("login_location = ?", loginLocation).Find(&results).Error

	return
}

// GetBatchFromLoginLocation 批量唯一主键查找 登录地点
func (obj *_SysLogininforMgr) GetBatchFromLoginLocation(loginLocations []string) (results []*SysLogininfor, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("login_location IN (?)", loginLocations).Find(&results).Error

	return
}

// GetFromBrowser 通过browser获取内容 浏览器类型
func (obj *_SysLogininforMgr) GetFromBrowser(browser string) (results []*SysLogininfor, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("browser = ?", browser).Find(&results).Error

	return
}

// GetBatchFromBrowser 批量唯一主键查找 浏览器类型
func (obj *_SysLogininforMgr) GetBatchFromBrowser(browsers []string) (results []*SysLogininfor, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("browser IN (?)", browsers).Find(&results).Error

	return
}

// GetFromOs 通过os获取内容 操作系统
func (obj *_SysLogininforMgr) GetFromOs(os string) (results []*SysLogininfor, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("os = ?", os).Find(&results).Error

	return
}

// GetBatchFromOs 批量唯一主键查找 操作系统
func (obj *_SysLogininforMgr) GetBatchFromOs(oss []string) (results []*SysLogininfor, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("os IN (?)", oss).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 登录状态（0成功 1失败）
func (obj *_SysLogininforMgr) GetFromStatus(status string) (results []*SysLogininfor, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量唯一主键查找 登录状态（0成功 1失败）
func (obj *_SysLogininforMgr) GetBatchFromStatus(statuss []string) (results []*SysLogininfor, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status IN (?)", statuss).Find(&results).Error

	return
}

// GetFromMsg 通过msg获取内容 提示消息
func (obj *_SysLogininforMgr) GetFromMsg(msg string) (results []*SysLogininfor, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("msg = ?", msg).Find(&results).Error

	return
}

// GetBatchFromMsg 批量唯一主键查找 提示消息
func (obj *_SysLogininforMgr) GetBatchFromMsg(msgs []string) (results []*SysLogininfor, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("msg IN (?)", msgs).Find(&results).Error

	return
}

// GetFromLoginTime 通过login_time获取内容 访问时间
func (obj *_SysLogininforMgr) GetFromLoginTime(loginTime time.Time) (results []*SysLogininfor, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("login_time = ?", loginTime).Find(&results).Error

	return
}

// GetBatchFromLoginTime 批量唯一主键查找 访问时间
func (obj *_SysLogininforMgr) GetBatchFromLoginTime(loginTimes []time.Time) (results []*SysLogininfor, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("login_time IN (?)", loginTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_SysLogininforMgr) FetchByPrimaryKey(infoID int64) (result SysLogininfor, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("info_id = ?", infoID).Find(&result).Error

	return
}
