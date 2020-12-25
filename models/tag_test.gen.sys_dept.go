package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _SysDeptMgr struct {
	*_BaseMgr
}

// SysDeptMgr open func
func SysDeptMgr(db *gorm.DB) *_SysDeptMgr {
	if db == nil {
		panic(fmt.Errorf("SysDeptMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SysDeptMgr{_BaseMgr: &_BaseMgr{DB: db.Table("sys_dept"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SysDeptMgr) GetTableName() string {
	return "sys_dept"
}

// Get 获取
func (obj *_SysDeptMgr) Get() (result SysDept, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SysDeptMgr) Gets() (results []*SysDept, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithDeptID dept_id获取 部门id
func (obj *_SysDeptMgr) WithDeptID(deptID int64) Option {
	return optionFunc(func(o *options) { o.query["dept_id"] = deptID })
}

// WithParentID parent_id获取 父部门id
func (obj *_SysDeptMgr) WithParentID(parentID int64) Option {
	return optionFunc(func(o *options) { o.query["parent_id"] = parentID })
}

// WithAncestors ancestors获取 祖级列表
func (obj *_SysDeptMgr) WithAncestors(ancestors string) Option {
	return optionFunc(func(o *options) { o.query["ancestors"] = ancestors })
}

// WithDeptName dept_name获取 部门名称
func (obj *_SysDeptMgr) WithDeptName(deptName string) Option {
	return optionFunc(func(o *options) { o.query["dept_name"] = deptName })
}

// WithOrderNum order_num获取 显示顺序
func (obj *_SysDeptMgr) WithOrderNum(orderNum int) Option {
	return optionFunc(func(o *options) { o.query["order_num"] = orderNum })
}

// WithLeader leader获取 负责人
func (obj *_SysDeptMgr) WithLeader(leader string) Option {
	return optionFunc(func(o *options) { o.query["leader"] = leader })
}

// WithPhone phone获取 联系电话
func (obj *_SysDeptMgr) WithPhone(phone string) Option {
	return optionFunc(func(o *options) { o.query["phone"] = phone })
}

// WithEmail email获取 邮箱
func (obj *_SysDeptMgr) WithEmail(email string) Option {
	return optionFunc(func(o *options) { o.query["email"] = email })
}

// WithStatus status获取 部门状态（0正常 1停用）
func (obj *_SysDeptMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithDelFlag del_flag获取 删除标志（0代表存在 2代表删除）
func (obj *_SysDeptMgr) WithDelFlag(delFlag string) Option {
	return optionFunc(func(o *options) { o.query["del_flag"] = delFlag })
}

// WithCreateBy create_by获取 创建者
func (obj *_SysDeptMgr) WithCreateBy(createBy string) Option {
	return optionFunc(func(o *options) { o.query["create_by"] = createBy })
}

// WithCreateTime create_time获取 创建时间
func (obj *_SysDeptMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateBy update_by获取 更新者
func (obj *_SysDeptMgr) WithUpdateBy(updateBy string) Option {
	return optionFunc(func(o *options) { o.query["update_by"] = updateBy })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_SysDeptMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// GetByOption 功能选项模式获取
func (obj *_SysDeptMgr) GetByOption(opts ...Option) (result SysDept, err error) {
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
func (obj *_SysDeptMgr) GetByOptions(opts ...Option) (results []*SysDept, err error) {
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

// GetFromDeptID 通过dept_id获取内容 部门id
func (obj *_SysDeptMgr) GetFromDeptID(deptID int64) (result SysDept, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("dept_id = ?", deptID).Find(&result).Error

	return
}

// GetBatchFromDeptID 批量唯一主键查找 部门id
func (obj *_SysDeptMgr) GetBatchFromDeptID(deptIDs []int64) (results []*SysDept, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("dept_id IN (?)", deptIDs).Find(&results).Error

	return
}

// GetFromParentID 通过parent_id获取内容 父部门id
func (obj *_SysDeptMgr) GetFromParentID(parentID int64) (results []*SysDept, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("parent_id = ?", parentID).Find(&results).Error

	return
}

// GetBatchFromParentID 批量唯一主键查找 父部门id
func (obj *_SysDeptMgr) GetBatchFromParentID(parentIDs []int64) (results []*SysDept, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("parent_id IN (?)", parentIDs).Find(&results).Error

	return
}

// GetFromAncestors 通过ancestors获取内容 祖级列表
func (obj *_SysDeptMgr) GetFromAncestors(ancestors string) (results []*SysDept, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ancestors = ?", ancestors).Find(&results).Error

	return
}

// GetBatchFromAncestors 批量唯一主键查找 祖级列表
func (obj *_SysDeptMgr) GetBatchFromAncestors(ancestorss []string) (results []*SysDept, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ancestors IN (?)", ancestorss).Find(&results).Error

	return
}

// GetFromDeptName 通过dept_name获取内容 部门名称
func (obj *_SysDeptMgr) GetFromDeptName(deptName string) (results []*SysDept, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("dept_name = ?", deptName).Find(&results).Error

	return
}

// GetBatchFromDeptName 批量唯一主键查找 部门名称
func (obj *_SysDeptMgr) GetBatchFromDeptName(deptNames []string) (results []*SysDept, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("dept_name IN (?)", deptNames).Find(&results).Error

	return
}

// GetFromOrderNum 通过order_num获取内容 显示顺序
func (obj *_SysDeptMgr) GetFromOrderNum(orderNum int) (results []*SysDept, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("order_num = ?", orderNum).Find(&results).Error

	return
}

// GetBatchFromOrderNum 批量唯一主键查找 显示顺序
func (obj *_SysDeptMgr) GetBatchFromOrderNum(orderNums []int) (results []*SysDept, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("order_num IN (?)", orderNums).Find(&results).Error

	return
}

// GetFromLeader 通过leader获取内容 负责人
func (obj *_SysDeptMgr) GetFromLeader(leader string) (results []*SysDept, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("leader = ?", leader).Find(&results).Error

	return
}

// GetBatchFromLeader 批量唯一主键查找 负责人
func (obj *_SysDeptMgr) GetBatchFromLeader(leaders []string) (results []*SysDept, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("leader IN (?)", leaders).Find(&results).Error

	return
}

// GetFromPhone 通过phone获取内容 联系电话
func (obj *_SysDeptMgr) GetFromPhone(phone string) (results []*SysDept, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("phone = ?", phone).Find(&results).Error

	return
}

// GetBatchFromPhone 批量唯一主键查找 联系电话
func (obj *_SysDeptMgr) GetBatchFromPhone(phones []string) (results []*SysDept, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("phone IN (?)", phones).Find(&results).Error

	return
}

// GetFromEmail 通过email获取内容 邮箱
func (obj *_SysDeptMgr) GetFromEmail(email string) (results []*SysDept, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("email = ?", email).Find(&results).Error

	return
}

// GetBatchFromEmail 批量唯一主键查找 邮箱
func (obj *_SysDeptMgr) GetBatchFromEmail(emails []string) (results []*SysDept, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("email IN (?)", emails).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 部门状态（0正常 1停用）
func (obj *_SysDeptMgr) GetFromStatus(status string) (results []*SysDept, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量唯一主键查找 部门状态（0正常 1停用）
func (obj *_SysDeptMgr) GetBatchFromStatus(statuss []string) (results []*SysDept, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status IN (?)", statuss).Find(&results).Error

	return
}

// GetFromDelFlag 通过del_flag获取内容 删除标志（0代表存在 2代表删除）
func (obj *_SysDeptMgr) GetFromDelFlag(delFlag string) (results []*SysDept, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("del_flag = ?", delFlag).Find(&results).Error

	return
}

// GetBatchFromDelFlag 批量唯一主键查找 删除标志（0代表存在 2代表删除）
func (obj *_SysDeptMgr) GetBatchFromDelFlag(delFlags []string) (results []*SysDept, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("del_flag IN (?)", delFlags).Find(&results).Error

	return
}

// GetFromCreateBy 通过create_by获取内容 创建者
func (obj *_SysDeptMgr) GetFromCreateBy(createBy string) (results []*SysDept, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_by = ?", createBy).Find(&results).Error

	return
}

// GetBatchFromCreateBy 批量唯一主键查找 创建者
func (obj *_SysDeptMgr) GetBatchFromCreateBy(createBys []string) (results []*SysDept, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_by IN (?)", createBys).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_SysDeptMgr) GetFromCreateTime(createTime time.Time) (results []*SysDept, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_time = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量唯一主键查找 创建时间
func (obj *_SysDeptMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*SysDept, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_time IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromUpdateBy 通过update_by获取内容 更新者
func (obj *_SysDeptMgr) GetFromUpdateBy(updateBy string) (results []*SysDept, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_by = ?", updateBy).Find(&results).Error

	return
}

// GetBatchFromUpdateBy 批量唯一主键查找 更新者
func (obj *_SysDeptMgr) GetBatchFromUpdateBy(updateBys []string) (results []*SysDept, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_by IN (?)", updateBys).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_SysDeptMgr) GetFromUpdateTime(updateTime time.Time) (results []*SysDept, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_time = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量唯一主键查找 更新时间
func (obj *_SysDeptMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*SysDept, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_time IN (?)", updateTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_SysDeptMgr) FetchByPrimaryKey(deptID int64) (result SysDept, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("dept_id = ?", deptID).Find(&result).Error

	return
}
