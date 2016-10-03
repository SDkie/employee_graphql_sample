package gq_test

import (
	"encoding/json"
	"fmt"

	"github.com/SDkie/employee_graphql_sample/data"
	"github.com/SDkie/employee_graphql_sample/db"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("getEmployee Graph Query", func() {

	type getEmployeeResponse struct {
		Data struct {
			Employee data.Employee `json:"getEmployee"`
		} `json:"data"`

		Errors []gqError `json:"errors"`
	}

	BeforeEach(func() {
		setup()
		Expect(db.GetDb().Create(&dept).Error).NotTo(HaveOccurred())
		emp.DeptNo = dept.DeptNo
		Expect(db.GetDb().Create(&emp).Error).NotTo(HaveOccurred())
	})

	Context("Sending valid graphql query", func() {

		It("We should get valid user response", func() {
			query := `
		query {
			getEmployee(EMPNO:%d){
				ENAME,
				JOB,
				MGR,
				SALARY,
				DEPT {
					DEPTNO,
					DNAME,
					LOC
				}
			}
		}`

			query = fmt.Sprintf(query, emp.EmpNo)

			resp := sendRequest(query)
			response := new(getEmployeeResponse)
			err := json.Unmarshal(resp.Body.Bytes(), response)

			Expect(err).NotTo(HaveOccurred())
			Expect(response.Errors).Should(HaveLen(0))
			empDb := response.Data.Employee
			Expect(empDb.EName).Should(Equal(emp.EName))
			Expect(empDb.Job).Should(Equal(emp.Job))
			Expect(empDb.Mgr).Should(Equal(emp.Mgr))
			Expect(empDb.Salary).Should(Equal(emp.Salary))
			Expect(empDb.Dept.Dname).Should(Equal(dept.Dname))
			Expect(empDb.Dept.DeptNo).Should(Equal(dept.DeptNo))
			Expect(empDb.Dept.Loc).Should(Equal(dept.Loc))
		})
	})

	Context("Sending Invalid EMPNO in query", func() {

		It("Should fail with Not Found error", func() {
			query := `
		query {
			getEmployee{
				ENAME,
				JOB,
				MGR,
				SALARY,
				DEPT {
					DEPTNO,
					DNAME,
					LOC
				}
			}
		}`

			resp := sendRequest(query)
			response := new(getEmployeeResponse)
			err := json.Unmarshal(resp.Body.Bytes(), response)
			Expect(err).NotTo(HaveOccurred())

			Expect(response.Errors).ShouldNot(HaveLen(0))
		})
	})

	AfterEach(func() {
		db.Close()
	})
})

var _ = Describe("listOfAllEmployees Graph Query", func() {
	type listOfAllEmployeeResponse struct {
		Data struct {
			Employees []data.Employee `json:"listOfAllEmployees"`
		} `json:"data"`

		Errors []gqError `json:"errors"`
	}

	BeforeEach(func() {
		setup()
		Expect(db.GetDb().Create(&dept).Error).NotTo(HaveOccurred())
		emp.DeptNo = dept.DeptNo
		Expect(db.GetDb().Create(&emp).Error).NotTo(HaveOccurred())
		emp.EmpNo = 0
		Expect(db.GetDb().Create(&emp).Error).NotTo(HaveOccurred())
	})

	Context("Sending valid graphql query", func() {

		It("Should get list of all users", func() {
			query := `
		query {
			listOfAllEmployees{
				EMPNO
				ENAME,
				JOB,
				MGR,
				SALARY,
				DEPT {
					DEPTNO,
					DNAME,
					LOC
				}
			}
		}`
			resp := sendRequest(query)
			response := new(listOfAllEmployeeResponse)
			Expect(json.Unmarshal(resp.Body.Bytes(), response)).ShouldNot(HaveOccurred())
			Expect(response.Errors).Should(HaveLen(0))
			Expect(response.Data.Employees).Should(HaveLen(2))
			for i, dbEmp := range response.Data.Employees {
				Expect(dbEmp.EmpNo).Should(Equal(i + 1))
				Expect(dbEmp.EName).Should(Equal(emp.EName))
				Expect(dbEmp.Job).Should(Equal(emp.Job))
				Expect(dbEmp.Mgr).Should(Equal(emp.Mgr))
				Expect(dbEmp.Salary).Should(Equal(emp.Salary))
				Expect(dbEmp.Dept.Dname).Should(Equal(dept.Dname))
				Expect(dbEmp.Dept.DeptNo).Should(Equal(dept.DeptNo))
				Expect(dbEmp.Dept.Loc).Should(Equal(dept.Loc))
			}
		})
	})

	Context("Not sending subsection of DEPT", func() {
		It("Should get error from graphql", func() {
			query := `
		query {
			listOfAllEmployees{
				EMPNO
				ENAME,
				JOB,
				MGR,
				SALARY,
				DEPT
			}
		}`

			resp := sendRequest(query)
			response := new(listOfAllEmployeeResponse)
			Expect(json.Unmarshal(resp.Body.Bytes(), response)).ShouldNot(HaveOccurred())
			Expect(response.Errors).ShouldNot(HaveLen(0))
		})
	})

	AfterEach(func() {
		db.Close()
	})
})
