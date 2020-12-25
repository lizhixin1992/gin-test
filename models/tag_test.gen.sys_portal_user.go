package	models	
import (	
"gorm.io/gorm"	
"time"	
"fmt"	
"context"	
)	

type _SysPortalUserMgr struct {
	*_BaseMgr
}

// SysPortalUserMgr open func
func SysPortalUserMgr(db *gorm.DB) *_SysPortalUserMgr {
	if db == nil {
		panic(fmt.Errorf("SysPortalUserMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SysPortalUserMgr{_BaseMgr: &_BaseMgr{DB: db.Table("sys_portal_user"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SysPortalUserMgr) GetTableName() string {
	return "sys_portal_user"
}

// Get 获取
func (obj *_SysPortalUserMgr) Get() (result SysPortalUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_SysPortalUserMgr) Gets() (results []*SysPortalUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithUserID user_id获取 Id
func (obj *_SysPortalUserMgr) WithUserID(userID int64) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithOpenID open_id获取 用户在钉钉的唯一ID
func (obj *_SysPortalUserMgr) WithOpenID(openID string) Option {
	return optionFunc(func(o *options) { o.query["open_id"] = openID })
}

// WithLoginName login_name获取 登录账号
func (obj *_SysPortalUserMgr) WithLoginName(loginName string) Option {
	return optionFunc(func(o *options) { o.query["login_name"] = loginName })
}

// WithUserName user_name获取 用户昵称
func (obj *_SysPortalUserMgr) WithUserName(userName string) Option {
	return optionFunc(func(o *options) { o.query["user_name"] = userName })
}

// WithEmail email获取 邮箱
func (obj *_SysPortalUserMgr) WithEmail(email string) Option {
	return optionFunc(func(o *options) { o.query["email"] = email })
}

// WithPhoneNumber phone_number获取 手机号
func (obj *_SysPortalUserMgr) WithPhoneNumber(phoneNumber string) Option {
	return optionFunc(func(o *options) { o.query["phone_number"] = phoneNumber })
}

// WithAvatar avatar获取 头像路径
func (obj *_SysPortalUserMgr) WithAvatar(avatar string) Option {
	return optionFunc(func(o *options) { o.query["avatar"] = avatar })
}

// WithStatus status获取 状态：false正常 true停用
func (obj *_SysPortalUserMgr) WithStatus(status bool) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithDelFlag del_flag获取 删除标志：false 正常 true 删除
func (obj *_SysPortalUserMgr) WithDelFlag(delFlag bool) Option {
	return optionFunc(func(o *options) { o.query["del_flag"] = delFlag })
}

// WithIfAdmin if_admin获取 是否为下游系统管理员：false 否， true 是
func (obj *_SysPortalUserMgr) WithIfAdmin(ifAdmin bool) Option {
	return optionFunc(func(o *options) { o.query["if_admin"] = ifAdmin })
}

// WithDept dept获取 组织id与名称对象
func (obj *_SysPortalUserMgr) WithDept(dept string) Option {
	return optionFunc(func(o *options) { o.query["dept"] = dept })
}

// WithTenantry tenantry获取 租户id与名称对象
func (obj *_SysPortalUserMgr) WithTenantry(tenantry string) Option {
	return optionFunc(func(o *options) { o.query["tenantry"] = tenantry })
}

// WithCreateTime create_time获取 创建时间
func (obj *_SysPortalUserMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateTime update_time获取 修改时间
func (obj *_SysPortalUserMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithDefaultTenantryID default_tenantry_id获取 默认租户id
func (obj *_SysPortalUserMgr) WithDefaultTenantryID(defaultTenantryID string) Option {
	return optionFunc(func(o *options) { o.query["default_tenantry_id"] = defaultTenantryID })
}

// WithUnionID union_id获取 钉钉unionId
func (obj *_SysPortalUserMgr) WithUnionID(unionID string) Option {
	return optionFunc(func(o *options) { o.query["union_id"] = unionID })
}

// WithDingUserID ding_user_id获取 钉钉用户id
func (obj *_SysPortalUserMgr) WithDingUserID(dingUserID string) Option {
	return optionFunc(func(o *options) { o.query["ding_user_id"] = dingUserID })
}

// WithPenName pen_name获取 用户笔名
func (obj *_SysPortalUserMgr) WithPenName(penName string) Option {
	return optionFunc(func(o *options) { o.query["pen_name"] = penName })
}


// GetByOption 功能选项模式获取
func (obj *_SysPortalUserMgr) GetByOption(opts ...Option) (result SysPortalUser, err error) {
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
func (obj *_SysPortalUserMgr) GetByOptions(opts ...Option) (results []*SysPortalUser, err error) {
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


// GetFromUserID 通过user_id获取内容 Id 
func (obj *_SysPortalUserMgr)  GetFromUserID(userID int64) (result SysPortalUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id = ?", userID).Find(&result).Error
	
	return
}

// GetBatchFromUserID 批量唯一主键查找 Id
func (obj *_SysPortalUserMgr) GetBatchFromUserID(userIDs []int64) (results []*SysPortalUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id IN (?)", userIDs).Find(&results).Error
	
	return
}
 
// GetFromOpenID 通过open_id获取内容 用户在钉钉的唯一ID 
func (obj *_SysPortalUserMgr) GetFromOpenID(openID string) (results []*SysPortalUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("open_id = ?", openID).Find(&results).Error
	
	return
}

// GetBatchFromOpenID 批量唯一主键查找 用户在钉钉的唯一ID
func (obj *_SysPortalUserMgr) GetBatchFromOpenID(openIDs []string) (results []*SysPortalUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("open_id IN (?)", openIDs).Find(&results).Error
	
	return
}
 
// GetFromLoginName 通过login_name获取内容 登录账号 
func (obj *_SysPortalUserMgr)  GetFromLoginName(loginName string) (result SysPortalUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("login_name = ?", loginName).Find(&result).Error
	
	return
}

// GetBatchFromLoginName 批量唯一主键查找 登录账号
func (obj *_SysPortalUserMgr) GetBatchFromLoginName(loginNames []string) (results []*SysPortalUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("login_name IN (?)", loginNames).Find(&results).Error
	
	return
}
 
// GetFromUserName 通过user_name获取内容 用户昵称 
func (obj *_SysPortalUserMgr) GetFromUserName(userName string) (results []*SysPortalUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_name = ?", userName).Find(&results).Error
	
	return
}

// GetBatchFromUserName 批量唯一主键查找 用户昵称
func (obj *_SysPortalUserMgr) GetBatchFromUserName(userNames []string) (results []*SysPortalUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_name IN (?)", userNames).Find(&results).Error
	
	return
}
 
// GetFromEmail 通过email获取内容 邮箱 
func (obj *_SysPortalUserMgr) GetFromEmail(email string) (results []*SysPortalUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("email = ?", email).Find(&results).Error
	
	return
}

// GetBatchFromEmail 批量唯一主键查找 邮箱
func (obj *_SysPortalUserMgr) GetBatchFromEmail(emails []string) (results []*SysPortalUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("email IN (?)", emails).Find(&results).Error
	
	return
}
 
// GetFromPhoneNumber 通过phone_number获取内容 手机号 
func (obj *_SysPortalUserMgr) GetFromPhoneNumber(phoneNumber string) (results []*SysPortalUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("phone_number = ?", phoneNumber).Find(&results).Error
	
	return
}

// GetBatchFromPhoneNumber 批量唯一主键查找 手机号
func (obj *_SysPortalUserMgr) GetBatchFromPhoneNumber(phoneNumbers []string) (results []*SysPortalUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("phone_number IN (?)", phoneNumbers).Find(&results).Error
	
	return
}
 
// GetFromAvatar 通过avatar获取内容 头像路径 
func (obj *_SysPortalUserMgr) GetFromAvatar(avatar string) (results []*SysPortalUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("avatar = ?", avatar).Find(&results).Error
	
	return
}

// GetBatchFromAvatar 批量唯一主键查找 头像路径
func (obj *_SysPortalUserMgr) GetBatchFromAvatar(avatars []string) (results []*SysPortalUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("avatar IN (?)", avatars).Find(&results).Error
	
	return
}
 
// GetFromStatus 通过status获取内容 状态：false正常 true停用 
func (obj *_SysPortalUserMgr) GetFromStatus(status bool) (results []*SysPortalUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status = ?", status).Find(&results).Error
	
	return
}

// GetBatchFromStatus 批量唯一主键查找 状态：false正常 true停用
func (obj *_SysPortalUserMgr) GetBatchFromStatus(statuss []bool) (results []*SysPortalUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status IN (?)", statuss).Find(&results).Error
	
	return
}
 
// GetFromDelFlag 通过del_flag获取内容 删除标志：false 正常 true 删除 
func (obj *_SysPortalUserMgr) GetFromDelFlag(delFlag bool) (results []*SysPortalUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("del_flag = ?", delFlag).Find(&results).Error
	
	return
}

// GetBatchFromDelFlag 批量唯一主键查找 删除标志：false 正常 true 删除
func (obj *_SysPortalUserMgr) GetBatchFromDelFlag(delFlags []bool) (results []*SysPortalUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("del_flag IN (?)", delFlags).Find(&results).Error
	
	return
}
 
// GetFromIfAdmin 通过if_admin获取内容 是否为下游系统管理员：false 否， true 是 
func (obj *_SysPortalUserMgr) GetFromIfAdmin(ifAdmin bool) (results []*SysPortalUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("if_admin = ?", ifAdmin).Find(&results).Error
	
	return
}

// GetBatchFromIfAdmin 批量唯一主键查找 是否为下游系统管理员：false 否， true 是
func (obj *_SysPortalUserMgr) GetBatchFromIfAdmin(ifAdmins []bool) (results []*SysPortalUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("if_admin IN (?)", ifAdmins).Find(&results).Error
	
	return
}
 
// GetFromDept 通过dept获取内容 组织id与名称对象 
func (obj *_SysPortalUserMgr) GetFromDept(dept string) (results []*SysPortalUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("dept = ?", dept).Find(&results).Error
	
	return
}

// GetBatchFromDept 批量唯一主键查找 组织id与名称对象
func (obj *_SysPortalUserMgr) GetBatchFromDept(depts []string) (results []*SysPortalUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("dept IN (?)", depts).Find(&results).Error
	
	return
}
 
// GetFromTenantry 通过tenantry获取内容 租户id与名称对象 
func (obj *_SysPortalUserMgr) GetFromTenantry(tenantry string) (results []*SysPortalUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("tenantry = ?", tenantry).Find(&results).Error
	
	return
}

// GetBatchFromTenantry 批量唯一主键查找 租户id与名称对象
func (obj *_SysPortalUserMgr) GetBatchFromTenantry(tenantrys []string) (results []*SysPortalUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("tenantry IN (?)", tenantrys).Find(&results).Error
	
	return
}
 
// GetFromCreateTime 通过create_time获取内容 创建时间 
func (obj *_SysPortalUserMgr) GetFromCreateTime(createTime time.Time) (results []*SysPortalUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_time = ?", createTime).Find(&results).Error
	
	return
}

// GetBatchFromCreateTime 批量唯一主键查找 创建时间
func (obj *_SysPortalUserMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*SysPortalUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_time IN (?)", createTimes).Find(&results).Error
	
	return
}
 
// GetFromUpdateTime 通过update_time获取内容 修改时间 
func (obj *_SysPortalUserMgr) GetFromUpdateTime(updateTime time.Time) (results []*SysPortalUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_time = ?", updateTime).Find(&results).Error
	
	return
}

// GetBatchFromUpdateTime 批量唯一主键查找 修改时间
func (obj *_SysPortalUserMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*SysPortalUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_time IN (?)", updateTimes).Find(&results).Error
	
	return
}
 
// GetFromDefaultTenantryID 通过default_tenantry_id获取内容 默认租户id 
func (obj *_SysPortalUserMgr) GetFromDefaultTenantryID(defaultTenantryID string) (results []*SysPortalUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("default_tenantry_id = ?", defaultTenantryID).Find(&results).Error
	
	return
}

// GetBatchFromDefaultTenantryID 批量唯一主键查找 默认租户id
func (obj *_SysPortalUserMgr) GetBatchFromDefaultTenantryID(defaultTenantryIDs []string) (results []*SysPortalUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("default_tenantry_id IN (?)", defaultTenantryIDs).Find(&results).Error
	
	return
}
 
// GetFromUnionID 通过union_id获取内容 钉钉unionId 
func (obj *_SysPortalUserMgr) GetFromUnionID(unionID string) (results []*SysPortalUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("union_id = ?", unionID).Find(&results).Error
	
	return
}

// GetBatchFromUnionID 批量唯一主键查找 钉钉unionId
func (obj *_SysPortalUserMgr) GetBatchFromUnionID(unionIDs []string) (results []*SysPortalUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("union_id IN (?)", unionIDs).Find(&results).Error
	
	return
}
 
// GetFromDingUserID 通过ding_user_id获取内容 钉钉用户id 
func (obj *_SysPortalUserMgr) GetFromDingUserID(dingUserID string) (results []*SysPortalUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ding_user_id = ?", dingUserID).Find(&results).Error
	
	return
}

// GetBatchFromDingUserID 批量唯一主键查找 钉钉用户id
func (obj *_SysPortalUserMgr) GetBatchFromDingUserID(dingUserIDs []string) (results []*SysPortalUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ding_user_id IN (?)", dingUserIDs).Find(&results).Error
	
	return
}
 
// GetFromPenName 通过pen_name获取内容 用户笔名 
func (obj *_SysPortalUserMgr) GetFromPenName(penName string) (results []*SysPortalUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("pen_name = ?", penName).Find(&results).Error
	
	return
}

// GetBatchFromPenName 批量唯一主键查找 用户笔名
func (obj *_SysPortalUserMgr) GetBatchFromPenName(penNames []string) (results []*SysPortalUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("pen_name IN (?)", penNames).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_SysPortalUserMgr) FetchByPrimaryKey(userID int64 ) (result SysPortalUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id = ?", userID).Find(&result).Error
	
	return
}
 
 // FetchUniqueByPortalLoginName primay or index 获取唯一内容
 func (obj *_SysPortalUserMgr) FetchUniqueByPortalLoginName(loginName string ) (result SysPortalUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("login_name = ?", loginName).Find(&result).Error
	
	return
}
 

 

	

