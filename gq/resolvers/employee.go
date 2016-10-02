package resolvers

import (
	"github.com/SDkie/employee_graphql_sample/data"
	log "github.com/Sirupsen/logrus"
	"github.com/graphql-go/graphql"
)

func GetEmployee(params graphql.ResolveParams) (interface{}, error) {
	// EMPNO validation is done by graphql
	empNo := params.Args["EMPNO"].(int)

	emp, err := data.GetEmployeeByEmpNo(empNo)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return emp, nil
}

func ListOfAllEmployees(params graphql.ResolveParams) (interface{}, error) {
	emps, err := data.ListOfAllEmployees()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return emps, err
}

func ListOfAllEmployeesByDname(params graphql.ResolveParams) (interface{}, error) {
	dname := params.Args["DNAME"].(string)

	emps, err := data.ListOfAllEmployeesByDname(dname)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return emps, nil
}

func CreateEmployee(params graphql.ResolveParams) (interface{}, error) {
	var eName, job string
	var salary float32
	var mgr, deptNo int

	// Compulsory
	eName = params.Args["ENAME"].(string)
	deptNo = params.Args["DEPTNO"].(int)

	// Optional
	temp, ok := params.Args["JOB"]
	if ok {
		job, _ = temp.(string)
	}
	temp, ok = params.Args["MGR"]
	if ok {
		mgr, _ = temp.(int)
	}
	temp, ok = params.Args["SALARY"]
	if ok {
		salary, _ = temp.(float32)
	}

	emp, err := data.CreateEmployee(eName, job, mgr, salary, deptNo)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return emp, nil
}
