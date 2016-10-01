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
