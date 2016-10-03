package gq_test

import (
	"net/http"
	"net/http/httptest"
	"os"

	"github.com/SDkie/employee_graphql_sample/data"
	"github.com/SDkie/employee_graphql_sample/db"
	"github.com/SDkie/employee_graphql_sample/gq"
	p "github.com/SDkie/employee_graphql_sample/preferences"
	log "github.com/Sirupsen/logrus"

	. "github.com/onsi/gomega"
)

var (
	dept data.Department
	emp  data.Employee
)

func setup() {
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
	handler := http.HandlerFunc(gq.GraphQlHandler)
	handler.ServeHTTP(resp, req)
	log.Debugln(query)
	log.Debugln(resp.Body)
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
