package data

import (
	"github.com/SDkie/employee_graphql_sample/db"
	log "github.com/Sirupsen/logrus"
)

type Employee struct {
	EmpNo  int     `json:"EMPNO" sql:"emp_no" gorm:"primary_key"`
	EName  string  `json:"ENAME" sql:"e_name"`
	Job    string  `json:"JOB" sql:"job"`
	Mgr    int     `json:"MGR" sql:"mgr"`
	Salary float32 `json:"SALARY" sql:"salary"`
	DeptNo int     `json:"DEPTNO" sql:"dept_no"`
}

// Get Employee using EmployeeNo
func GetEmployeeByEmpNo(empNo int) (*Employee, error) {
	emp := new(Employee)
	log.Debug("EMPNO:", empNo)
	err := db.GetDb().Where(&Employee{EmpNo: empNo}).First(emp).Error
	return emp, err
}

// Get List of All Employees
func ListOfAllEmployees() ([]Employee, error) {
	emps := new([]Employee)
	err := db.GetDb().Find(&emps).Error
	return *emps, err
}
