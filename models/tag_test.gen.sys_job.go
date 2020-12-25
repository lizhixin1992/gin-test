package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _SysJobMgr struct {
	*_BaseMgr
}

// SysJobMgr open func
func SysJobMgr(db *gorm.DB) *_SysJobMgr {
	if db == nil {
		panic(fmt.Errorf("SysJobMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SysJobMgr{_BaseMgr: &_BaseMgr{DB: db.Table("sys_job"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SysJobMgr) GetTableName() string {
	return "sys_job"
}

// Get 获取
func (obj *_SysJobMgr) Get() (result SysJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SysJobMgr) Gets() (results []*SysJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithJobID job_id获取 任务ID
func (obj *_SysJobMgr) WithJobID(jobID int64) Option {
	return optionFunc(func(o *options) { o.query["job_id"] = jobID })
}

// WithJobName job_name获取 任务名称
func (obj *_SysJobMgr) WithJobName(jobName string) Option {
	return optionFunc(func(o *options) { o.query["job_name"] = jobName })
}

// WithJobGroup job_group获取 任务组名
func (obj *_SysJobMgr) WithJobGroup(jobGroup string) Option {
	return optionFunc(func(o *options) { o.query["job_group"] = jobGroup })
}

// WithInvokeTarget invoke_target获取 调用目标字符串
func (obj *_SysJobMgr) WithInvokeTarget(invokeTarget string) Option {
	return optionFunc(func(o *options) { o.query["invoke_target"] = invokeTarget })
}

// WithCronExpression cron_expression获取 cron执行表达式
func (obj *_SysJobMgr) WithCronExpression(cronExpression string) Option {
	return optionFunc(func(o *options) { o.query["cron_expression"] = cronExpression })
}

// WithMisfirePolicy misfire_policy获取 计划执行错误策略（1立即执行 2执行一次 3放弃执行）
func (obj *_SysJobMgr) WithMisfirePolicy(misfirePolicy string) Option {
	return optionFunc(func(o *options) { o.query["misfire_policy"] = misfirePolicy })
}

// WithConcurrent concurrent获取 是否并发执行（0允许 1禁止）
func (obj *_SysJobMgr) WithConcurrent(concurrent string) Option {
	return optionFunc(func(o *options) { o.query["concurrent"] = concurrent })
}

// WithStatus status获取 状态（0正常 1暂停）
func (obj *_SysJobMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithCreateBy create_by获取 创建者
func (obj *_SysJobMgr) WithCreateBy(createBy string) Option {
	return optionFunc(func(o *options) { o.query["create_by"] = createBy })
}

// WithCreateTime create_time获取 创建时间
func (obj *_SysJobMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateBy update_by获取 更新者
func (obj *_SysJobMgr) WithUpdateBy(updateBy string) Option {
	return optionFunc(func(o *options) { o.query["update_by"] = updateBy })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_SysJobMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithRemark remark获取 备注信息
func (obj *_SysJobMgr) WithRemark(remark string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// GetByOption 功能选项模式获取
func (obj *_SysJobMgr) GetByOption(opts ...Option) (result SysJob, err error) {
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
func (obj *_SysJobMgr) GetByOptions(opts ...Option) (results []*SysJob, err error) {
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

// GetFromJobID 通过job_id获取内容 任务ID
func (obj *_SysJobMgr) GetFromJobID(jobID int64) (results []*SysJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("job_id = ?", jobID).Find(&results).Error

	return
}

// GetBatchFromJobID 批量唯一主键查找 任务ID
func (obj *_SysJobMgr) GetBatchFromJobID(jobIDs []int64) (results []*SysJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("job_id IN (?)", jobIDs).Find(&results).Error

	return
}

// GetFromJobName 通过job_name获取内容 任务名称
func (obj *_SysJobMgr) GetFromJobName(jobName string) (results []*SysJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("job_name = ?", jobName).Find(&results).Error

	return
}

// GetBatchFromJobName 批量唯一主键查找 任务名称
func (obj *_SysJobMgr) GetBatchFromJobName(jobNames []string) (results []*SysJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("job_name IN (?)", jobNames).Find(&results).Error

	return
}

// GetFromJobGroup 通过job_group获取内容 任务组名
func (obj *_SysJobMgr) GetFromJobGroup(jobGroup string) (results []*SysJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("job_group = ?", jobGroup).Find(&results).Error

	return
}

// GetBatchFromJobGroup 批量唯一主键查找 任务组名
func (obj *_SysJobMgr) GetBatchFromJobGroup(jobGroups []string) (results []*SysJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("job_group IN (?)", jobGroups).Find(&results).Error

	return
}

// GetFromInvokeTarget 通过invoke_target获取内容 调用目标字符串
func (obj *_SysJobMgr) GetFromInvokeTarget(invokeTarget string) (results []*SysJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("invoke_target = ?", invokeTarget).Find(&results).Error

	return
}

// GetBatchFromInvokeTarget 批量唯一主键查找 调用目标字符串
func (obj *_SysJobMgr) GetBatchFromInvokeTarget(invokeTargets []string) (results []*SysJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("invoke_target IN (?)", invokeTargets).Find(&results).Error

	return
}

// GetFromCronExpression 通过cron_expression获取内容 cron执行表达式
func (obj *_SysJobMgr) GetFromCronExpression(cronExpression string) (results []*SysJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cron_expression = ?", cronExpression).Find(&results).Error

	return
}

// GetBatchFromCronExpression 批量唯一主键查找 cron执行表达式
func (obj *_SysJobMgr) GetBatchFromCronExpression(cronExpressions []string) (results []*SysJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cron_expression IN (?)", cronExpressions).Find(&results).Error

	return
}

// GetFromMisfirePolicy 通过misfire_policy获取内容 计划执行错误策略（1立即执行 2执行一次 3放弃执行）
func (obj *_SysJobMgr) GetFromMisfirePolicy(misfirePolicy string) (results []*SysJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("misfire_policy = ?", misfirePolicy).Find(&results).Error

	return
}

// GetBatchFromMisfirePolicy 批量唯一主键查找 计划执行错误策略（1立即执行 2执行一次 3放弃执行）
func (obj *_SysJobMgr) GetBatchFromMisfirePolicy(misfirePolicys []string) (results []*SysJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("misfire_policy IN (?)", misfirePolicys).Find(&results).Error

	return
}

// GetFromConcurrent 通过concurrent获取内容 是否并发执行（0允许 1禁止）
func (obj *_SysJobMgr) GetFromConcurrent(concurrent string) (results []*SysJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("concurrent = ?", concurrent).Find(&results).Error

	return
}

// GetBatchFromConcurrent 批量唯一主键查找 是否并发执行（0允许 1禁止）
func (obj *_SysJobMgr) GetBatchFromConcurrent(concurrents []string) (results []*SysJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("concurrent IN (?)", concurrents).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 状态（0正常 1暂停）
func (obj *_SysJobMgr) GetFromStatus(status string) (results []*SysJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量唯一主键查找 状态（0正常 1暂停）
func (obj *_SysJobMgr) GetBatchFromStatus(statuss []string) (results []*SysJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status IN (?)", statuss).Find(&results).Error

	return
}

// GetFromCreateBy 通过create_by获取内容 创建者
func (obj *_SysJobMgr) GetFromCreateBy(createBy string) (results []*SysJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_by = ?", createBy).Find(&results).Error

	return
}

// GetBatchFromCreateBy 批量唯一主键查找 创建者
func (obj *_SysJobMgr) GetBatchFromCreateBy(createBys []string) (results []*SysJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_by IN (?)", createBys).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_SysJobMgr) GetFromCreateTime(createTime time.Time) (results []*SysJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_time = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量唯一主键查找 创建时间
func (obj *_SysJobMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*SysJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_time IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromUpdateBy 通过update_by获取内容 更新者
func (obj *_SysJobMgr) GetFromUpdateBy(updateBy string) (results []*SysJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_by = ?", updateBy).Find(&results).Error

	return
}

// GetBatchFromUpdateBy 批量唯一主键查找 更新者
func (obj *_SysJobMgr) GetBatchFromUpdateBy(updateBys []string) (results []*SysJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_by IN (?)", updateBys).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_SysJobMgr) GetFromUpdateTime(updateTime time.Time) (results []*SysJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_time = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量唯一主键查找 更新时间
func (obj *_SysJobMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*SysJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_time IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注信息
func (obj *_SysJobMgr) GetFromRemark(remark string) (results []*SysJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("remark = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量唯一主键查找 备注信息
func (obj *_SysJobMgr) GetBatchFromRemark(remarks []string) (results []*SysJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("remark IN (?)", remarks).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_SysJobMgr) FetchByPrimaryKey(jobID int64, jobName string, jobGroup string) (result SysJob, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("job_id = ? AND job_name = ? AND job_group = ?", jobID, jobName, jobGroup).Find(&result).Error

	return
}
