package data

import "github.com/SDkie/employee_graphql_sample/db"

type Department struct {
	DeptNo int    `json:"DEPTNO" sql:"dept_no" gorm:"primary_key"`
	Dname  string `json:"DNAME" sql:"d_name" gorm:"not null;unique"`
	Loc    string `json:"LOC" sql:"loc"`
}

// Get Department using DeptNo
func GetDepartmentByDeptNo(deptNo int) (*Department, error) {
	dept := new(Department)
	err := db.GetDb().Where(&Department{DeptNo: deptNo}).First(dept).Error

	return dept, err
}

// Get Department By Department Name
func GetDepartmentByDname(dname string) (*Department, error) {
	dept := new(Department)
	err := db.GetDb().Where(&Department{Dname: dname}).First(dept).Error
	return dept, err
}
