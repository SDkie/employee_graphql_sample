package db

import (
	log "github.com/Sirupsen/logrus"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func Init(mysqlUrl string) error {
	var err error
	if db, err = gorm.Open("mysql", mysqlUrl); err != nil {
		log.Errorf("MYSQL Init Error: %s", err)
		return err
	}
	db.LogMode(true)
	db.SingularTable(true)

	log.Info("MYSQL Init : DONE")
	return nil
}

func Close() {
	err := db.Close()
	if err != nil {
		log.Errorf("MYSQL Close Error: %s", err)
	}
	log.Info("MYSQL Close: DONE")
}

func GetDb() *gorm.DB {
	return db
}
