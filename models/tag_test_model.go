package models

import (
	"time"
)

// LabelCharacter 人物表
type LabelCharacter struct {
	CharacterID              int       `gorm:"primaryKey;column:character_id;type:int(11);not null" json:"character_id,omitempty"`                // 人物主键
	CharacterUUID            string    `gorm:"primaryKey;column:character_uuid;type:varchar(20);not null" json:"character_uuid,omitempty"`        // 唯一标识
	CharacterName            string    `gorm:"column:character_name;type:varchar(50);not null" json:"character_name,omitempty"`                   // 人物名称
	CharacterNamePinyinFirst string    `gorm:"column:character_name_pinyin_first;type:varchar(255)" json:"character_name_pinyin_first,omitempty"` // 人物名称首字母
	CharacterType            int       `gorm:"column:character_type;type:int(11)" json:"character_type,omitempty"`                                // 人物类型
	CharacterImg             string    `gorm:"column:character_img;type:varchar(255)" json:"character_img,omitempty"`                             // 图片
	CharacterImgAliOssPath   string    `gorm:"column:character_img_ali_oss_path;type:varchar(255)" json:"character_img_ali_oss_path,omitempty"`   // 图片在阿里云OSS上地址
	CharacterSex             int       `gorm:"column:character_sex;type:int(1)" json:"character_sex,omitempty"`                                   // 性别,0:男,1:女
	CharacterBirthday        time.Time `gorm:"column:character_birthday;type:datetime" json:"character_birthday,omitempty"`                       // 出生日期
	CharacterCountry         string    `gorm:"column:character_country;type:varchar(20)" json:"character_country,omitempty"`                      // 国家
	ApplyStatus              int       `gorm:"column:apply_status;type:int(1)" json:"apply_status,omitempty"`                                     // 是否启用,0:启用,1:不启用
	DeleteStatus             int       `gorm:"column:delete_status;type:int(1)" json:"delete_status,omitempty"`                                   // 是否删除,0:未删除,1:已删除
	CreateTime               time.Time `gorm:"column:create_time;type:datetime" json:"create_time,omitempty"`                                     // 创建时间
	UpdateTime               time.Time `gorm:"column:update_time;type:datetime" json:"update_time,omitempty"`                                     // 修改时间
	CreateUser               string    `gorm:"column:create_user;type:varchar(50)" json:"create_user,omitempty"`                                  // 创建人
	UpdateUser               string    `gorm:"column:update_user;type:varchar(50)" json:"update_user,omitempty"`                                  // 修改人
	UploadUUID               string    `gorm:"column:upload_uuid;type:varchar(50)" json:"upload_uuid,omitempty"`
}

// LabelTag 标签表
type LabelTag struct {
	TagID            int       `gorm:"primaryKey;index:index_tag_id;column:tag_id;type:int(11);not null" json:"tag_id,omitempty"`           // 标签主键
	TagUUID          string    `gorm:"primaryKey;index:index_tag_uuid;column:tag_uuid;type:varchar(20);not null" json:"tag_uuid,omitempty"` // 唯一标识
	TagName          string    `gorm:"column:tag_name;type:varchar(50)" json:"tag_name,omitempty"`                                          // 标签名称
	TagType          int       `gorm:"column:tag_type;type:int(11)" json:"tag_type,omitempty"`                                              // 标签类型
	TagParentID      string    `gorm:"index:index_tag_parent_id;column:tag_parent_id;type:varchar(20)" json:"tag_parent_id,omitempty"`      // 父级标识
	TagLayer         int       `gorm:"column:tag_layer;type:int(11)" json:"tag_layer,omitempty"`                                            // 层级
	TagImg           string    `gorm:"column:tag_img;type:varchar(255)" json:"tag_img,omitempty"`                                           // 图片
	TagImgAliOssPath string    `gorm:"column:tag_img_ali_oss_path;type:varchar(255)" json:"tag_img_ali_oss_path,omitempty"`                 // 图片在阿里云OSS上地址
	ApplyStatus      int       `gorm:"column:apply_status;type:int(1)" json:"apply_status,omitempty"`                                       // 是否启用,0:启用,1:不启用
	DeleteStatus     int       `gorm:"column:delete_status;type:int(1)" json:"delete_status,omitempty"`                                     // 是否删除,0:未删除,1:已删除
	CreateTime       time.Time `gorm:"column:create_time;type:datetime" json:"create_time,omitempty"`                                       // 创建时间
	UpdateTime       time.Time `gorm:"column:update_time;type:datetime" json:"update_time,omitempty"`                                       // 修改时间
	CreateUser       string    `gorm:"column:create_user;type:varchar(50)" json:"create_user,omitempty"`                                    // 创建人
	UpdateUser       string    `gorm:"column:update_user;type:varchar(50)" json:"update_user,omitempty"`                                    // 修改人
	OrderNum         int       `gorm:"column:order_num;type:int(11)" json:"order_num,omitempty"`                                            // 排序序号
	UploadUUID       string    `gorm:"column:upload_uuid;type:varchar(50)" json:"upload_uuid,omitempty"`
	IsHaveChildren   int       `gorm:"column:is_have_children;type:int(1)" json:"is_have_children,omitempty"` // 是否可以有下级,0:可以,1:不可以
}

// LabelTagCharacter 标签和人物关联表
type LabelTagCharacter struct {
	RelationUUID  string    `gorm:"column:relation_uuid;type:varchar(20)" json:"relation_uuid,omitempty"`                       // 唯一标识
	TagUUID       string    `gorm:"primaryKey;column:tag_uuid;type:varchar(20);not null" json:"tag_uuid,omitempty"`             // 标签标识
	CharacterUUID string    `gorm:"primaryKey;column:character_uuid;type:varchar(20);not null" json:"character_uuid,omitempty"` // 人物标识
	CreateTime    time.Time `gorm:"column:create_time;type:datetime" json:"create_time,omitempty"`                              // 创建时间
	UpdateTime    time.Time `gorm:"column:update_time;type:datetime" json:"update_time,omitempty"`                              // 修改时间
	CreateUser    string    `gorm:"column:create_user;type:varchar(50)" json:"create_user,omitempty"`                           // 创建人
	UpdateUser    string    `gorm:"column:update_user;type:varchar(50)" json:"update_user,omitempty"`                           // 修改人
	DeleteStatus  int       `gorm:"column:delete_status;type:int(1)" json:"delete_status,omitempty"`                            // 是否删除,0:未删除,1:已删除
}

// LabelUploadRel 上传数据记录表
type LabelUploadRel struct {
	UploadUUID string    `gorm:"primaryKey;column:upload_uuid;type:varchar(50);not null" json:"upload_uuid,omitempty"` // 该条上传的唯一标识
	UploadType string    `gorm:"column:upload_type;type:varchar(255)" json:"upload_type,omitempty"`                    // 上传类型(labelTag,labelCharacter)
	InfoID     string    `gorm:"column:info_id;type:varchar(50)" json:"info_id,omitempty"`                             // 消息ID，用于异步回调时通知调用者的唯一性标识
	LabelUUID  string    `gorm:"column:label_uuid;type:varchar(50)" json:"label_uuid,omitempty"`                       // 关联数据的唯一标识
	UploadPath string    `gorm:"column:upload_path;type:varchar(255)" json:"upload_path,omitempty"`                    // 上传地址
	CreateTime time.Time `gorm:"column:create_time;type:datetime" json:"create_time,omitempty"`                        // 创建时间
	UpdateTime time.Time `gorm:"column:update_time;type:datetime" json:"update_time,omitempty"`                        // 修改时间
	Status     int       `gorm:"column:status;type:int(1)" json:"status,omitempty"`                                    // 状态:0:未成功;1:成功
	UploadUser string    `gorm:"column:upload_user;type:varchar(50)" json:"upload_user,omitempty"`                     // 上传人
}

// SysConfig 参数配置表
type SysConfig struct {
	ConfigID    int       `gorm:"primaryKey;column:config_id;type:int(5);not null" json:"config_id,omitempty"` // 参数主键
	ConfigName  string    `gorm:"column:config_name;type:varchar(100)" json:"config_name,omitempty"`           // 参数名称
	ConfigKey   string    `gorm:"column:config_key;type:varchar(100)" json:"config_key,omitempty"`             // 参数键名
	ConfigValue string    `gorm:"column:config_value;type:varchar(500)" json:"config_value,omitempty"`         // 参数键值
	ConfigType  string    `gorm:"column:config_type;type:char(1)" json:"config_type,omitempty"`                // 系统内置（Y是 N否）
	CreateBy    string    `gorm:"column:create_by;type:varchar(64)" json:"create_by,omitempty"`                // 创建者
	CreateTime  time.Time `gorm:"column:create_time;type:datetime" json:"create_time,omitempty"`               // 创建时间
	UpdateBy    string    `gorm:"column:update_by;type:varchar(64)" json:"update_by,omitempty"`                // 更新者
	UpdateTime  time.Time `gorm:"column:update_time;type:datetime" json:"update_time,omitempty"`               // 更新时间
	Remark      string    `gorm:"column:remark;type:varchar(500)" json:"remark,omitempty"`                     // 备注
}

// SysCountry 国家信息
type SysCountry struct {
	CountryID           int       `gorm:"primaryKey;column:country_id;type:int(11);not null" json:"country_id,omitempty"`        // 主键
	CountryNameChinese  string    `gorm:"column:country_name_chinese;type:varchar(255)" json:"country_name_chinese,omitempty"`   // 国家中文名称
	CountryCode         string    `gorm:"column:country_code;type:varchar(255)" json:"country_code,omitempty"`                   // 国家代码
	CountryNameEnglish  string    `gorm:"column:country_name_english;type:varchar(255)" json:"country_name_english,omitempty"`   // 国家英文名称
	CountryNationalFlag string    `gorm:"column:country_national_flag;type:varchar(255)" json:"country_national_flag,omitempty"` // 国旗
	CountryUUID         string    `gorm:"column:country_uuid;type:varchar(20)" json:"country_uuid,omitempty"`                    // 唯一标识
	ApplyStatus         int       `gorm:"column:apply_status;type:int(1)" json:"apply_status,omitempty"`                         // 是否启用,0:启用,1:不启用
	DeleteStatus        int       `gorm:"column:delete_status;type:int(1)" json:"delete_status,omitempty"`                       // 是否删除,0:未删除,1:已删除
	CreateTime          time.Time `gorm:"column:create_time;type:datetime" json:"create_time,omitempty"`                         // 创建时间
	UpdateTime          time.Time `gorm:"column:update_time;type:datetime" json:"update_time,omitempty"`                         // 修改时间
	CreateUser          string    `gorm:"column:create_user;type:varchar(50)" json:"create_user,omitempty"`                      // 创建人
	UpdateUser          string    `gorm:"column:update_user;type:varchar(50)" json:"update_user,omitempty"`                      // 修改人
	CountryType         int       `gorm:"column:country_type;type:int(1)" json:"country_type,omitempty"`                         // 类型,0:国家;1:地区
	BelongContinent     string    `gorm:"column:belong_continent;type:varchar(255)" json:"belong_continent,omitempty"`           // 所属大洲
}

// SysDept 部门表
type SysDept struct {
	DeptID     int64     `gorm:"primaryKey;column:dept_id;type:bigint(20);not null" json:"dept_id,omitempty"` // 部门id
	ParentID   int64     `gorm:"column:parent_id;type:bigint(20)" json:"parent_id,omitempty"`                 // 父部门id
	Ancestors  string    `gorm:"column:ancestors;type:varchar(50)" json:"ancestors,omitempty"`                // 祖级列表
	DeptName   string    `gorm:"column:dept_name;type:varchar(30)" json:"dept_name,omitempty"`                // 部门名称
	OrderNum   int       `gorm:"column:order_num;type:int(4)" json:"order_num,omitempty"`                     // 显示顺序
	Leader     string    `gorm:"column:leader;type:varchar(20)" json:"leader,omitempty"`                      // 负责人
	Phone      string    `gorm:"column:phone;type:varchar(11)" json:"phone,omitempty"`                        // 联系电话
	Email      string    `gorm:"column:email;type:varchar(50)" json:"email,omitempty"`                        // 邮箱
	Status     string    `gorm:"column:status;type:char(1)" json:"status,omitempty"`                          // 部门状态（0正常 1停用）
	DelFlag    string    `gorm:"column:del_flag;type:char(1)" json:"del_flag,omitempty"`                      // 删除标志（0代表存在 2代表删除）
	CreateBy   string    `gorm:"column:create_by;type:varchar(64)" json:"create_by,omitempty"`                // 创建者
	CreateTime time.Time `gorm:"column:create_time;type:datetime" json:"create_time,omitempty"`               // 创建时间
	UpdateBy   string    `gorm:"column:update_by;type:varchar(64)" json:"update_by,omitempty"`                // 更新者
	UpdateTime time.Time `gorm:"column:update_time;type:datetime" json:"update_time,omitempty"`               // 更新时间
}

// SysDictData 字典数据表
type SysDictData struct {
	DictCode   int64     `gorm:"primaryKey;column:dict_code;type:bigint(20);not null" json:"dict_code,omitempty"` // 字典编码
	DictSort   int       `gorm:"column:dict_sort;type:int(4)" json:"dict_sort,omitempty"`                         // 字典排序
	DictLabel  string    `gorm:"column:dict_label;type:varchar(100)" json:"dict_label,omitempty"`                 // 字典标签
	DictValue  string    `gorm:"column:dict_value;type:varchar(100)" json:"dict_value,omitempty"`                 // 字典键值
	DictType   string    `gorm:"column:dict_type;type:varchar(100)" json:"dict_type,omitempty"`                   // 字典类型
	CSSClass   string    `gorm:"column:css_class;type:varchar(100)" json:"css_class,omitempty"`                   // 样式属性（其他样式扩展）
	ListClass  string    `gorm:"column:list_class;type:varchar(100)" json:"list_class,omitempty"`                 // 表格回显样式
	IsDefault  string    `gorm:"column:is_default;type:char(1)" json:"is_default,omitempty"`                      // 是否默认（Y是 N否）
	Status     string    `gorm:"column:status;type:char(1)" json:"status,omitempty"`                              // 状态（0正常 1停用）
	CreateBy   string    `gorm:"column:create_by;type:varchar(64)" json:"create_by,omitempty"`                    // 创建者
	CreateTime time.Time `gorm:"column:create_time;type:datetime" json:"create_time,omitempty"`                   // 创建时间
	UpdateBy   string    `gorm:"column:update_by;type:varchar(64)" json:"update_by,omitempty"`                    // 更新者
	UpdateTime time.Time `gorm:"column:update_time;type:datetime" json:"update_time,omitempty"`                   // 更新时间
	Remark     string    `gorm:"column:remark;type:varchar(500)" json:"remark,omitempty"`                         // 备注
}

// SysDictType 字典类型表
type SysDictType struct {
	DictID     int64     `gorm:"primaryKey;column:dict_id;type:bigint(20);not null" json:"dict_id,omitempty"` // 字典主键
	DictName   string    `gorm:"column:dict_name;type:varchar(100)" json:"dict_name,omitempty"`               // 字典名称
	DictType   string    `gorm:"unique;column:dict_type;type:varchar(100)" json:"dict_type,omitempty"`        // 字典类型
	Status     string    `gorm:"column:status;type:char(1)" json:"status,omitempty"`                          // 状态（0正常 1停用）
	CreateBy   string    `gorm:"column:create_by;type:varchar(64)" json:"create_by,omitempty"`                // 创建者
	CreateTime time.Time `gorm:"column:create_time;type:datetime" json:"create_time,omitempty"`               // 创建时间
	UpdateBy   string    `gorm:"column:update_by;type:varchar(64)" json:"update_by,omitempty"`                // 更新者
	UpdateTime time.Time `gorm:"column:update_time;type:datetime" json:"update_time,omitempty"`               // 更新时间
	Remark     string    `gorm:"column:remark;type:varchar(500)" json:"remark,omitempty"`                     // 备注
}

// SysJob 定时任务调度表
type SysJob struct {
	JobID          int64     `gorm:"primaryKey;column:job_id;type:bigint(20);not null" json:"job_id,omitempty"`        // 任务ID
	JobName        string    `gorm:"primaryKey;column:job_name;type:varchar(64);not null" json:"job_name,omitempty"`   // 任务名称
	JobGroup       string    `gorm:"primaryKey;column:job_group;type:varchar(64);not null" json:"job_group,omitempty"` // 任务组名
	InvokeTarget   string    `gorm:"column:invoke_target;type:varchar(500);not null" json:"invoke_target,omitempty"`   // 调用目标字符串
	CronExpression string    `gorm:"column:cron_expression;type:varchar(255)" json:"cron_expression,omitempty"`        // cron执行表达式
	MisfirePolicy  string    `gorm:"column:misfire_policy;type:varchar(20)" json:"misfire_policy,omitempty"`           // 计划执行错误策略（1立即执行 2执行一次 3放弃执行）
	Concurrent     string    `gorm:"column:concurrent;type:char(1)" json:"concurrent,omitempty"`                       // 是否并发执行（0允许 1禁止）
	Status         string    `gorm:"column:status;type:char(1)" json:"status,omitempty"`                               // 状态（0正常 1暂停）
	CreateBy       string    `gorm:"column:create_by;type:varchar(64)" json:"create_by,omitempty"`                     // 创建者
	CreateTime     time.Time `gorm:"column:create_time;type:datetime" json:"create_time,omitempty"`                    // 创建时间
	UpdateBy       string    `gorm:"column:update_by;type:varchar(64)" json:"update_by,omitempty"`                     // 更新者
	UpdateTime     time.Time `gorm:"column:update_time;type:datetime" json:"update_time,omitempty"`                    // 更新时间
	Remark         string    `gorm:"column:remark;type:varchar(500)" json:"remark,omitempty"`                          // 备注信息
}

// SysJobLog 定时任务调度日志表
type SysJobLog struct {
	JobLogID      int64     `gorm:"primaryKey;column:job_log_id;type:bigint(20);not null" json:"job_log_id,omitempty"` // 任务日志ID
	JobName       string    `gorm:"column:job_name;type:varchar(64);not null" json:"job_name,omitempty"`               // 任务名称
	JobGroup      string    `gorm:"column:job_group;type:varchar(64);not null" json:"job_group,omitempty"`             // 任务组名
	InvokeTarget  string    `gorm:"column:invoke_target;type:varchar(500);not null" json:"invoke_target,omitempty"`    // 调用目标字符串
	JobMessage    string    `gorm:"column:job_message;type:varchar(500)" json:"job_message,omitempty"`                 // 日志信息
	Status        string    `gorm:"column:status;type:char(1)" json:"status,omitempty"`                                // 执行状态（0正常 1失败）
	ExceptionInfo string    `gorm:"column:exception_info;type:varchar(2000)" json:"exception_info,omitempty"`          // 异常信息
	CreateTime    time.Time `gorm:"column:create_time;type:datetime" json:"create_time,omitempty"`                     // 创建时间
}

// SysLogininfor 系统访问记录
type SysLogininfor struct {
	InfoID        int64     `gorm:"primaryKey;column:info_id;type:bigint(20);not null" json:"info_id,omitempty"` // 访问ID
	LoginName     string    `gorm:"column:login_name;type:varchar(50)" json:"login_name,omitempty"`              // 登录账号
	IPaddr        string    `gorm:"column:ipaddr;type:varchar(50)" json:"i_paddr,omitempty"`                     // 登录IP地址
	LoginLocation string    `gorm:"column:login_location;type:varchar(255)" json:"login_location,omitempty"`     // 登录地点
	Browser       string    `gorm:"column:browser;type:varchar(50)" json:"browser,omitempty"`                    // 浏览器类型
	Os            string    `gorm:"column:os;type:varchar(50)" json:"os,omitempty"`                              // 操作系统
	Status        string    `gorm:"column:status;type:char(1)" json:"status,omitempty"`                          // 登录状态（0成功 1失败）
	Msg           string    `gorm:"column:msg;type:varchar(255)" json:"msg,omitempty"`                           // 提示消息
	LoginTime     time.Time `gorm:"column:login_time;type:datetime" json:"login_time,omitempty"`                 // 访问时间
}

// SysMenu 菜单权限表
type SysMenu struct {
	MenuID     int64     `gorm:"primaryKey;column:menu_id;type:bigint(20);not null" json:"menu_id,omitempty"` // 菜单ID
	MenuName   string    `gorm:"column:menu_name;type:varchar(50);not null" json:"menu_name,omitempty"`       // 菜单名称
	ParentID   int64     `gorm:"column:parent_id;type:bigint(20)" json:"parent_id,omitempty"`                 // 父菜单ID
	OrderNum   int       `gorm:"column:order_num;type:int(4)" json:"order_num,omitempty"`                     // 显示顺序
	URL        string    `gorm:"column:url;type:varchar(200)" json:"url,omitempty"`                           // 请求地址
	Target     string    `gorm:"column:target;type:varchar(20)" json:"target,omitempty"`                      // 打开方式（menuItem页签 menuBlank新窗口）
	MenuType   string    `gorm:"column:menu_type;type:char(1)" json:"menu_type,omitempty"`                    // 菜单类型（M目录 C菜单 F按钮）
	Visible    string    `gorm:"column:visible;type:char(1)" json:"visible,omitempty"`                        // 菜单状态（0显示 1隐藏）
	Perms      string    `gorm:"column:perms;type:varchar(100)" json:"perms,omitempty"`                       // 权限标识
	Icon       string    `gorm:"column:icon;type:varchar(100)" json:"icon,omitempty"`                         // 菜单图标
	CreateBy   string    `gorm:"column:create_by;type:varchar(64)" json:"create_by,omitempty"`                // 创建者
	CreateTime time.Time `gorm:"column:create_time;type:datetime" json:"create_time,omitempty"`               // 创建时间
	UpdateBy   string    `gorm:"column:update_by;type:varchar(64)" json:"update_by,omitempty"`                // 更新者
	UpdateTime time.Time `gorm:"column:update_time;type:datetime" json:"update_time,omitempty"`               // 更新时间
	Remark     string    `gorm:"column:remark;type:varchar(500)" json:"remark,omitempty"`                     // 备注
}

// SysNotice 通知公告表
type SysNotice struct {
	NoticeID      int       `gorm:"primaryKey;column:notice_id;type:int(4);not null" json:"notice_id,omitempty"` // 公告ID
	NoticeTitle   string    `gorm:"column:notice_title;type:varchar(50);not null" json:"notice_title,omitempty"` // 公告标题
	NoticeType    string    `gorm:"column:notice_type;type:char(1);not null" json:"notice_type,omitempty"`       // 公告类型（1通知 2公告）
	NoticeContent string    `gorm:"column:notice_content;type:varchar(2000)" json:"notice_content,omitempty"`    // 公告内容
	Status        string    `gorm:"column:status;type:char(1)" json:"status,omitempty"`                          // 公告状态（0正常 1关闭）
	CreateBy      string    `gorm:"column:create_by;type:varchar(64)" json:"create_by,omitempty"`                // 创建者
	CreateTime    time.Time `gorm:"column:create_time;type:datetime" json:"create_time,omitempty"`               // 创建时间
	UpdateBy      string    `gorm:"column:update_by;type:varchar(64)" json:"update_by,omitempty"`                // 更新者
	UpdateTime    time.Time `gorm:"column:update_time;type:datetime" json:"update_time,omitempty"`               // 更新时间
	Remark        string    `gorm:"column:remark;type:varchar(255)" json:"remark,omitempty"`                     // 备注
}

// SysOperLog 操作日志记录
type SysOperLog struct {
	OperID        int64     `gorm:"primaryKey;column:oper_id;type:bigint(20);not null" json:"oper_id,omitempty"` // 日志主键
	Title         string    `gorm:"column:title;type:varchar(50)" json:"title,omitempty"`                        // 模块标题
	BusinessType  int       `gorm:"column:business_type;type:int(2)" json:"business_type,omitempty"`             // 业务类型（0其它 1新增 2修改 3删除）
	Method        string    `gorm:"column:method;type:varchar(100)" json:"method,omitempty"`                     // 方法名称
	RequestMethod string    `gorm:"column:request_method;type:varchar(10)" json:"request_method,omitempty"`      // 请求方式
	OperatorType  int       `gorm:"column:operator_type;type:int(1)" json:"operator_type,omitempty"`             // 操作类别（0其它 1后台用户 2手机端用户）
	OperName      string    `gorm:"column:oper_name;type:varchar(50)" json:"oper_name,omitempty"`                // 操作人员
	DeptName      string    `gorm:"column:dept_name;type:varchar(50)" json:"dept_name,omitempty"`                // 部门名称
	OperURL       string    `gorm:"column:oper_url;type:varchar(255)" json:"oper_url,omitempty"`                 // 请求URL
	OperIP        string    `gorm:"column:oper_ip;type:varchar(50)" json:"oper_ip,omitempty"`                    // 主机地址
	OperLocation  string    `gorm:"column:oper_location;type:varchar(255)" json:"oper_location,omitempty"`       // 操作地点
	OperParam     string    `gorm:"column:oper_param;type:varchar(2000)" json:"oper_param,omitempty"`            // 请求参数
	Status        int       `gorm:"column:status;type:int(1)" json:"status,omitempty"`                           // 操作状态（0正常 1异常）
	ErrorMsg      string    `gorm:"column:error_msg;type:varchar(2000)" json:"error_msg,omitempty"`              // 错误消息
	OperTime      time.Time `gorm:"column:oper_time;type:datetime" json:"oper_time,omitempty"`                   // 操作时间
}

// SysPortalUser Portal同步用户数据表
type SysPortalUser struct {
	UserID            int64     `gorm:"primaryKey;column:user_id;type:bigint(20);not null" json:"user_id,omitempty"`      // Id
	OpenID            string    `gorm:"column:open_id;type:varchar(128)" json:"open_id,omitempty"`                        // 用户在钉钉的唯一ID
	LoginName         string    `gorm:"unique;column:login_name;type:varchar(128)" json:"login_name,omitempty"`           // 登录账号
	UserName          string    `gorm:"column:user_name;type:varchar(128)" json:"user_name,omitempty"`                    // 用户昵称
	Email             string    `gorm:"column:email;type:varchar(500)" json:"email,omitempty"`                            // 邮箱
	PhoneNumber       string    `gorm:"column:phone_number;type:varchar(40)" json:"phone_number,omitempty"`               // 手机号
	Avatar            string    `gorm:"column:avatar;type:varchar(500)" json:"avatar,omitempty"`                          // 头像路径
	Status            bool      `gorm:"column:status;type:tinyint(1)" json:"status,omitempty"`                            // 状态：false正常 true停用
	DelFlag           bool      `gorm:"column:del_flag;type:tinyint(1)" json:"del_flag,omitempty"`                        // 删除标志：false 正常 true 删除
	IfAdmin           bool      `gorm:"column:if_admin;type:tinyint(1)" json:"if_admin,omitempty"`                        // 是否为下游系统管理员：false 否， true 是
	Dept              string    `gorm:"column:dept;type:varchar(500)" json:"dept,omitempty"`                              // 组织id与名称对象
	Tenantry          string    `gorm:"column:tenantry;type:varchar(500)" json:"tenantry,omitempty"`                      // 租户id与名称对象
	CreateTime        time.Time `gorm:"column:create_time;type:datetime" json:"create_time,omitempty"`                    // 创建时间
	UpdateTime        time.Time `gorm:"column:update_time;type:datetime" json:"update_time,omitempty"`                    // 修改时间
	DefaultTenantryID string    `gorm:"column:default_tenantry_id;type:varchar(32)" json:"default_tenantry_id,omitempty"` // 默认租户id
	UnionID           string    `gorm:"column:union_id;type:varchar(128)" json:"union_id,omitempty"`                      // 钉钉unionId
	DingUserID        string    `gorm:"column:ding_user_id;type:varchar(128)" json:"ding_user_id,omitempty"`              // 钉钉用户id
	PenName           string    `gorm:"column:pen_name;type:varchar(128)" json:"pen_name,omitempty"`                      // 用户笔名
}

// SysPost 岗位信息表
type SysPost struct {
	PostID     int64     `gorm:"primaryKey;column:post_id;type:bigint(20);not null" json:"post_id,omitempty"` // 岗位ID
	PostCode   string    `gorm:"column:post_code;type:varchar(64);not null" json:"post_code,omitempty"`       // 岗位编码
	PostName   string    `gorm:"column:post_name;type:varchar(50);not null" json:"post_name,omitempty"`       // 岗位名称
	PostSort   int       `gorm:"column:post_sort;type:int(4);not null" json:"post_sort,omitempty"`            // 显示顺序
	Status     string    `gorm:"column:status;type:char(1);not null" json:"status,omitempty"`                 // 状态（0正常 1停用）
	CreateBy   string    `gorm:"column:create_by;type:varchar(64)" json:"create_by,omitempty"`                // 创建者
	CreateTime time.Time `gorm:"column:create_time;type:datetime" json:"create_time,omitempty"`               // 创建时间
	UpdateBy   string    `gorm:"column:update_by;type:varchar(64)" json:"update_by,omitempty"`                // 更新者
	UpdateTime time.Time `gorm:"column:update_time;type:datetime" json:"update_time,omitempty"`               // 更新时间
	Remark     string    `gorm:"column:remark;type:varchar(500)" json:"remark,omitempty"`                     // 备注
}

// SysRole 角色信息表
type SysRole struct {
	RoleID     int64     `gorm:"primaryKey;column:role_id;type:bigint(20);not null" json:"role_id,omitempty"` // 角色ID
	RoleName   string    `gorm:"column:role_name;type:varchar(30);not null" json:"role_name,omitempty"`       // 角色名称
	RoleKey    string    `gorm:"column:role_key;type:varchar(100);not null" json:"role_key,omitempty"`        // 角色权限字符串
	RoleSort   int       `gorm:"column:role_sort;type:int(4);not null" json:"role_sort,omitempty"`            // 显示顺序
	DataScope  string    `gorm:"column:data_scope;type:char(1)" json:"data_scope,omitempty"`                  // 数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）
	Status     string    `gorm:"column:status;type:char(1);not null" json:"status,omitempty"`                 // 角色状态（0正常 1停用）
	DelFlag    string    `gorm:"column:del_flag;type:char(1)" json:"del_flag,omitempty"`                      // 删除标志（0代表存在 2代表删除）
	CreateBy   string    `gorm:"column:create_by;type:varchar(64)" json:"create_by,omitempty"`                // 创建者
	CreateTime time.Time `gorm:"column:create_time;type:datetime" json:"create_time,omitempty"`               // 创建时间
	UpdateBy   string    `gorm:"column:update_by;type:varchar(64)" json:"update_by,omitempty"`                // 更新者
	UpdateTime time.Time `gorm:"column:update_time;type:datetime" json:"update_time,omitempty"`               // 更新时间
	Remark     string    `gorm:"column:remark;type:varchar(500)" json:"remark,omitempty"`                     // 备注
}

// SysRoleDept 角色和部门关联表
type SysRoleDept struct {
	RoleID int64 `gorm:"primaryKey;column:role_id;type:bigint(20);not null" json:"role_id,omitempty"` // 角色ID
	DeptID int64 `gorm:"primaryKey;column:dept_id;type:bigint(20);not null" json:"dept_id,omitempty"` // 部门ID
}

// SysRoleMenu 角色和菜单关联表
type SysRoleMenu struct {
	RoleID int64 `gorm:"primaryKey;column:role_id;type:bigint(20);not null" json:"role_id,omitempty"` // 角色ID
	MenuID int64 `gorm:"primaryKey;column:menu_id;type:bigint(20);not null" json:"menu_id,omitempty"` // 菜单ID
}

// SysUser 用户信息表
type SysUser struct {
	UserID       int64     `gorm:"primaryKey;column:user_id;type:bigint(20);not null" json:"user_id,omitempty"` // 用户ID
	DeptID       int64     `gorm:"column:dept_id;type:bigint(20)" json:"dept_id,omitempty"`                     // 部门ID
	LoginName    string    `gorm:"column:login_name;type:varchar(30);not null" json:"login_name,omitempty"`     // 登录账号
	UserName     string    `gorm:"column:user_name;type:varchar(30);not null" json:"user_name,omitempty"`       // 用户昵称
	UserType     string    `gorm:"column:user_type;type:varchar(2)" json:"user_type,omitempty"`                 // 用户类型（00系统用户）
	Email        string    `gorm:"column:email;type:varchar(50)" json:"email,omitempty"`                        // 用户邮箱
	Phonenumber  string    `gorm:"column:phonenumber;type:varchar(11)" json:"phonenumber,omitempty"`            // 手机号码
	Sex          string    `gorm:"column:sex;type:char(1)" json:"sex,omitempty"`                                // 用户性别（0男 1女 2未知）
	Avatar       string    `gorm:"column:avatar;type:varchar(100)" json:"avatar,omitempty"`                     // 头像路径
	Password     string    `gorm:"column:password;type:varchar(50)" json:"password,omitempty"`                  // 密码
	Salt         string    `gorm:"column:salt;type:varchar(20)" json:"salt,omitempty"`                          // 盐加密
	Status       string    `gorm:"column:status;type:char(1)" json:"status,omitempty"`                          // 帐号状态（0正常 1停用）
	DelFlag      string    `gorm:"column:del_flag;type:char(1)" json:"del_flag,omitempty"`                      // 删除标志（0代表存在 2代表删除）
	LoginIP      string    `gorm:"column:login_ip;type:varchar(50)" json:"login_ip,omitempty"`                  // 最后登陆IP
	LoginDate    time.Time `gorm:"column:login_date;type:datetime" json:"login_date,omitempty"`                 // 最后登陆时间
	CreateBy     string    `gorm:"column:create_by;type:varchar(64)" json:"create_by,omitempty"`                // 创建者
	CreateTime   time.Time `gorm:"column:create_time;type:datetime" json:"create_time,omitempty"`               // 创建时间
	UpdateBy     string    `gorm:"column:update_by;type:varchar(64)" json:"update_by,omitempty"`                // 更新者
	UpdateTime   time.Time `gorm:"column:update_time;type:datetime" json:"update_time,omitempty"`               // 更新时间
	Remark       string    `gorm:"column:remark;type:varchar(500)" json:"remark,omitempty"`                     // 备注
	PortalUserID int64     `gorm:"column:portal_user_id;type:bigint(20)" json:"portal_user_id,omitempty"`       // portal用户表主键
}

// SysUserCms 标签系统与cms系统用户ID关联
type SysUserCms struct {
	PearUserID string    `gorm:"primaryKey;column:pear_user_id;type:varchar(20);not null" json:"pear_user_id,omitempty"` // 标签系统用户ID
	CmsUserID  string    `gorm:"primaryKey;column:cms_user_id;type:varchar(20);not null" json:"cms_user_id,omitempty"`   // cms系统用户ID
	CreateTime time.Time `gorm:"column:create_time;type:datetime" json:"create_time,omitempty"`                          // 创建时间
	UpdateTime time.Time `gorm:"column:update_time;type:datetime" json:"update_time,omitempty"`                          // 修改时间
	CreateUser string    `gorm:"column:create_user;type:varchar(20)" json:"create_user,omitempty"`                       // 创建人
	UpdateUser string    `gorm:"column:update_user;type:varchar(20)" json:"update_user,omitempty"`                       // 修改人
}

// SysUserDictData 用户与字典数据关联表
type SysUserDictData struct {
	UserID   int64 `gorm:"primaryKey;column:user_id;type:bigint(20);not null" json:"user_id,omitempty"`     // 用户ID
	DictCode int64 `gorm:"primaryKey;column:dict_code;type:bigint(20);not null" json:"dict_code,omitempty"` // 字典数据表ID
}

// SysUserOnline 在线用户记录
type SysUserOnline struct {
	SessionID      string    `gorm:"primaryKey;column:sessionId;type:varchar(50);not null" json:"session_id,omitempty"` // 用户会话id
	LoginName      string    `gorm:"column:login_name;type:varchar(50)" json:"login_name,omitempty"`                    // 登录账号
	DeptName       string    `gorm:"column:dept_name;type:varchar(50)" json:"dept_name,omitempty"`                      // 部门名称
	IPaddr         string    `gorm:"column:ipaddr;type:varchar(50)" json:"i_paddr,omitempty"`                           // 登录IP地址
	LoginLocation  string    `gorm:"column:login_location;type:varchar(255)" json:"login_location,omitempty"`           // 登录地点
	Browser        string    `gorm:"column:browser;type:varchar(50)" json:"browser,omitempty"`                          // 浏览器类型
	Os             string    `gorm:"column:os;type:varchar(50)" json:"os,omitempty"`                                    // 操作系统
	Status         string    `gorm:"column:status;type:varchar(10)" json:"status,omitempty"`                            // 在线状态on_line在线off_line离线
	StartTimestamp time.Time `gorm:"column:start_timestamp;type:datetime" json:"start_timestamp,omitempty"`             // session创建时间
	LastAccessTime time.Time `gorm:"column:last_access_time;type:datetime" json:"last_access_time,omitempty"`           // session最后访问时间
	ExpireTime     int       `gorm:"column:expire_time;type:int(5)" json:"expire_time,omitempty"`                       // 超时时间，单位为分钟
}

// SysUserPost 用户与岗位关联表
type SysUserPost struct {
	UserID int64 `gorm:"primaryKey;column:user_id;type:bigint(20);not null" json:"user_id,omitempty"` // 用户ID
	PostID int64 `gorm:"primaryKey;column:post_id;type:bigint(20);not null" json:"post_id,omitempty"` // 岗位ID
}

// SysUserRole 用户和角色关联表
type SysUserRole struct {
	UserID int64 `gorm:"primaryKey;column:user_id;type:bigint(20);not null" json:"user_id,omitempty"` // 用户ID
	RoleID int64 `gorm:"primaryKey;column:role_id;type:bigint(20);not null" json:"role_id,omitempty"` // 角色ID
}

// TableXiaaohui [...]
type TableXiaaohui struct {
	ID   int16  `gorm:"primaryKey;column:id;type:smallint(6);not null" json:"id,omitempty"`
	Name string `gorm:"column:name;type:varchar(100)" json:"name,omitempty"`
	Item string `gorm:"column:item;type:varchar(255)" json:"item,omitempty"`
}
