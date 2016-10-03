package data_test

import (
	. "github.com/SDkie/employee_graphql_sample/data"
	"github.com/SDkie/employee_graphql_sample/db"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Creating Employee", func() {
	BeforeEach(func() {
		setup()
		Expect(db.GetDb().Create(&dept).Error).NotTo(HaveOccurred())
		emp.DeptNo = dept.DeptNo
	})

	Context("Valid User Creation", func() {
		It("User Should Create Successfully", func() {
			newEmp, err := CreateEmployee(emp.EName, emp.Job, emp.Mgr, emp.Salary, emp.DeptNo)
			Expect(err).NotTo(HaveOccurred())
			Expect(newEmp.EName).Should(Equal(emp.EName))
			Expect(newEmp.Job).Should(Equal(emp.Job))
			Expect(newEmp.Mgr).Should(Equal(emp.Mgr))
			Expect(newEmp.Salary).Should(Equal(emp.Salary))
			Expect(newEmp.DeptNo).Should(Equal(emp.DeptNo))
		})

	})

	Context("User Creation with invalid DeptNo", func() {
		It("User Creation should fail", func() {
			_, err := CreateEmployee(emp.EName, emp.Job, emp.Mgr, emp.Salary, -1)
			Expect(err).Should(HaveOccurred())
		})
	})

	AfterEach(func() {
		db.Close()
	})
})

var _ = Describe("Reading Employee Data", func() {
	BeforeEach(func() {
		setup()
		Expect(db.GetDb().Create(&dept).Error).NotTo(HaveOccurred())
		emp.DeptNo = dept.DeptNo
		Expect(db.GetDb().Create(&emp).Error).NotTo(HaveOccurred())
		emp.EName += "1"
		emp.EmpNo = 0
		emp.DeptNo = dept.DeptNo
		Expect(db.GetDb().Create(&emp).Error).NotTo(HaveOccurred())
	})

	Context("Reading User List", func() {
		It("Should Get All DB Users", func() {
			emps, err := ListOfAllEmployees()
			Expect(err).NotTo(HaveOccurred())
			Expect(emps).Should(HaveLen(2))
		})
	})

	Context("Reading Single User", func() {
		It("Should Get Single DB User", func() {
			newEmp, err := GetEmployeeByEmpNo(emp.EmpNo)
			Expect(err).NotTo(HaveOccurred())
			Expect(newEmp.EmpNo).Should(Equal(emp.EmpNo))
			Expect(newEmp.EName).Should(Equal(emp.EName))
			Expect(newEmp.Job).Should(Equal(emp.Job))
			Expect(newEmp.Mgr).Should(Equal(emp.Mgr))
			Expect(newEmp.Salary).Should(Equal(emp.Salary))
			Expect(newEmp.DeptNo).Should(Equal(emp.DeptNo))
		})
	})

	Context("Reading User List by DName", func() {
		It("Should Get All DB Users", func() {
			emps, err := ListOfAllEmployeesByDname(dept.Dname)
			Expect(err).NotTo(HaveOccurred())
			Expect(emps).Should(HaveLen(2))
		})
	})

	AfterEach(func() {
		db.Close()
	})
})
