package db

import (
	log "github.com/Sirupsen/logrus"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

// Initialize MYSQL Db
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

// Close MYSQL Db
func Close() {
	err := db.Close()
	if err != nil {
		log.Errorf("MYSQL Close Error: %s", err)
	}
	log.Info("MYSQL Close: DONE")
}

// Get MySql Db Connection
func GetDb() *gorm.DB {
	return db
}
