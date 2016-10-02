package data

import (
	"github.com/SDkie/employee_graphql_sample/db"
	log "github.com/Sirupsen/logrus"
)

// This function will setup Database tables
func Init() {
	var err error
	if err = db.GetDb().AutoMigrate(&Employee{}).Error; err != nil {
		log.Errorf("AutoMigrate Error: %s", err)
	}
	if err = db.GetDb().AutoMigrate(&Department{}).Error; err != nil {
		log.Errorf("AutoMigrate Error: %s", err)
	}

	db.GetDb().Model(&Employee{}).AddForeignKey("dept_no", "department(dept_no)", "RESTRICT", "RESTRICT")

	log.Info("Data Init: DONE")
}
