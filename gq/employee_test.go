package gq_test

import (
	"encoding/json"
	"fmt"

	. "github.com/SDkie/employee_graphql_sample/gq"

	"net/http"
	"net/http/httptest"
	"os"

	"github.com/SDkie/employee_graphql_sample/data"
	"github.com/SDkie/employee_graphql_sample/db"
	p "github.com/SDkie/employee_graphql_sample/preferences"
	log "github.com/Sirupsen/logrus"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	dept data.Department
	emp  data.Employee
)

func testingSetup() {
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

	emp.EName = "QWERTY"
	emp.Job = "Backend Engineer"
	emp.Mgr = 0
	emp.Salary = 100.50

}

func sendRequest(query string) *httptest.ResponseRecorder {
	req, err := http.NewRequest("GET", "/graphql", nil)
	Expect(err).NotTo(HaveOccurred())
	urlQuery := req.URL.Query()

	urlQuery.Set("query", query)
	req.URL.RawQuery = urlQuery.Encode()

	resp := httptest.NewRecorder()
	handler := http.HandlerFunc(GraphQlHandler)
	handler.ServeHTTP(resp, req)
	return resp
}

type gqLocation struct {
	Line   int `json:"line"`
	Column int `json:"column"`
}

type gqError struct {
	Message   string       `json:"message"`
	Locations []gqLocation `json:"locations"`
}

var _ = Describe("createEmployee Graph Query", func() {

	type createEmployeeResponse struct {
		Data struct {
			CreateEmployee data.Employee `json:"createEmployee"`
		} `json:"data"`

		Errors []gqError `json:"errors"`
	}

	BeforeEach(func() {
		testingSetup()
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
			Expect(response.Data.CreateEmployee.EName).Should(Equal(emp.EName))
			Expect(response.Data.CreateEmployee.Job).Should(Equal(emp.Job))
			Expect(response.Data.CreateEmployee.Mgr).Should(Equal(emp.Mgr))
			Expect(response.Data.CreateEmployee.Salary).Should(Equal(emp.Salary))
			Expect(response.Data.CreateEmployee.Dept.Dname).Should(Equal(dept.Dname))
			Expect(response.Data.CreateEmployee.Dept.DeptNo).Should(Equal(dept.DeptNo))
			Expect(response.Data.CreateEmployee.Dept.Loc).Should(Equal(dept.Loc))
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

var _ = Describe("getEmployee Graph Query", func() {

	type getEmployeeResponse struct {
		Data struct {
			CreateEmployee data.Employee `json:"getEmployee"`
		} `json:"data"`

		Errors []gqError `json:"errors"`
	}

	BeforeEach(func() {
		testingSetup()
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
			Expect(response.Data.CreateEmployee.EName).Should(Equal(emp.EName))
			Expect(response.Data.CreateEmployee.Job).Should(Equal(emp.Job))
			Expect(response.Data.CreateEmployee.Mgr).Should(Equal(emp.Mgr))
			Expect(response.Data.CreateEmployee.Salary).Should(Equal(emp.Salary))
			Expect(response.Data.CreateEmployee.Dept.Dname).Should(Equal(dept.Dname))
			Expect(response.Data.CreateEmployee.Dept.DeptNo).Should(Equal(dept.DeptNo))
			Expect(response.Data.CreateEmployee.Dept.Loc).Should(Equal(dept.Loc))
		})
	})

	Context("Sending Invalid EMPNO in query", func() {

		It("Not sending EMPNO in request", func() {
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
