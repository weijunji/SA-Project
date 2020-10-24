package db

import (
	"log"

	"github.com/henrylee2cn/erpc/v6"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var _db *gorm.DB

func init() {
	erpc.Infof("Init database")
	var err error
	_db, err = gorm.Open(sqlite.Open("sap.sqlite"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
}

//GetDB get inited db connection
func GetDB() *gorm.DB {
	return _db
}
