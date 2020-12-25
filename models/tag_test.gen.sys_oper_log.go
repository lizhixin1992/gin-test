package	models	
import (	
"time"	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _SysOperLogMgr struct {
	*_BaseMgr
}

// SysOperLogMgr open func
func SysOperLogMgr(db *gorm.DB) *_SysOperLogMgr {
	if db == nil {
		panic(fmt.Errorf("SysOperLogMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SysOperLogMgr{_BaseMgr: &_BaseMgr{DB: db.Table("sys_oper_log"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SysOperLogMgr) GetTableName() string {
	return "sys_oper_log"
}

// Get 获取
func (obj *_SysOperLogMgr) Get() (result SysOperLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_SysOperLogMgr) Gets() (results []*SysOperLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithOperID oper_id获取 日志主键
func (obj *_SysOperLogMgr) WithOperID(operID int64) Option {
	return optionFunc(func(o *options) { o.query["oper_id"] = operID })
}

// WithTitle title获取 模块标题
func (obj *_SysOperLogMgr) WithTitle(title string) Option {
	return optionFunc(func(o *options) { o.query["title"] = title })
}

// WithBusinessType business_type获取 业务类型（0其它 1新增 2修改 3删除）
func (obj *_SysOperLogMgr) WithBusinessType(businessType int) Option {
	return optionFunc(func(o *options) { o.query["business_type"] = businessType })
}

// WithMethod method获取 方法名称
func (obj *_SysOperLogMgr) WithMethod(method string) Option {
	return optionFunc(func(o *options) { o.query["method"] = method })
}

// WithRequestMethod request_method获取 请求方式
func (obj *_SysOperLogMgr) WithRequestMethod(requestMethod string) Option {
	return optionFunc(func(o *options) { o.query["request_method"] = requestMethod })
}

// WithOperatorType operator_type获取 操作类别（0其它 1后台用户 2手机端用户）
func (obj *_SysOperLogMgr) WithOperatorType(operatorType int) Option {
	return optionFunc(func(o *options) { o.query["operator_type"] = operatorType })
}

// WithOperName oper_name获取 操作人员
func (obj *_SysOperLogMgr) WithOperName(operName string) Option {
	return optionFunc(func(o *options) { o.query["oper_name"] = operName })
}

// WithDeptName dept_name获取 部门名称
func (obj *_SysOperLogMgr) WithDeptName(deptName string) Option {
	return optionFunc(func(o *options) { o.query["dept_name"] = deptName })
}

// WithOperURL oper_url获取 请求URL
func (obj *_SysOperLogMgr) WithOperURL(operURL string) Option {
	return optionFunc(func(o *options) { o.query["oper_url"] = operURL })
}

// WithOperIP oper_ip获取 主机地址
func (obj *_SysOperLogMgr) WithOperIP(operIP string) Option {
	return optionFunc(func(o *options) { o.query["oper_ip"] = operIP })
}

// WithOperLocation oper_location获取 操作地点
func (obj *_SysOperLogMgr) WithOperLocation(operLocation string) Option {
	return optionFunc(func(o *options) { o.query["oper_location"] = operLocation })
}

// WithOperParam oper_param获取 请求参数
func (obj *_SysOperLogMgr) WithOperParam(operParam string) Option {
	return optionFunc(func(o *options) { o.query["oper_param"] = operParam })
}

// WithStatus status获取 操作状态（0正常 1异常）
func (obj *_SysOperLogMgr) WithStatus(status int) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithErrorMsg error_msg获取 错误消息
func (obj *_SysOperLogMgr) WithErrorMsg(errorMsg string) Option {
	return optionFunc(func(o *options) { o.query["error_msg"] = errorMsg })
}

// WithOperTime oper_time获取 操作时间
func (obj *_SysOperLogMgr) WithOperTime(operTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["oper_time"] = operTime })
}


// GetByOption 功能选项模式获取
func (obj *_SysOperLogMgr) GetByOption(opts ...Option) (result SysOperLog, err error) {
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
func (obj *_SysOperLogMgr) GetByOptions(opts ...Option) (results []*SysOperLog, err error) {
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


// GetFromOperID 通过oper_id获取内容 日志主键 
func (obj *_SysOperLogMgr)  GetFromOperID(operID int64) (result SysOperLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("oper_id = ?", operID).Find(&result).Error
	
	return
}

// GetBatchFromOperID 批量唯一主键查找 日志主键
func (obj *_SysOperLogMgr) GetBatchFromOperID(operIDs []int64) (results []*SysOperLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("oper_id IN (?)", operIDs).Find(&results).Error
	
	return
}
 
// GetFromTitle 通过title获取内容 模块标题 
func (obj *_SysOperLogMgr) GetFromTitle(title string) (results []*SysOperLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("title = ?", title).Find(&results).Error
	
	return
}

// GetBatchFromTitle 批量唯一主键查找 模块标题
func (obj *_SysOperLogMgr) GetBatchFromTitle(titles []string) (results []*SysOperLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("title IN (?)", titles).Find(&results).Error
	
	return
}
 
// GetFromBusinessType 通过business_type获取内容 业务类型（0其它 1新增 2修改 3删除） 
func (obj *_SysOperLogMgr) GetFromBusinessType(businessType int) (results []*SysOperLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("business_type = ?", businessType).Find(&results).Error
	
	return
}

// GetBatchFromBusinessType 批量唯一主键查找 业务类型（0其它 1新增 2修改 3删除）
func (obj *_SysOperLogMgr) GetBatchFromBusinessType(businessTypes []int) (results []*SysOperLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("business_type IN (?)", businessTypes).Find(&results).Error
	
	return
}
 
// GetFromMethod 通过method获取内容 方法名称 
func (obj *_SysOperLogMgr) GetFromMethod(method string) (results []*SysOperLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("method = ?", method).Find(&results).Error
	
	return
}

// GetBatchFromMethod 批量唯一主键查找 方法名称
func (obj *_SysOperLogMgr) GetBatchFromMethod(methods []string) (results []*SysOperLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("method IN (?)", methods).Find(&results).Error
	
	return
}
 
// GetFromRequestMethod 通过request_method获取内容 请求方式 
func (obj *_SysOperLogMgr) GetFromRequestMethod(requestMethod string) (results []*SysOperLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("request_method = ?", requestMethod).Find(&results).Error
	
	return
}

// GetBatchFromRequestMethod 批量唯一主键查找 请求方式
func (obj *_SysOperLogMgr) GetBatchFromRequestMethod(requestMethods []string) (results []*SysOperLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("request_method IN (?)", requestMethods).Find(&results).Error
	
	return
}
 
// GetFromOperatorType 通过operator_type获取内容 操作类别（0其它 1后台用户 2手机端用户） 
func (obj *_SysOperLogMgr) GetFromOperatorType(operatorType int) (results []*SysOperLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("operator_type = ?", operatorType).Find(&results).Error
	
	return
}

// GetBatchFromOperatorType 批量唯一主键查找 操作类别（0其它 1后台用户 2手机端用户）
func (obj *_SysOperLogMgr) GetBatchFromOperatorType(operatorTypes []int) (results []*SysOperLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("operator_type IN (?)", operatorTypes).Find(&results).Error
	
	return
}
 
// GetFromOperName 通过oper_name获取内容 操作人员 
func (obj *_SysOperLogMgr) GetFromOperName(operName string) (results []*SysOperLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("oper_name = ?", operName).Find(&results).Error
	
	return
}

// GetBatchFromOperName 批量唯一主键查找 操作人员
func (obj *_SysOperLogMgr) GetBatchFromOperName(operNames []string) (results []*SysOperLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("oper_name IN (?)", operNames).Find(&results).Error
	
	return
}
 
// GetFromDeptName 通过dept_name获取内容 部门名称 
func (obj *_SysOperLogMgr) GetFromDeptName(deptName string) (results []*SysOperLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("dept_name = ?", deptName).Find(&results).Error
	
	return
}

// GetBatchFromDeptName 批量唯一主键查找 部门名称
func (obj *_SysOperLogMgr) GetBatchFromDeptName(deptNames []string) (results []*SysOperLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("dept_name IN (?)", deptNames).Find(&results).Error
	
	return
}
 
// GetFromOperURL 通过oper_url获取内容 请求URL 
func (obj *_SysOperLogMgr) GetFromOperURL(operURL string) (results []*SysOperLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("oper_url = ?", operURL).Find(&results).Error
	
	return
}

// GetBatchFromOperURL 批量唯一主键查找 请求URL
func (obj *_SysOperLogMgr) GetBatchFromOperURL(operURLs []string) (results []*SysOperLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("oper_url IN (?)", operURLs).Find(&results).Error
	
	return
}
 
// GetFromOperIP 通过oper_ip获取内容 主机地址 
func (obj *_SysOperLogMgr) GetFromOperIP(operIP string) (results []*SysOperLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("oper_ip = ?", operIP).Find(&results).Error
	
	return
}

// GetBatchFromOperIP 批量唯一主键查找 主机地址
func (obj *_SysOperLogMgr) GetBatchFromOperIP(operIPs []string) (results []*SysOperLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("oper_ip IN (?)", operIPs).Find(&results).Error
	
	return
}
 
// GetFromOperLocation 通过oper_location获取内容 操作地点 
func (obj *_SysOperLogMgr) GetFromOperLocation(operLocation string) (results []*SysOperLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("oper_location = ?", operLocation).Find(&results).Error
	
	return
}

// GetBatchFromOperLocation 批量唯一主键查找 操作地点
func (obj *_SysOperLogMgr) GetBatchFromOperLocation(operLocations []string) (results []*SysOperLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("oper_location IN (?)", operLocations).Find(&results).Error
	
	return
}
 
// GetFromOperParam 通过oper_param获取内容 请求参数 
func (obj *_SysOperLogMgr) GetFromOperParam(operParam string) (results []*SysOperLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("oper_param = ?", operParam).Find(&results).Error
	
	return
}

// GetBatchFromOperParam 批量唯一主键查找 请求参数
func (obj *_SysOperLogMgr) GetBatchFromOperParam(operParams []string) (results []*SysOperLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("oper_param IN (?)", operParams).Find(&results).Error
	
	return
}
 
// GetFromStatus 通过status获取内容 操作状态（0正常 1异常） 
func (obj *_SysOperLogMgr) GetFromStatus(status int) (results []*SysOperLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status = ?", status).Find(&results).Error
	
	return
}

// GetBatchFromStatus 批量唯一主键查找 操作状态（0正常 1异常）
func (obj *_SysOperLogMgr) GetBatchFromStatus(statuss []int) (results []*SysOperLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status IN (?)", statuss).Find(&results).Error
	
	return
}
 
// GetFromErrorMsg 通过error_msg获取内容 错误消息 
func (obj *_SysOperLogMgr) GetFromErrorMsg(errorMsg string) (results []*SysOperLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("error_msg = ?", errorMsg).Find(&results).Error
	
	return
}

// GetBatchFromErrorMsg 批量唯一主键查找 错误消息
func (obj *_SysOperLogMgr) GetBatchFromErrorMsg(errorMsgs []string) (results []*SysOperLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("error_msg IN (?)", errorMsgs).Find(&results).Error
	
	return
}
 
// GetFromOperTime 通过oper_time获取内容 操作时间 
func (obj *_SysOperLogMgr) GetFromOperTime(operTime time.Time) (results []*SysOperLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("oper_time = ?", operTime).Find(&results).Error
	
	return
}

// GetBatchFromOperTime 批量唯一主键查找 操作时间
func (obj *_SysOperLogMgr) GetBatchFromOperTime(operTimes []time.Time) (results []*SysOperLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("oper_time IN (?)", operTimes).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_SysOperLogMgr) FetchByPrimaryKey(operID int64 ) (result SysOperLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("oper_id = ?", operID).Find(&result).Error
	
	return
}
 

 

	

