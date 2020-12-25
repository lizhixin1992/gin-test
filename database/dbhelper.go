package database

import (
	"fmt"
	"github.com/lizhixin1992/gin-test/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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

	masterEngine = db
	return masterEngine
}
