package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/lizhixin1992/gin-test/conf"
	"log"
	"sync"
)

var (
	masterEngine *sql.DB
	lock         sync.Mutex
)

func InstanceMaster() *sql.DB {
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
	driverSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
		c.UserName, c.Password, c.Host, c.Port, c.DbName)

	db, err := sql.Open(conf.DriverName, driverSource)
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	//defer db.Close()

	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	} else {
		log.Println("db open is success")
	}

	masterEngine = db
	return masterEngine
}
