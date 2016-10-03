package gq_test

import (
	"encoding/json"
	"fmt"

	"github.com/SDkie/employee_graphql_sample/data"
	"github.com/SDkie/employee_graphql_sample/db"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("createEmployee", func() {

	type createEmployeeResponse struct {
		Data struct {
			Employee data.Employee `json:"createEmployee"`
		} `json:"data"`

		Errors []gqError `json:"errors"`
	}

	BeforeEach(func() {
		setup()
		Expect(db.GetDb().Create(&dept).Error).NotTo(HaveOccurred())
		emp.DeptNo = dept.DeptNo
	})

	Context("Sending valid graphql query", func() {

		It("User should create successfully", func() {
			query := `
		mutation {
			createEmployee(ENAME:"%s", JOB:"%s", MGR:%d, SALARY:%f, DEPTNO:%d){
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

			query = fmt.Sprintf(query, emp.EName, emp.Job, emp.Mgr, emp.Salary, dept.DeptNo)

			resp := sendRequest(query)
			response := new(createEmployeeResponse)
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

	Context("Not Sending valid DeptNo", func() {

		It("User creation should fail", func() {
			query := `
		mutation {
			createEmployee(ENAME:"%s", JOB:"%s", MGR:%d, SALARY:%f, DEPTNO:%d){
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

			query = fmt.Sprintf(query, emp.EName, emp.Job, emp.Mgr, emp.Salary, -1)
			resp := sendRequest(query)

			response := new(createEmployeeResponse)
			err := json.Unmarshal(resp.Body.Bytes(), response)

			Expect(err).NotTo(HaveOccurred())
			Expect(response.Errors).ShouldNot(HaveLen(0))
		})
	})

	AfterEach(func() {
		db.Close()
	})
})
