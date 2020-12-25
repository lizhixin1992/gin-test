package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _SysUserOnlineMgr struct {
	*_BaseMgr
}

// SysUserOnlineMgr open func
func SysUserOnlineMgr(db *gorm.DB) *_SysUserOnlineMgr {
	if db == nil {
		panic(fmt.Errorf("SysUserOnlineMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SysUserOnlineMgr{_BaseMgr: &_BaseMgr{DB: db.Table("sys_user_online"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SysUserOnlineMgr) GetTableName() string {
	return "sys_user_online"
}

// Get 获取
func (obj *_SysUserOnlineMgr) Get() (result SysUserOnline, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SysUserOnlineMgr) Gets() (results []*SysUserOnline, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithSessionID sessionId获取 用户会话id
func (obj *_SysUserOnlineMgr) WithSessionID(sessionID string) Option {
	return optionFunc(func(o *options) { o.query["sessionId"] = sessionID })
}

// WithLoginName login_name获取 登录账号
func (obj *_SysUserOnlineMgr) WithLoginName(loginName string) Option {
	return optionFunc(func(o *options) { o.query["login_name"] = loginName })
}

// WithDeptName dept_name获取 部门名称
func (obj *_SysUserOnlineMgr) WithDeptName(deptName string) Option {
	return optionFunc(func(o *options) { o.query["dept_name"] = deptName })
}

// WithIPaddr ipaddr获取 登录IP地址
func (obj *_SysUserOnlineMgr) WithIPaddr(ipaddr string) Option {
	return optionFunc(func(o *options) { o.query["ipaddr"] = ipaddr })
}

// WithLoginLocation login_location获取 登录地点
func (obj *_SysUserOnlineMgr) WithLoginLocation(loginLocation string) Option {
	return optionFunc(func(o *options) { o.query["login_location"] = loginLocation })
}

// WithBrowser browser获取 浏览器类型
func (obj *_SysUserOnlineMgr) WithBrowser(browser string) Option {
	return optionFunc(func(o *options) { o.query["browser"] = browser })
}

// WithOs os获取 操作系统
func (obj *_SysUserOnlineMgr) WithOs(os string) Option {
	return optionFunc(func(o *options) { o.query["os"] = os })
}

// WithStatus status获取 在线状态on_line在线off_line离线
func (obj *_SysUserOnlineMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithStartTimestamp start_timestamp获取 session创建时间
func (obj *_SysUserOnlineMgr) WithStartTimestamp(startTimestamp time.Time) Option {
	return optionFunc(func(o *options) { o.query["start_timestamp"] = startTimestamp })
}

// WithLastAccessTime last_access_time获取 session最后访问时间
func (obj *_SysUserOnlineMgr) WithLastAccessTime(lastAccessTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["last_access_time"] = lastAccessTime })
}

// WithExpireTime expire_time获取 超时时间，单位为分钟
func (obj *_SysUserOnlineMgr) WithExpireTime(expireTime int) Option {
	return optionFunc(func(o *options) { o.query["expire_time"] = expireTime })
}

// GetByOption 功能选项模式获取
func (obj *_SysUserOnlineMgr) GetByOption(opts ...Option) (result SysUserOnline, err error) {
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
func (obj *_SysUserOnlineMgr) GetByOptions(opts ...Option) (results []*SysUserOnline, err error) {
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

// GetFromSessionID 通过sessionId获取内容 用户会话id
func (obj *_SysUserOnlineMgr) GetFromSessionID(sessionID string) (result SysUserOnline, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("sessionId = ?", sessionID).Find(&result).Error

	return
}

// GetBatchFromSessionID 批量唯一主键查找 用户会话id
func (obj *_SysUserOnlineMgr) GetBatchFromSessionID(sessionIDs []string) (results []*SysUserOnline, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("sessionId IN (?)", sessionIDs).Find(&results).Error

	return
}

// GetFromLoginName 通过login_name获取内容 登录账号
func (obj *_SysUserOnlineMgr) GetFromLoginName(loginName string) (results []*SysUserOnline, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("login_name = ?", loginName).Find(&results).Error

	return
}

// GetBatchFromLoginName 批量唯一主键查找 登录账号
func (obj *_SysUserOnlineMgr) GetBatchFromLoginName(loginNames []string) (results []*SysUserOnline, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("login_name IN (?)", loginNames).Find(&results).Error

	return
}

// GetFromDeptName 通过dept_name获取内容 部门名称
func (obj *_SysUserOnlineMgr) GetFromDeptName(deptName string) (results []*SysUserOnline, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("dept_name = ?", deptName).Find(&results).Error

	return
}

// GetBatchFromDeptName 批量唯一主键查找 部门名称
func (obj *_SysUserOnlineMgr) GetBatchFromDeptName(deptNames []string) (results []*SysUserOnline, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("dept_name IN (?)", deptNames).Find(&results).Error

	return
}

// GetFromIPaddr 通过ipaddr获取内容 登录IP地址
func (obj *_SysUserOnlineMgr) GetFromIPaddr(ipaddr string) (results []*SysUserOnline, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ipaddr = ?", ipaddr).Find(&results).Error

	return
}

// GetBatchFromIPaddr 批量唯一主键查找 登录IP地址
func (obj *_SysUserOnlineMgr) GetBatchFromIPaddr(ipaddrs []string) (results []*SysUserOnline, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ipaddr IN (?)", ipaddrs).Find(&results).Error

	return
}

// GetFromLoginLocation 通过login_location获取内容 登录地点
func (obj *_SysUserOnlineMgr) GetFromLoginLocation(loginLocation string) (results []*SysUserOnline, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("login_location = ?", loginLocation).Find(&results).Error

	return
}

// GetBatchFromLoginLocation 批量唯一主键查找 登录地点
func (obj *_SysUserOnlineMgr) GetBatchFromLoginLocation(loginLocations []string) (results []*SysUserOnline, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("login_location IN (?)", loginLocations).Find(&results).Error

	return
}

// GetFromBrowser 通过browser获取内容 浏览器类型
func (obj *_SysUserOnlineMgr) GetFromBrowser(browser string) (results []*SysUserOnline, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("browser = ?", browser).Find(&results).Error

	return
}

// GetBatchFromBrowser 批量唯一主键查找 浏览器类型
func (obj *_SysUserOnlineMgr) GetBatchFromBrowser(browsers []string) (results []*SysUserOnline, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("browser IN (?)", browsers).Find(&results).Error

	return
}

// GetFromOs 通过os获取内容 操作系统
func (obj *_SysUserOnlineMgr) GetFromOs(os string) (results []*SysUserOnline, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("os = ?", os).Find(&results).Error

	return
}

// GetBatchFromOs 批量唯一主键查找 操作系统
func (obj *_SysUserOnlineMgr) GetBatchFromOs(oss []string) (results []*SysUserOnline, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("os IN (?)", oss).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 在线状态on_line在线off_line离线
func (obj *_SysUserOnlineMgr) GetFromStatus(status string) (results []*SysUserOnline, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量唯一主键查找 在线状态on_line在线off_line离线
func (obj *_SysUserOnlineMgr) GetBatchFromStatus(statuss []string) (results []*SysUserOnline, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status IN (?)", statuss).Find(&results).Error

	return
}

// GetFromStartTimestamp 通过start_timestamp获取内容 session创建时间
func (obj *_SysUserOnlineMgr) GetFromStartTimestamp(startTimestamp time.Time) (results []*SysUserOnline, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("start_timestamp = ?", startTimestamp).Find(&results).Error

	return
}

// GetBatchFromStartTimestamp 批量唯一主键查找 session创建时间
func (obj *_SysUserOnlineMgr) GetBatchFromStartTimestamp(startTimestamps []time.Time) (results []*SysUserOnline, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("start_timestamp IN (?)", startTimestamps).Find(&results).Error

	return
}

// GetFromLastAccessTime 通过last_access_time获取内容 session最后访问时间
func (obj *_SysUserOnlineMgr) GetFromLastAccessTime(lastAccessTime time.Time) (results []*SysUserOnline, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("last_access_time = ?", lastAccessTime).Find(&results).Error

	return
}

// GetBatchFromLastAccessTime 批量唯一主键查找 session最后访问时间
func (obj *_SysUserOnlineMgr) GetBatchFromLastAccessTime(lastAccessTimes []time.Time) (results []*SysUserOnline, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("last_access_time IN (?)", lastAccessTimes).Find(&results).Error

	return
}

// GetFromExpireTime 通过expire_time获取内容 超时时间，单位为分钟
func (obj *_SysUserOnlineMgr) GetFromExpireTime(expireTime int) (results []*SysUserOnline, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("expire_time = ?", expireTime).Find(&results).Error

	return
}

// GetBatchFromExpireTime 批量唯一主键查找 超时时间，单位为分钟
func (obj *_SysUserOnlineMgr) GetBatchFromExpireTime(expireTimes []int) (results []*SysUserOnline, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("expire_time IN (?)", expireTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_SysUserOnlineMgr) FetchByPrimaryKey(sessionID string) (result SysUserOnline, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("sessionId = ?", sessionID).Find(&result).Error

	return
}
