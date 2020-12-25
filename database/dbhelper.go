package database

import (
	"fmt"
	"github.com/lizhixin1992/gin-test/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
	"sync"
)

var (
	masterEngine *gorm.DB
	lock         sync.Mutex
)

func InstanceMaster() *gorm.DB {
	if masterEngine != nil {
		return masterEngine
	}

	lock.Lock()
	defer lock.Unlock()

	//再次判断，确保不重复创建
	if masterEngine != nil {
		return masterEngine
	}

	c := conf.MasterDbConf
	driverSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true",
		c.UserName, c.Password, c.Host, c.Port, c.DbName)

	db, err := gorm.Open(mysql.Open(driverSource), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	//多数据库连接，主/从，读写分离，负载均衡
	driverSource1 := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true",
		c.UserName, c.Password, c.Host, c.Port, "tag_test_1")

	driverSource2 := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true",
		c.UserName, c.Password, c.Host, c.Port, "tag_test_2")

	db.Use(dbresolver.Register(dbresolver.Config{
		Sources:  []gorm.Dialector{mysql.Open(driverSource)},
		Replicas: []gorm.Dialector{mysql.Open(driverSource1), mysql.Open(driverSource2)},
		// sources/replicas 负载均衡策略
		Policy: dbresolver.RandomPolicy{},
	}))

	masterEngine = db
	return masterEngine
}
