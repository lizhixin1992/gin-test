package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type SysUserMgr struct {
	*_BaseMgr
}

// SysUserMgr open func
func NewSysUserMgr(db *gorm.DB) *SysUserMgr {
	if db == nil {
		panic(fmt.Errorf("SysUserMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &SysUserMgr{_BaseMgr: &_BaseMgr{DB: db.Table("sys_user"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *SysUserMgr) GetTableName() string {
	return "sys_user"
}

// Get 获取
func (obj *SysUserMgr) Get() (result SysUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *SysUserMgr) Gets() (results []*SysUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithUserID user_id获取 用户ID
func (obj *SysUserMgr) WithUserID(userID int64) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithDeptID dept_id获取 部门ID
func (obj *SysUserMgr) WithDeptID(deptID int64) Option {
	return optionFunc(func(o *options) { o.query["dept_id"] = deptID })
}

// WithLoginName login_name获取 登录账号
func (obj *SysUserMgr) WithLoginName(loginName string) Option {
	return optionFunc(func(o *options) { o.query["login_name"] = loginName })
}

// WithUserName user_name获取 用户昵称
func (obj *SysUserMgr) WithUserName(userName string) Option {
	return optionFunc(func(o *options) { o.query["user_name"] = userName })
}

// WithUserType user_type获取 用户类型（00系统用户）
func (obj *SysUserMgr) WithUserType(userType string) Option {
	return optionFunc(func(o *options) { o.query["user_type"] = userType })
}

// WithEmail email获取 用户邮箱
func (obj *SysUserMgr) WithEmail(email string) Option {
	return optionFunc(func(o *options) { o.query["email"] = email })
}

// WithPhonenumber phonenumber获取 手机号码
func (obj *SysUserMgr) WithPhonenumber(phonenumber string) Option {
	return optionFunc(func(o *options) { o.query["phonenumber"] = phonenumber })
}

// WithSex sex获取 用户性别（0男 1女 2未知）
func (obj *SysUserMgr) WithSex(sex string) Option {
	return optionFunc(func(o *options) { o.query["sex"] = sex })
}

// WithAvatar avatar获取 头像路径
func (obj *SysUserMgr) WithAvatar(avatar string) Option {
	return optionFunc(func(o *options) { o.query["avatar"] = avatar })
}

// WithPassword password获取 密码
func (obj *SysUserMgr) WithPassword(password string) Option {
	return optionFunc(func(o *options) { o.query["password"] = password })
}

// WithSalt salt获取 盐加密
func (obj *SysUserMgr) WithSalt(salt string) Option {
	return optionFunc(func(o *options) { o.query["salt"] = salt })
}

// WithStatus status获取 帐号状态（0正常 1停用）
func (obj *SysUserMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithDelFlag del_flag获取 删除标志（0代表存在 2代表删除）
func (obj *SysUserMgr) WithDelFlag(delFlag string) Option {
	return optionFunc(func(o *options) { o.query["del_flag"] = delFlag })
}

// WithLoginIP login_ip获取 最后登陆IP
func (obj *SysUserMgr) WithLoginIP(loginIP string) Option {
	return optionFunc(func(o *options) { o.query["login_ip"] = loginIP })
}

// WithLoginDate login_date获取 最后登陆时间
func (obj *SysUserMgr) WithLoginDate(loginDate time.Time) Option {
	return optionFunc(func(o *options) { o.query["login_date"] = loginDate })
}

// WithCreateBy create_by获取 创建者
func (obj *SysUserMgr) WithCreateBy(createBy string) Option {
	return optionFunc(func(o *options) { o.query["create_by"] = createBy })
}

// WithCreateTime create_time获取 创建时间
func (obj *SysUserMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateBy update_by获取 更新者
func (obj *SysUserMgr) WithUpdateBy(updateBy string) Option {
	return optionFunc(func(o *options) { o.query["update_by"] = updateBy })
}

// WithUpdateTime update_time获取 更新时间
func (obj *SysUserMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithRemark remark获取 备注
func (obj *SysUserMgr) WithRemark(remark string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithPortalUserID portal_user_id获取 portal用户表主键
func (obj *SysUserMgr) WithPortalUserID(portalUserID int64) Option {
	return optionFunc(func(o *options) { o.query["portal_user_id"] = portalUserID })
}

// GetByOption 功能选项模式获取
func (obj *SysUserMgr) GetByOption(opts ...Option) (result SysUser, err error) {
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
func (obj *SysUserMgr) GetByOptions(opts ...Option) (results []*SysUser, err error) {
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

// GetFromUserID 通过user_id获取内容 用户ID
func (obj *SysUserMgr) GetFromUserID(userID int64) (result SysUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id = ?", userID).Find(&result).Error

	return
}

// GetBatchFromUserID 批量唯一主键查找 用户ID
func (obj *SysUserMgr) GetBatchFromUserID(userIDs []int64) (results []*SysUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromDeptID 通过dept_id获取内容 部门ID
func (obj *SysUserMgr) GetFromDeptID(deptID int64) (results []*SysUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("dept_id = ?", deptID).Find(&results).Error

	return
}

// GetBatchFromDeptID 批量唯一主键查找 部门ID
func (obj *SysUserMgr) GetBatchFromDeptID(deptIDs []int64) (results []*SysUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("dept_id IN (?)", deptIDs).Find(&results).Error

	return
}

// GetFromLoginName 通过login_name获取内容 登录账号
func (obj *SysUserMgr) GetFromLoginName(loginName string) (results []*SysUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("login_name = ?", loginName).Find(&results).Error

	return
}

// GetBatchFromLoginName 批量唯一主键查找 登录账号
func (obj *SysUserMgr) GetBatchFromLoginName(loginNames []string) (results []*SysUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("login_name IN (?)", loginNames).Find(&results).Error

	return
}

// GetFromUserName 通过user_name获取内容 用户昵称
func (obj *SysUserMgr) GetFromUserName(userName string) (results []*SysUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_name = ?", userName).Find(&results).Error

	return
}

// GetBatchFromUserName 批量唯一主键查找 用户昵称
func (obj *SysUserMgr) GetBatchFromUserName(userNames []string) (results []*SysUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_name IN (?)", userNames).Find(&results).Error

	return
}

// GetFromUserType 通过user_type获取内容 用户类型（00系统用户）
func (obj *SysUserMgr) GetFromUserType(userType string) (results []*SysUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_type = ?", userType).Find(&results).Error

	return
}

// GetBatchFromUserType 批量唯一主键查找 用户类型（00系统用户）
func (obj *SysUserMgr) GetBatchFromUserType(userTypes []string) (results []*SysUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_type IN (?)", userTypes).Find(&results).Error

	return
}

// GetFromEmail 通过email获取内容 用户邮箱
func (obj *SysUserMgr) GetFromEmail(email string) (results []*SysUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("email = ?", email).Find(&results).Error

	return
}

// GetBatchFromEmail 批量唯一主键查找 用户邮箱
func (obj *SysUserMgr) GetBatchFromEmail(emails []string) (results []*SysUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("email IN (?)", emails).Find(&results).Error

	return
}

// GetFromPhonenumber 通过phonenumber获取内容 手机号码
func (obj *SysUserMgr) GetFromPhonenumber(phonenumber string) (results []*SysUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("phonenumber = ?", phonenumber).Find(&results).Error

	return
}

// GetBatchFromPhonenumber 批量唯一主键查找 手机号码
func (obj *SysUserMgr) GetBatchFromPhonenumber(phonenumbers []string) (results []*SysUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("phonenumber IN (?)", phonenumbers).Find(&results).Error

	return
}

// GetFromSex 通过sex获取内容 用户性别（0男 1女 2未知）
func (obj *SysUserMgr) GetFromSex(sex string) (results []*SysUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("sex = ?", sex).Find(&results).Error

	return
}

// GetBatchFromSex 批量唯一主键查找 用户性别（0男 1女 2未知）
func (obj *SysUserMgr) GetBatchFromSex(sexs []string) (results []*SysUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("sex IN (?)", sexs).Find(&results).Error

	return
}

// GetFromAvatar 通过avatar获取内容 头像路径
func (obj *SysUserMgr) GetFromAvatar(avatar string) (results []*SysUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("avatar = ?", avatar).Find(&results).Error

	return
}

// GetBatchFromAvatar 批量唯一主键查找 头像路径
func (obj *SysUserMgr) GetBatchFromAvatar(avatars []string) (results []*SysUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("avatar IN (?)", avatars).Find(&results).Error

	return
}

// GetFromPassword 通过password获取内容 密码
func (obj *SysUserMgr) GetFromPassword(password string) (results []*SysUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("password = ?", password).Find(&results).Error

	return
}

// GetBatchFromPassword 批量唯一主键查找 密码
func (obj *SysUserMgr) GetBatchFromPassword(passwords []string) (results []*SysUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("password IN (?)", passwords).Find(&results).Error

	return
}

// GetFromSalt 通过salt获取内容 盐加密
func (obj *SysUserMgr) GetFromSalt(salt string) (results []*SysUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("salt = ?", salt).Find(&results).Error

	return
}

// GetBatchFromSalt 批量唯一主键查找 盐加密
func (obj *SysUserMgr) GetBatchFromSalt(salts []string) (results []*SysUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("salt IN (?)", salts).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 帐号状态（0正常 1停用）
func (obj *SysUserMgr) GetFromStatus(status string) (results []*SysUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量唯一主键查找 帐号状态（0正常 1停用）
func (obj *SysUserMgr) GetBatchFromStatus(statuss []string) (results []*SysUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status IN (?)", statuss).Find(&results).Error

	return
}

// GetFromDelFlag 通过del_flag获取内容 删除标志（0代表存在 2代表删除）
func (obj *SysUserMgr) GetFromDelFlag(delFlag string) (results []*SysUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("del_flag = ?", delFlag).Find(&results).Error

	return
}

// GetBatchFromDelFlag 批量唯一主键查找 删除标志（0代表存在 2代表删除）
func (obj *SysUserMgr) GetBatchFromDelFlag(delFlags []string) (results []*SysUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("del_flag IN (?)", delFlags).Find(&results).Error

	return
}

// GetFromLoginIP 通过login_ip获取内容 最后登陆IP
func (obj *SysUserMgr) GetFromLoginIP(loginIP string) (results []*SysUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("login_ip = ?", loginIP).Find(&results).Error

	return
}

// GetBatchFromLoginIP 批量唯一主键查找 最后登陆IP
func (obj *SysUserMgr) GetBatchFromLoginIP(loginIPs []string) (results []*SysUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("login_ip IN (?)", loginIPs).Find(&results).Error

	return
}

// GetFromLoginDate 通过login_date获取内容 最后登陆时间
func (obj *SysUserMgr) GetFromLoginDate(loginDate time.Time) (results []*SysUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("login_date = ?", loginDate).Find(&results).Error

	return
}

// GetBatchFromLoginDate 批量唯一主键查找 最后登陆时间
func (obj *SysUserMgr) GetBatchFromLoginDate(loginDates []time.Time) (results []*SysUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("login_date IN (?)", loginDates).Find(&results).Error

	return
}

// GetFromCreateBy 通过create_by获取内容 创建者
func (obj *SysUserMgr) GetFromCreateBy(createBy string) (results []*SysUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_by = ?", createBy).Find(&results).Error

	return
}

// GetBatchFromCreateBy 批量唯一主键查找 创建者
func (obj *SysUserMgr) GetBatchFromCreateBy(createBys []string) (results []*SysUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_by IN (?)", createBys).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *SysUserMgr) GetFromCreateTime(createTime time.Time) (results []*SysUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_time = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量唯一主键查找 创建时间
func (obj *SysUserMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*SysUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_time IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromUpdateBy 通过update_by获取内容 更新者
func (obj *SysUserMgr) GetFromUpdateBy(updateBy string) (results []*SysUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_by = ?", updateBy).Find(&results).Error

	return
}

// GetBatchFromUpdateBy 批量唯一主键查找 更新者
func (obj *SysUserMgr) GetBatchFromUpdateBy(updateBys []string) (results []*SysUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_by IN (?)", updateBys).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *SysUserMgr) GetFromUpdateTime(updateTime time.Time) (results []*SysUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_time = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量唯一主键查找 更新时间
func (obj *SysUserMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*SysUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_time IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *SysUserMgr) GetFromRemark(remark string) (results []*SysUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("remark = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量唯一主键查找 备注
func (obj *SysUserMgr) GetBatchFromRemark(remarks []string) (results []*SysUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("remark IN (?)", remarks).Find(&results).Error

	return
}

// GetFromPortalUserID 通过portal_user_id获取内容 portal用户表主键
func (obj *SysUserMgr) GetFromPortalUserID(portalUserID int64) (results []*SysUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("portal_user_id = ?", portalUserID).Find(&results).Error

	return
}

// GetBatchFromPortalUserID 批量唯一主键查找 portal用户表主键
func (obj *SysUserMgr) GetBatchFromPortalUserID(portalUserIDs []int64) (results []*SysUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("portal_user_id IN (?)", portalUserIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *SysUserMgr) FetchByPrimaryKey(userID int64) (result SysUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id = ?", userID).Find(&result).Error

	return
}
