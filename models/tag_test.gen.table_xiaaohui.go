package	models	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _TableXiaaohuiMgr struct {
	*_BaseMgr
}

// TableXiaaohuiMgr open func
func TableXiaaohuiMgr(db *gorm.DB) *_TableXiaaohuiMgr {
	if db == nil {
		panic(fmt.Errorf("TableXiaaohuiMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_TableXiaaohuiMgr{_BaseMgr: &_BaseMgr{DB: db.Table("table_xiaaohui"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_TableXiaaohuiMgr) GetTableName() string {
	return "table_xiaaohui"
}

// Get 获取
func (obj *_TableXiaaohuiMgr) Get() (result TableXiaaohui, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_TableXiaaohuiMgr) Gets() (results []*TableXiaaohui, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 
func (obj *_TableXiaaohuiMgr) WithID(id int16) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取 
func (obj *_TableXiaaohuiMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithItem item获取 
func (obj *_TableXiaaohuiMgr) WithItem(item string) Option {
	return optionFunc(func(o *options) { o.query["item"] = item })
}


// GetByOption 功能选项模式获取
func (obj *_TableXiaaohuiMgr) GetByOption(opts ...Option) (result TableXiaaohui, err error) {
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
func (obj *_TableXiaaohuiMgr) GetByOptions(opts ...Option) (results []*TableXiaaohui, err error) {
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


// GetFromID 通过id获取内容  
func (obj *_TableXiaaohuiMgr)  GetFromID(id int16) (result TableXiaaohui, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error
	
	return
}

// GetBatchFromID 批量唯一主键查找 
func (obj *_TableXiaaohuiMgr) GetBatchFromID(ids []int16) (results []*TableXiaaohui, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error
	
	return
}
 
// GetFromName 通过name获取内容  
func (obj *_TableXiaaohuiMgr) GetFromName(name string) (results []*TableXiaaohui, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error
	
	return
}

// GetBatchFromName 批量唯一主键查找 
func (obj *_TableXiaaohuiMgr) GetBatchFromName(names []string) (results []*TableXiaaohui, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error
	
	return
}
 
// GetFromItem 通过item获取内容  
func (obj *_TableXiaaohuiMgr) GetFromItem(item string) (results []*TableXiaaohui, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("item = ?", item).Find(&results).Error
	
	return
}

// GetBatchFromItem 批量唯一主键查找 
func (obj *_TableXiaaohuiMgr) GetBatchFromItem(items []string) (results []*TableXiaaohui, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("item IN (?)", items).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_TableXiaaohuiMgr) FetchByPrimaryKey(id int16 ) (result TableXiaaohui, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error
	
	return
}
 

 

	

