package data_test

import (
	"os"

	"github.com/SDkie/employee_graphql_sample/db"
	p "github.com/SDkie/employee_graphql_sample/preferences"
	log "github.com/Sirupsen/logrus"

	. "github.com/SDkie/employee_graphql_sample/data"
	. "github.com/onsi/gomega"
)

var (
	dept Department
	emp  Employee
)

func setup() {
	os.Setenv("ENV", "test")
	log.SetLevel(log.ErrorLevel)

	p.Init("../config.ini")
	mysqlURL := p.GetMysqlURL()
	Expect(db.Init(mysqlURL)).NotTo(HaveOccurred())
	db.GetDb().LogMode(false)
	Expect(db.GetDb().DropTableIfExists(Employee{}).Error).NotTo(HaveOccurred())
	Expect(db.GetDb().DropTableIfExists(Department{}).Error).NotTo(HaveOccurred())
	Init()

	dept.Dname = "Software development"
	dept.Loc = "Pune"

	emp.EName = "QWERTY"
	emp.Job = "Backend Engineer"
	emp.Mgr = 0
	emp.Salary = 100.50
}
