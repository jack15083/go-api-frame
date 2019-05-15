package models

import (
	"../core"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var dbDefault *gorm.DB
var BaseModel baseModel

type baseModel struct{}

func (baseModel) ConnectDB(DBName string) *gorm.DB {
	dBConf := core.Config.Database[DBName]
	db, err := gorm.Open(dBConf.DriverName, dBConf.DataSourceName)
	if err != nil {
		panic("failed to connect database:" + err.Error())
	}

	return db
}

func (baseModel) CloseDB(db *gorm.DB) {
	db.Close()
}
