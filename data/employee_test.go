package data_test

import (
	"os"

	"github.com/SDkie/employee_graphql_sample/data"
	"github.com/SDkie/employee_graphql_sample/db"
	p "github.com/SDkie/employee_graphql_sample/preferences"
	log "github.com/Sirupsen/logrus"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Creating Employee", func() {
	var dept data.Department
	var emp data.Employee

	BeforeEach(func() {
		os.Setenv("ENV", "test")
		log.SetLevel(log.ErrorLevel)

		p.Init("../config.ini")
		mysqlURL := p.GetMysqlURL()
		Expect(db.Init(mysqlURL)).NotTo(HaveOccurred())
		db.GetDb().LogMode(false)
		Expect(db.GetDb().DropTableIfExists(data.Employee{}).Error).NotTo(HaveOccurred())
		Expect(db.GetDb().DropTableIfExists(data.Department{}).Error).NotTo(HaveOccurred())
		data.Init()

		dept.Dname = "Software development"
		dept.Loc = "Pune"
		Expect(db.GetDb().Create(&dept).Error).NotTo(HaveOccurred())

		emp.EName = "QWERTY"
		emp.DeptNo = dept.DeptNo
		emp.Job = "Backend Engineer"
		emp.Mgr = 0
		emp.Salary = 100.50
	})

	Context("Valid User Creation", func() {
		It("User Should Create Successfully", func() {
			_, err := data.CreateEmployee(emp.EName, emp.Job, emp.Mgr, emp.Salary, emp.DeptNo)
			Expect(err).NotTo(HaveOccurred())
		})

	})

	Context("Invalid User Creation", func() {
		It("User Creation should fail", func() {
			_, err := data.CreateEmployee(emp.EName, emp.Job, emp.Mgr, emp.Salary, -1)
			Expect(err).Should(HaveOccurred())
		})
	})

	AfterEach(func() {
		db.Close()
	})
})
