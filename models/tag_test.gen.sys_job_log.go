package	models	
import (	
"gorm.io/gorm"	
"time"	
"fmt"	
"context"	
)	

type _SysJobLogMgr struct {
	*_BaseMgr
}

// SysJobLogMgr open func
func SysJobLogMgr(db *gorm.DB) *_SysJobLogMgr {
	if db == nil {
		panic(fmt.Errorf("SysJobLogMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SysJobLogMgr{_BaseMgr: &_BaseMgr{DB: db.Table("sys_job_log"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SysJobLogMgr) GetTableName() string {
	return "sys_job_log"
}

// Get 获取
func (obj *_SysJobLogMgr) Get() (result SysJobLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_SysJobLogMgr) Gets() (results []*SysJobLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithJobLogID job_log_id获取 任务日志ID
func (obj *_SysJobLogMgr) WithJobLogID(jobLogID int64) Option {
	return optionFunc(func(o *options) { o.query["job_log_id"] = jobLogID })
}

// WithJobName job_name获取 任务名称
func (obj *_SysJobLogMgr) WithJobName(jobName string) Option {
	return optionFunc(func(o *options) { o.query["job_name"] = jobName })
}

// WithJobGroup job_group获取 任务组名
func (obj *_SysJobLogMgr) WithJobGroup(jobGroup string) Option {
	return optionFunc(func(o *options) { o.query["job_group"] = jobGroup })
}

// WithInvokeTarget invoke_target获取 调用目标字符串
func (obj *_SysJobLogMgr) WithInvokeTarget(invokeTarget string) Option {
	return optionFunc(func(o *options) { o.query["invoke_target"] = invokeTarget })
}

// WithJobMessage job_message获取 日志信息
func (obj *_SysJobLogMgr) WithJobMessage(jobMessage string) Option {
	return optionFunc(func(o *options) { o.query["job_message"] = jobMessage })
}

// WithStatus status获取 执行状态（0正常 1失败）
func (obj *_SysJobLogMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithExceptionInfo exception_info获取 异常信息
func (obj *_SysJobLogMgr) WithExceptionInfo(exceptionInfo string) Option {
	return optionFunc(func(o *options) { o.query["exception_info"] = exceptionInfo })
}

// WithCreateTime create_time获取 创建时间
func (obj *_SysJobLogMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}


// GetByOption 功能选项模式获取
func (obj *_SysJobLogMgr) GetByOption(opts ...Option) (result SysJobLog, err error) {
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
func (obj *_SysJobLogMgr) GetByOptions(opts ...Option) (results []*SysJobLog, err error) {
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


// GetFromJobLogID 通过job_log_id获取内容 任务日志ID 
func (obj *_SysJobLogMgr)  GetFromJobLogID(jobLogID int64) (result SysJobLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("job_log_id = ?", jobLogID).Find(&result).Error
	
	return
}

// GetBatchFromJobLogID 批量唯一主键查找 任务日志ID
func (obj *_SysJobLogMgr) GetBatchFromJobLogID(jobLogIDs []int64) (results []*SysJobLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("job_log_id IN (?)", jobLogIDs).Find(&results).Error
	
	return
}
 
// GetFromJobName 通过job_name获取内容 任务名称 
func (obj *_SysJobLogMgr) GetFromJobName(jobName string) (results []*SysJobLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("job_name = ?", jobName).Find(&results).Error
	
	return
}

// GetBatchFromJobName 批量唯一主键查找 任务名称
func (obj *_SysJobLogMgr) GetBatchFromJobName(jobNames []string) (results []*SysJobLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("job_name IN (?)", jobNames).Find(&results).Error
	
	return
}
 
// GetFromJobGroup 通过job_group获取内容 任务组名 
func (obj *_SysJobLogMgr) GetFromJobGroup(jobGroup string) (results []*SysJobLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("job_group = ?", jobGroup).Find(&results).Error
	
	return
}

// GetBatchFromJobGroup 批量唯一主键查找 任务组名
func (obj *_SysJobLogMgr) GetBatchFromJobGroup(jobGroups []string) (results []*SysJobLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("job_group IN (?)", jobGroups).Find(&results).Error
	
	return
}
 
// GetFromInvokeTarget 通过invoke_target获取内容 调用目标字符串 
func (obj *_SysJobLogMgr) GetFromInvokeTarget(invokeTarget string) (results []*SysJobLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("invoke_target = ?", invokeTarget).Find(&results).Error
	
	return
}

// GetBatchFromInvokeTarget 批量唯一主键查找 调用目标字符串
func (obj *_SysJobLogMgr) GetBatchFromInvokeTarget(invokeTargets []string) (results []*SysJobLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("invoke_target IN (?)", invokeTargets).Find(&results).Error
	
	return
}
 
// GetFromJobMessage 通过job_message获取内容 日志信息 
func (obj *_SysJobLogMgr) GetFromJobMessage(jobMessage string) (results []*SysJobLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("job_message = ?", jobMessage).Find(&results).Error
	
	return
}

// GetBatchFromJobMessage 批量唯一主键查找 日志信息
func (obj *_SysJobLogMgr) GetBatchFromJobMessage(jobMessages []string) (results []*SysJobLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("job_message IN (?)", jobMessages).Find(&results).Error
	
	return
}
 
// GetFromStatus 通过status获取内容 执行状态（0正常 1失败） 
func (obj *_SysJobLogMgr) GetFromStatus(status string) (results []*SysJobLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status = ?", status).Find(&results).Error
	
	return
}

// GetBatchFromStatus 批量唯一主键查找 执行状态（0正常 1失败）
func (obj *_SysJobLogMgr) GetBatchFromStatus(statuss []string) (results []*SysJobLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status IN (?)", statuss).Find(&results).Error
	
	return
}
 
// GetFromExceptionInfo 通过exception_info获取内容 异常信息 
func (obj *_SysJobLogMgr) GetFromExceptionInfo(exceptionInfo string) (results []*SysJobLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("exception_info = ?", exceptionInfo).Find(&results).Error
	
	return
}

// GetBatchFromExceptionInfo 批量唯一主键查找 异常信息
func (obj *_SysJobLogMgr) GetBatchFromExceptionInfo(exceptionInfos []string) (results []*SysJobLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("exception_info IN (?)", exceptionInfos).Find(&results).Error
	
	return
}
 
// GetFromCreateTime 通过create_time获取内容 创建时间 
func (obj *_SysJobLogMgr) GetFromCreateTime(createTime time.Time) (results []*SysJobLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_time = ?", createTime).Find(&results).Error
	
	return
}

// GetBatchFromCreateTime 批量唯一主键查找 创建时间
func (obj *_SysJobLogMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*SysJobLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_time IN (?)", createTimes).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_SysJobLogMgr) FetchByPrimaryKey(jobLogID int64 ) (result SysJobLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("job_log_id = ?", jobLogID).Find(&result).Error
	
	return
}
 

 

	

