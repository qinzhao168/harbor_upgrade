package db

import (
	"fmt"

	"harbor_upgrade/src/config"

	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func init() {
	newDBClient1(config.GetConfig().Db1.Type)
	newDBClient5(config.GetConfig().Db5.Type)
}

var db1 *gorm.DB
var db5 *gorm.DB

func newDBClient1(t string) {
	var err error
	switch t {
	case "mysql":
		glog.V(2).Info("init mysql...")
		addr := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
			config.GetConfig().Db1.Username,
			config.GetConfig().Db1.Password,
			config.GetConfig().Db1.Addr,
			config.GetConfig().Db1.Database)
		db1, err = gorm.Open("mysql", addr)
		if err != nil {
			fmt.Printf("mysql db1 connect error: ", err)
			return
		}

		db1.DB().SetMaxIdleConns(10)
		db1.DB().SetMaxOpenConns(1000)
		db1.LogMode(true)
	default:
		fmt.Printf("Not support db type %s", t)

	}
}

func GetDB1() *gorm.DB {
	return db1
}

func newDBClient5(t string) {
	var err error
	switch t {
	case "mysql":
		glog.V(2).Info("init mysql...")
		addr := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
			config.GetConfig().Db5.Username,
			config.GetConfig().Db5.Password,
			config.GetConfig().Db5.Addr,
			config.GetConfig().Db5.Database)
		db5, err = gorm.Open("mysql", addr)
		if err != nil {
			fmt.Printf("mysql db5 connect error: ", err)
			return
		}

		db5.DB().SetMaxIdleConns(10)
		db5.DB().SetMaxOpenConns(1000)
		db5.LogMode(true)
	default:
		fmt.Printf("Not support db type %s", t)

	}
}

func GetDB5() *gorm.DB {
	return db5
}
