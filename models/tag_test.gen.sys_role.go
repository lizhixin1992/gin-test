package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _SysRoleMgr struct {
	*_BaseMgr
}

// SysRoleMgr open func
func SysRoleMgr(db *gorm.DB) *_SysRoleMgr {
	if db == nil {
		panic(fmt.Errorf("SysRoleMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SysRoleMgr{_BaseMgr: &_BaseMgr{DB: db.Table("sys_role"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SysRoleMgr) GetTableName() string {
	return "sys_role"
}

// Get 获取
func (obj *_SysRoleMgr) Get() (result SysRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SysRoleMgr) Gets() (results []*SysRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithRoleID role_id获取 角色ID
func (obj *_SysRoleMgr) WithRoleID(roleID int64) Option {
	return optionFunc(func(o *options) { o.query["role_id"] = roleID })
}

// WithRoleName role_name获取 角色名称
func (obj *_SysRoleMgr) WithRoleName(roleName string) Option {
	return optionFunc(func(o *options) { o.query["role_name"] = roleName })
}

// WithRoleKey role_key获取 角色权限字符串
func (obj *_SysRoleMgr) WithRoleKey(roleKey string) Option {
	return optionFunc(func(o *options) { o.query["role_key"] = roleKey })
}

// WithRoleSort role_sort获取 显示顺序
func (obj *_SysRoleMgr) WithRoleSort(roleSort int) Option {
	return optionFunc(func(o *options) { o.query["role_sort"] = roleSort })
}

// WithDataScope data_scope获取 数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）
func (obj *_SysRoleMgr) WithDataScope(dataScope string) Option {
	return optionFunc(func(o *options) { o.query["data_scope"] = dataScope })
}

// WithStatus status获取 角色状态（0正常 1停用）
func (obj *_SysRoleMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithDelFlag del_flag获取 删除标志（0代表存在 2代表删除）
func (obj *_SysRoleMgr) WithDelFlag(delFlag string) Option {
	return optionFunc(func(o *options) { o.query["del_flag"] = delFlag })
}

// WithCreateBy create_by获取 创建者
func (obj *_SysRoleMgr) WithCreateBy(createBy string) Option {
	return optionFunc(func(o *options) { o.query["create_by"] = createBy })
}

// WithCreateTime create_time获取 创建时间
func (obj *_SysRoleMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateBy update_by获取 更新者
func (obj *_SysRoleMgr) WithUpdateBy(updateBy string) Option {
	return optionFunc(func(o *options) { o.query["update_by"] = updateBy })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_SysRoleMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithRemark remark获取 备注
func (obj *_SysRoleMgr) WithRemark(remark string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// GetByOption 功能选项模式获取
func (obj *_SysRoleMgr) GetByOption(opts ...Option) (result SysRole, err error) {
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
func (obj *_SysRoleMgr) GetByOptions(opts ...Option) (results []*SysRole, err error) {
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

// GetFromRoleID 通过role_id获取内容 角色ID
func (obj *_SysRoleMgr) GetFromRoleID(roleID int64) (result SysRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("role_id = ?", roleID).Find(&result).Error

	return
}

// GetBatchFromRoleID 批量唯一主键查找 角色ID
func (obj *_SysRoleMgr) GetBatchFromRoleID(roleIDs []int64) (results []*SysRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("role_id IN (?)", roleIDs).Find(&results).Error

	return
}

// GetFromRoleName 通过role_name获取内容 角色名称
func (obj *_SysRoleMgr) GetFromRoleName(roleName string) (results []*SysRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("role_name = ?", roleName).Find(&results).Error

	return
}

// GetBatchFromRoleName 批量唯一主键查找 角色名称
func (obj *_SysRoleMgr) GetBatchFromRoleName(roleNames []string) (results []*SysRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("role_name IN (?)", roleNames).Find(&results).Error

	return
}

// GetFromRoleKey 通过role_key获取内容 角色权限字符串
func (obj *_SysRoleMgr) GetFromRoleKey(roleKey string) (results []*SysRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("role_key = ?", roleKey).Find(&results).Error

	return
}

// GetBatchFromRoleKey 批量唯一主键查找 角色权限字符串
func (obj *_SysRoleMgr) GetBatchFromRoleKey(roleKeys []string) (results []*SysRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("role_key IN (?)", roleKeys).Find(&results).Error

	return
}

// GetFromRoleSort 通过role_sort获取内容 显示顺序
func (obj *_SysRoleMgr) GetFromRoleSort(roleSort int) (results []*SysRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("role_sort = ?", roleSort).Find(&results).Error

	return
}

// GetBatchFromRoleSort 批量唯一主键查找 显示顺序
func (obj *_SysRoleMgr) GetBatchFromRoleSort(roleSorts []int) (results []*SysRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("role_sort IN (?)", roleSorts).Find(&results).Error

	return
}

// GetFromDataScope 通过data_scope获取内容 数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）
func (obj *_SysRoleMgr) GetFromDataScope(dataScope string) (results []*SysRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("data_scope = ?", dataScope).Find(&results).Error

	return
}

// GetBatchFromDataScope 批量唯一主键查找 数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）
func (obj *_SysRoleMgr) GetBatchFromDataScope(dataScopes []string) (results []*SysRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("data_scope IN (?)", dataScopes).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 角色状态（0正常 1停用）
func (obj *_SysRoleMgr) GetFromStatus(status string) (results []*SysRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量唯一主键查找 角色状态（0正常 1停用）
func (obj *_SysRoleMgr) GetBatchFromStatus(statuss []string) (results []*SysRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status IN (?)", statuss).Find(&results).Error

	return
}

// GetFromDelFlag 通过del_flag获取内容 删除标志（0代表存在 2代表删除）
func (obj *_SysRoleMgr) GetFromDelFlag(delFlag string) (results []*SysRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("del_flag = ?", delFlag).Find(&results).Error

	return
}

// GetBatchFromDelFlag 批量唯一主键查找 删除标志（0代表存在 2代表删除）
func (obj *_SysRoleMgr) GetBatchFromDelFlag(delFlags []string) (results []*SysRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("del_flag IN (?)", delFlags).Find(&results).Error

	return
}

// GetFromCreateBy 通过create_by获取内容 创建者
func (obj *_SysRoleMgr) GetFromCreateBy(createBy string) (results []*SysRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_by = ?", createBy).Find(&results).Error

	return
}

// GetBatchFromCreateBy 批量唯一主键查找 创建者
func (obj *_SysRoleMgr) GetBatchFromCreateBy(createBys []string) (results []*SysRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_by IN (?)", createBys).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_SysRoleMgr) GetFromCreateTime(createTime time.Time) (results []*SysRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_time = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量唯一主键查找 创建时间
func (obj *_SysRoleMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*SysRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_time IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromUpdateBy 通过update_by获取内容 更新者
func (obj *_SysRoleMgr) GetFromUpdateBy(updateBy string) (results []*SysRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_by = ?", updateBy).Find(&results).Error

	return
}

// GetBatchFromUpdateBy 批量唯一主键查找 更新者
func (obj *_SysRoleMgr) GetBatchFromUpdateBy(updateBys []string) (results []*SysRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_by IN (?)", updateBys).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_SysRoleMgr) GetFromUpdateTime(updateTime time.Time) (results []*SysRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_time = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量唯一主键查找 更新时间
func (obj *_SysRoleMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*SysRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_time IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_SysRoleMgr) GetFromRemark(remark string) (results []*SysRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("remark = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量唯一主键查找 备注
func (obj *_SysRoleMgr) GetBatchFromRemark(remarks []string) (results []*SysRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("remark IN (?)", remarks).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_SysRoleMgr) FetchByPrimaryKey(roleID int64) (result SysRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("role_id = ?", roleID).Find(&result).Error

	return
}
