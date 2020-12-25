package	models	
import (	
"gorm.io/gorm"	
"time"	
"fmt"	
"context"	
)	

type _SysUserCmsMgr struct {
	*_BaseMgr
}

// SysUserCmsMgr open func
func SysUserCmsMgr(db *gorm.DB) *_SysUserCmsMgr {
	if db == nil {
		panic(fmt.Errorf("SysUserCmsMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SysUserCmsMgr{_BaseMgr: &_BaseMgr{DB: db.Table("sys_user_cms"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SysUserCmsMgr) GetTableName() string {
	return "sys_user_cms"
}

// Get 获取
func (obj *_SysUserCmsMgr) Get() (result SysUserCms, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_SysUserCmsMgr) Gets() (results []*SysUserCms, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithPearUserID pear_user_id获取 标签系统用户ID
func (obj *_SysUserCmsMgr) WithPearUserID(pearUserID string) Option {
	return optionFunc(func(o *options) { o.query["pear_user_id"] = pearUserID })
}

// WithCmsUserID cms_user_id获取 cms系统用户ID
func (obj *_SysUserCmsMgr) WithCmsUserID(cmsUserID string) Option {
	return optionFunc(func(o *options) { o.query["cms_user_id"] = cmsUserID })
}

// WithCreateTime create_time获取 创建时间
func (obj *_SysUserCmsMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateTime update_time获取 修改时间
func (obj *_SysUserCmsMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithCreateUser create_user获取 创建人
func (obj *_SysUserCmsMgr) WithCreateUser(createUser string) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateUser update_user获取 修改人
func (obj *_SysUserCmsMgr) WithUpdateUser(updateUser string) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}


// GetByOption 功能选项模式获取
func (obj *_SysUserCmsMgr) GetByOption(opts ...Option) (result SysUserCms, err error) {
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
func (obj *_SysUserCmsMgr) GetByOptions(opts ...Option) (results []*SysUserCms, err error) {
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


// GetFromPearUserID 通过pear_user_id获取内容 标签系统用户ID 
func (obj *_SysUserCmsMgr) GetFromPearUserID(pearUserID string) (results []*SysUserCms, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("pear_user_id = ?", pearUserID).Find(&results).Error
	
	return
}

// GetBatchFromPearUserID 批量唯一主键查找 标签系统用户ID
func (obj *_SysUserCmsMgr) GetBatchFromPearUserID(pearUserIDs []string) (results []*SysUserCms, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("pear_user_id IN (?)", pearUserIDs).Find(&results).Error
	
	return
}
 
// GetFromCmsUserID 通过cms_user_id获取内容 cms系统用户ID 
func (obj *_SysUserCmsMgr) GetFromCmsUserID(cmsUserID string) (results []*SysUserCms, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cms_user_id = ?", cmsUserID).Find(&results).Error
	
	return
}

// GetBatchFromCmsUserID 批量唯一主键查找 cms系统用户ID
func (obj *_SysUserCmsMgr) GetBatchFromCmsUserID(cmsUserIDs []string) (results []*SysUserCms, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cms_user_id IN (?)", cmsUserIDs).Find(&results).Error
	
	return
}
 
// GetFromCreateTime 通过create_time获取内容 创建时间 
func (obj *_SysUserCmsMgr) GetFromCreateTime(createTime time.Time) (results []*SysUserCms, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_time = ?", createTime).Find(&results).Error
	
	return
}

// GetBatchFromCreateTime 批量唯一主键查找 创建时间
func (obj *_SysUserCmsMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*SysUserCms, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_time IN (?)", createTimes).Find(&results).Error
	
	return
}
 
// GetFromUpdateTime 通过update_time获取内容 修改时间 
func (obj *_SysUserCmsMgr) GetFromUpdateTime(updateTime time.Time) (results []*SysUserCms, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_time = ?", updateTime).Find(&results).Error
	
	return
}

// GetBatchFromUpdateTime 批量唯一主键查找 修改时间
func (obj *_SysUserCmsMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*SysUserCms, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_time IN (?)", updateTimes).Find(&results).Error
	
	return
}
 
// GetFromCreateUser 通过create_user获取内容 创建人 
func (obj *_SysUserCmsMgr) GetFromCreateUser(createUser string) (results []*SysUserCms, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_user = ?", createUser).Find(&results).Error
	
	return
}

// GetBatchFromCreateUser 批量唯一主键查找 创建人
func (obj *_SysUserCmsMgr) GetBatchFromCreateUser(createUsers []string) (results []*SysUserCms, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_user IN (?)", createUsers).Find(&results).Error
	
	return
}
 
// GetFromUpdateUser 通过update_user获取内容 修改人 
func (obj *_SysUserCmsMgr) GetFromUpdateUser(updateUser string) (results []*SysUserCms, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_user = ?", updateUser).Find(&results).Error
	
	return
}

// GetBatchFromUpdateUser 批量唯一主键查找 修改人
func (obj *_SysUserCmsMgr) GetBatchFromUpdateUser(updateUsers []string) (results []*SysUserCms, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_user IN (?)", updateUsers).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_SysUserCmsMgr) FetchByPrimaryKey(pearUserID string ,cmsUserID string ) (result SysUserCms, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("pear_user_id = ? AND cms_user_id = ?", pearUserID , cmsUserID).Find(&result).Error
	
	return
}
 

 

	

