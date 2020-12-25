package	models	
import (	
"gorm.io/gorm"	
"time"	
"fmt"	
"context"	
)	

type _SysNoticeMgr struct {
	*_BaseMgr
}

// SysNoticeMgr open func
func SysNoticeMgr(db *gorm.DB) *_SysNoticeMgr {
	if db == nil {
		panic(fmt.Errorf("SysNoticeMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SysNoticeMgr{_BaseMgr: &_BaseMgr{DB: db.Table("sys_notice"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SysNoticeMgr) GetTableName() string {
	return "sys_notice"
}

// Get 获取
func (obj *_SysNoticeMgr) Get() (result SysNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_SysNoticeMgr) Gets() (results []*SysNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithNoticeID notice_id获取 公告ID
func (obj *_SysNoticeMgr) WithNoticeID(noticeID int) Option {
	return optionFunc(func(o *options) { o.query["notice_id"] = noticeID })
}

// WithNoticeTitle notice_title获取 公告标题
func (obj *_SysNoticeMgr) WithNoticeTitle(noticeTitle string) Option {
	return optionFunc(func(o *options) { o.query["notice_title"] = noticeTitle })
}

// WithNoticeType notice_type获取 公告类型（1通知 2公告）
func (obj *_SysNoticeMgr) WithNoticeType(noticeType string) Option {
	return optionFunc(func(o *options) { o.query["notice_type"] = noticeType })
}

// WithNoticeContent notice_content获取 公告内容
func (obj *_SysNoticeMgr) WithNoticeContent(noticeContent string) Option {
	return optionFunc(func(o *options) { o.query["notice_content"] = noticeContent })
}

// WithStatus status获取 公告状态（0正常 1关闭）
func (obj *_SysNoticeMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithCreateBy create_by获取 创建者
func (obj *_SysNoticeMgr) WithCreateBy(createBy string) Option {
	return optionFunc(func(o *options) { o.query["create_by"] = createBy })
}

// WithCreateTime create_time获取 创建时间
func (obj *_SysNoticeMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateBy update_by获取 更新者
func (obj *_SysNoticeMgr) WithUpdateBy(updateBy string) Option {
	return optionFunc(func(o *options) { o.query["update_by"] = updateBy })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_SysNoticeMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithRemark remark获取 备注
func (obj *_SysNoticeMgr) WithRemark(remark string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}


// GetByOption 功能选项模式获取
func (obj *_SysNoticeMgr) GetByOption(opts ...Option) (result SysNotice, err error) {
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
func (obj *_SysNoticeMgr) GetByOptions(opts ...Option) (results []*SysNotice, err error) {
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


// GetFromNoticeID 通过notice_id获取内容 公告ID 
func (obj *_SysNoticeMgr)  GetFromNoticeID(noticeID int) (result SysNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("notice_id = ?", noticeID).Find(&result).Error
	
	return
}

// GetBatchFromNoticeID 批量唯一主键查找 公告ID
func (obj *_SysNoticeMgr) GetBatchFromNoticeID(noticeIDs []int) (results []*SysNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("notice_id IN (?)", noticeIDs).Find(&results).Error
	
	return
}
 
// GetFromNoticeTitle 通过notice_title获取内容 公告标题 
func (obj *_SysNoticeMgr) GetFromNoticeTitle(noticeTitle string) (results []*SysNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("notice_title = ?", noticeTitle).Find(&results).Error
	
	return
}

// GetBatchFromNoticeTitle 批量唯一主键查找 公告标题
func (obj *_SysNoticeMgr) GetBatchFromNoticeTitle(noticeTitles []string) (results []*SysNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("notice_title IN (?)", noticeTitles).Find(&results).Error
	
	return
}
 
// GetFromNoticeType 通过notice_type获取内容 公告类型（1通知 2公告） 
func (obj *_SysNoticeMgr) GetFromNoticeType(noticeType string) (results []*SysNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("notice_type = ?", noticeType).Find(&results).Error
	
	return
}

// GetBatchFromNoticeType 批量唯一主键查找 公告类型（1通知 2公告）
func (obj *_SysNoticeMgr) GetBatchFromNoticeType(noticeTypes []string) (results []*SysNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("notice_type IN (?)", noticeTypes).Find(&results).Error
	
	return
}
 
// GetFromNoticeContent 通过notice_content获取内容 公告内容 
func (obj *_SysNoticeMgr) GetFromNoticeContent(noticeContent string) (results []*SysNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("notice_content = ?", noticeContent).Find(&results).Error
	
	return
}

// GetBatchFromNoticeContent 批量唯一主键查找 公告内容
func (obj *_SysNoticeMgr) GetBatchFromNoticeContent(noticeContents []string) (results []*SysNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("notice_content IN (?)", noticeContents).Find(&results).Error
	
	return
}
 
// GetFromStatus 通过status获取内容 公告状态（0正常 1关闭） 
func (obj *_SysNoticeMgr) GetFromStatus(status string) (results []*SysNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status = ?", status).Find(&results).Error
	
	return
}

// GetBatchFromStatus 批量唯一主键查找 公告状态（0正常 1关闭）
func (obj *_SysNoticeMgr) GetBatchFromStatus(statuss []string) (results []*SysNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status IN (?)", statuss).Find(&results).Error
	
	return
}
 
// GetFromCreateBy 通过create_by获取内容 创建者 
func (obj *_SysNoticeMgr) GetFromCreateBy(createBy string) (results []*SysNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_by = ?", createBy).Find(&results).Error
	
	return
}

// GetBatchFromCreateBy 批量唯一主键查找 创建者
func (obj *_SysNoticeMgr) GetBatchFromCreateBy(createBys []string) (results []*SysNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_by IN (?)", createBys).Find(&results).Error
	
	return
}
 
// GetFromCreateTime 通过create_time获取内容 创建时间 
func (obj *_SysNoticeMgr) GetFromCreateTime(createTime time.Time) (results []*SysNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_time = ?", createTime).Find(&results).Error
	
	return
}

// GetBatchFromCreateTime 批量唯一主键查找 创建时间
func (obj *_SysNoticeMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*SysNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_time IN (?)", createTimes).Find(&results).Error
	
	return
}
 
// GetFromUpdateBy 通过update_by获取内容 更新者 
func (obj *_SysNoticeMgr) GetFromUpdateBy(updateBy string) (results []*SysNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_by = ?", updateBy).Find(&results).Error
	
	return
}

// GetBatchFromUpdateBy 批量唯一主键查找 更新者
func (obj *_SysNoticeMgr) GetBatchFromUpdateBy(updateBys []string) (results []*SysNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_by IN (?)", updateBys).Find(&results).Error
	
	return
}
 
// GetFromUpdateTime 通过update_time获取内容 更新时间 
func (obj *_SysNoticeMgr) GetFromUpdateTime(updateTime time.Time) (results []*SysNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_time = ?", updateTime).Find(&results).Error
	
	return
}

// GetBatchFromUpdateTime 批量唯一主键查找 更新时间
func (obj *_SysNoticeMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*SysNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_time IN (?)", updateTimes).Find(&results).Error
	
	return
}
 
// GetFromRemark 通过remark获取内容 备注 
func (obj *_SysNoticeMgr) GetFromRemark(remark string) (results []*SysNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("remark = ?", remark).Find(&results).Error
	
	return
}

// GetBatchFromRemark 批量唯一主键查找 备注
func (obj *_SysNoticeMgr) GetBatchFromRemark(remarks []string) (results []*SysNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("remark IN (?)", remarks).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_SysNoticeMgr) FetchByPrimaryKey(noticeID int ) (result SysNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("notice_id = ?", noticeID).Find(&result).Error
	
	return
}
 

 

	

