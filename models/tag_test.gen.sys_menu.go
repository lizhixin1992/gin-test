package	models	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
"time"	
)	

type _SysMenuMgr struct {
	*_BaseMgr
}

// SysMenuMgr open func
func SysMenuMgr(db *gorm.DB) *_SysMenuMgr {
	if db == nil {
		panic(fmt.Errorf("SysMenuMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SysMenuMgr{_BaseMgr: &_BaseMgr{DB: db.Table("sys_menu"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SysMenuMgr) GetTableName() string {
	return "sys_menu"
}

// Get 获取
func (obj *_SysMenuMgr) Get() (result SysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_SysMenuMgr) Gets() (results []*SysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithMenuID menu_id获取 菜单ID
func (obj *_SysMenuMgr) WithMenuID(menuID int64) Option {
	return optionFunc(func(o *options) { o.query["menu_id"] = menuID })
}

// WithMenuName menu_name获取 菜单名称
func (obj *_SysMenuMgr) WithMenuName(menuName string) Option {
	return optionFunc(func(o *options) { o.query["menu_name"] = menuName })
}

// WithParentID parent_id获取 父菜单ID
func (obj *_SysMenuMgr) WithParentID(parentID int64) Option {
	return optionFunc(func(o *options) { o.query["parent_id"] = parentID })
}

// WithOrderNum order_num获取 显示顺序
func (obj *_SysMenuMgr) WithOrderNum(orderNum int) Option {
	return optionFunc(func(o *options) { o.query["order_num"] = orderNum })
}

// WithURL url获取 请求地址
func (obj *_SysMenuMgr) WithURL(url string) Option {
	return optionFunc(func(o *options) { o.query["url"] = url })
}

// WithTarget target获取 打开方式（menuItem页签 menuBlank新窗口）
func (obj *_SysMenuMgr) WithTarget(target string) Option {
	return optionFunc(func(o *options) { o.query["target"] = target })
}

// WithMenuType menu_type获取 菜单类型（M目录 C菜单 F按钮）
func (obj *_SysMenuMgr) WithMenuType(menuType string) Option {
	return optionFunc(func(o *options) { o.query["menu_type"] = menuType })
}

// WithVisible visible获取 菜单状态（0显示 1隐藏）
func (obj *_SysMenuMgr) WithVisible(visible string) Option {
	return optionFunc(func(o *options) { o.query["visible"] = visible })
}

// WithPerms perms获取 权限标识
func (obj *_SysMenuMgr) WithPerms(perms string) Option {
	return optionFunc(func(o *options) { o.query["perms"] = perms })
}

// WithIcon icon获取 菜单图标
func (obj *_SysMenuMgr) WithIcon(icon string) Option {
	return optionFunc(func(o *options) { o.query["icon"] = icon })
}

// WithCreateBy create_by获取 创建者
func (obj *_SysMenuMgr) WithCreateBy(createBy string) Option {
	return optionFunc(func(o *options) { o.query["create_by"] = createBy })
}

// WithCreateTime create_time获取 创建时间
func (obj *_SysMenuMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateBy update_by获取 更新者
func (obj *_SysMenuMgr) WithUpdateBy(updateBy string) Option {
	return optionFunc(func(o *options) { o.query["update_by"] = updateBy })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_SysMenuMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithRemark remark获取 备注
func (obj *_SysMenuMgr) WithRemark(remark string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}


// GetByOption 功能选项模式获取
func (obj *_SysMenuMgr) GetByOption(opts ...Option) (result SysMenu, err error) {
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
func (obj *_SysMenuMgr) GetByOptions(opts ...Option) (results []*SysMenu, err error) {
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


// GetFromMenuID 通过menu_id获取内容 菜单ID 
func (obj *_SysMenuMgr)  GetFromMenuID(menuID int64) (result SysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("menu_id = ?", menuID).Find(&result).Error
	
	return
}

// GetBatchFromMenuID 批量唯一主键查找 菜单ID
func (obj *_SysMenuMgr) GetBatchFromMenuID(menuIDs []int64) (results []*SysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("menu_id IN (?)", menuIDs).Find(&results).Error
	
	return
}
 
// GetFromMenuName 通过menu_name获取内容 菜单名称 
func (obj *_SysMenuMgr) GetFromMenuName(menuName string) (results []*SysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("menu_name = ?", menuName).Find(&results).Error
	
	return
}

// GetBatchFromMenuName 批量唯一主键查找 菜单名称
func (obj *_SysMenuMgr) GetBatchFromMenuName(menuNames []string) (results []*SysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("menu_name IN (?)", menuNames).Find(&results).Error
	
	return
}
 
// GetFromParentID 通过parent_id获取内容 父菜单ID 
func (obj *_SysMenuMgr) GetFromParentID(parentID int64) (results []*SysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("parent_id = ?", parentID).Find(&results).Error
	
	return
}

// GetBatchFromParentID 批量唯一主键查找 父菜单ID
func (obj *_SysMenuMgr) GetBatchFromParentID(parentIDs []int64) (results []*SysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("parent_id IN (?)", parentIDs).Find(&results).Error
	
	return
}
 
// GetFromOrderNum 通过order_num获取内容 显示顺序 
func (obj *_SysMenuMgr) GetFromOrderNum(orderNum int) (results []*SysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("order_num = ?", orderNum).Find(&results).Error
	
	return
}

// GetBatchFromOrderNum 批量唯一主键查找 显示顺序
func (obj *_SysMenuMgr) GetBatchFromOrderNum(orderNums []int) (results []*SysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("order_num IN (?)", orderNums).Find(&results).Error
	
	return
}
 
// GetFromURL 通过url获取内容 请求地址 
func (obj *_SysMenuMgr) GetFromURL(url string) (results []*SysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("url = ?", url).Find(&results).Error
	
	return
}

// GetBatchFromURL 批量唯一主键查找 请求地址
func (obj *_SysMenuMgr) GetBatchFromURL(urls []string) (results []*SysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("url IN (?)", urls).Find(&results).Error
	
	return
}
 
// GetFromTarget 通过target获取内容 打开方式（menuItem页签 menuBlank新窗口） 
func (obj *_SysMenuMgr) GetFromTarget(target string) (results []*SysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("target = ?", target).Find(&results).Error
	
	return
}

// GetBatchFromTarget 批量唯一主键查找 打开方式（menuItem页签 menuBlank新窗口）
func (obj *_SysMenuMgr) GetBatchFromTarget(targets []string) (results []*SysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("target IN (?)", targets).Find(&results).Error
	
	return
}
 
// GetFromMenuType 通过menu_type获取内容 菜单类型（M目录 C菜单 F按钮） 
func (obj *_SysMenuMgr) GetFromMenuType(menuType string) (results []*SysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("menu_type = ?", menuType).Find(&results).Error
	
	return
}

// GetBatchFromMenuType 批量唯一主键查找 菜单类型（M目录 C菜单 F按钮）
func (obj *_SysMenuMgr) GetBatchFromMenuType(menuTypes []string) (results []*SysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("menu_type IN (?)", menuTypes).Find(&results).Error
	
	return
}
 
// GetFromVisible 通过visible获取内容 菜单状态（0显示 1隐藏） 
func (obj *_SysMenuMgr) GetFromVisible(visible string) (results []*SysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("visible = ?", visible).Find(&results).Error
	
	return
}

// GetBatchFromVisible 批量唯一主键查找 菜单状态（0显示 1隐藏）
func (obj *_SysMenuMgr) GetBatchFromVisible(visibles []string) (results []*SysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("visible IN (?)", visibles).Find(&results).Error
	
	return
}
 
// GetFromPerms 通过perms获取内容 权限标识 
func (obj *_SysMenuMgr) GetFromPerms(perms string) (results []*SysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("perms = ?", perms).Find(&results).Error
	
	return
}

// GetBatchFromPerms 批量唯一主键查找 权限标识
func (obj *_SysMenuMgr) GetBatchFromPerms(permss []string) (results []*SysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("perms IN (?)", permss).Find(&results).Error
	
	return
}
 
// GetFromIcon 通过icon获取内容 菜单图标 
func (obj *_SysMenuMgr) GetFromIcon(icon string) (results []*SysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("icon = ?", icon).Find(&results).Error
	
	return
}

// GetBatchFromIcon 批量唯一主键查找 菜单图标
func (obj *_SysMenuMgr) GetBatchFromIcon(icons []string) (results []*SysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("icon IN (?)", icons).Find(&results).Error
	
	return
}
 
// GetFromCreateBy 通过create_by获取内容 创建者 
func (obj *_SysMenuMgr) GetFromCreateBy(createBy string) (results []*SysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_by = ?", createBy).Find(&results).Error
	
	return
}

// GetBatchFromCreateBy 批量唯一主键查找 创建者
func (obj *_SysMenuMgr) GetBatchFromCreateBy(createBys []string) (results []*SysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_by IN (?)", createBys).Find(&results).Error
	
	return
}
 
// GetFromCreateTime 通过create_time获取内容 创建时间 
func (obj *_SysMenuMgr) GetFromCreateTime(createTime time.Time) (results []*SysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_time = ?", createTime).Find(&results).Error
	
	return
}

// GetBatchFromCreateTime 批量唯一主键查找 创建时间
func (obj *_SysMenuMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*SysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("create_time IN (?)", createTimes).Find(&results).Error
	
	return
}
 
// GetFromUpdateBy 通过update_by获取内容 更新者 
func (obj *_SysMenuMgr) GetFromUpdateBy(updateBy string) (results []*SysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_by = ?", updateBy).Find(&results).Error
	
	return
}

// GetBatchFromUpdateBy 批量唯一主键查找 更新者
func (obj *_SysMenuMgr) GetBatchFromUpdateBy(updateBys []string) (results []*SysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_by IN (?)", updateBys).Find(&results).Error
	
	return
}
 
// GetFromUpdateTime 通过update_time获取内容 更新时间 
func (obj *_SysMenuMgr) GetFromUpdateTime(updateTime time.Time) (results []*SysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_time = ?", updateTime).Find(&results).Error
	
	return
}

// GetBatchFromUpdateTime 批量唯一主键查找 更新时间
func (obj *_SysMenuMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*SysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("update_time IN (?)", updateTimes).Find(&results).Error
	
	return
}
 
// GetFromRemark 通过remark获取内容 备注 
func (obj *_SysMenuMgr) GetFromRemark(remark string) (results []*SysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("remark = ?", remark).Find(&results).Error
	
	return
}

// GetBatchFromRemark 批量唯一主键查找 备注
func (obj *_SysMenuMgr) GetBatchFromRemark(remarks []string) (results []*SysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("remark IN (?)", remarks).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_SysMenuMgr) FetchByPrimaryKey(menuID int64 ) (result SysMenu, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("menu_id = ?", menuID).Find(&result).Error
	
	return
}
 

 

	

