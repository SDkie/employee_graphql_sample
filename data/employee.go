package data

import (
	"errors"

	"github.com/SDkie/employee_graphql_sample/db"
)

type Employee struct {
	EmpNo  int         `json:"EMPNO" sql:"emp_no" gorm:"primary_key"`
	EName  string      `json:"ENAME" sql:"e_name" gorm:"not null"`
	Job    string      `json:"JOB" sql:"job"`
	Mgr    int         `json:"MGR" sql:"mgr"`
	Salary float32     `json:"SALARY" sql:"salary"`
	DeptNo int         `json:"DEPTNO" sql:"dept_no" gorm:"not null"`
	Dept   *Department `json:"DEPT" sql:"-"`
}

// Get Employee using EmployeeNo
func GetEmployeeByEmpNo(empNo int) (*Employee, error) {
	emp := new(Employee)
	err := db.GetDb().Where(&Employee{EmpNo: empNo}).First(emp).Error
	if err != nil {
		return nil, err
	}

	emp.Dept, err = GetDepartmentByDeptNo(emp.DeptNo)
	return emp, err
}

// Get List of All Employees By Dname
func ListOfAllEmployeesByDname(dname string) ([]Employee, error) {
	dept, err := GetDepartmentByDname(dname)
	if err != nil {
		return nil, err
	}

	emps := []Employee{}
	err = db.GetDb().Where(&Employee{DeptNo: dept.DeptNo}).Find(&emps).Error
	if err != nil {
		return nil, err
	}

	for i, _ := range emps {
		emps[i].Dept = dept
	}

	return emps, nil
}

// Get List of All Employees
func ListOfAllEmployees() ([]Employee, error) {
	emps := []Employee{}
	err := db.GetDb().Find(&emps).Error
	if err != nil {
		return nil, err
	}

	for i, emp := range emps {
		emps[i].Dept, err = GetDepartmentByDeptNo(emp.DeptNo)
		if err != nil {
			return nil, err
		}
	}

	return emps, nil
}

func CreateEmployee(eName string, job string, mgr int, salary float32, deptNo int) (*Employee, error) {
	// Check if deptNo is valid or not
	dept, err := GetDepartmentByDeptNo(deptNo)
	if err != nil {
		err = errors.New("Invalid DeptNo")
		return nil, err
	}

	emp := new(Employee)
	emp.EName = eName
	emp.Job = job
	emp.Mgr = mgr
	emp.Salary = salary
	emp.DeptNo = deptNo
	emp.Dept = dept
	err = db.GetDb().Create(emp).Error
	return emp, err
}

// Update Employee
func UpdateEmployee(empNo int, eName string, job string, mgr int, salary float32, deptNo int) (*Employee, error) {
	// Check if deptNo is valid or not
	_, err := GetDepartmentByDeptNo(deptNo)
	if err != nil {
		err = errors.New("Invalid DeptNo")
		return nil, err
	}

	// Check if empNo is valid or not
	emp, err := GetEmployeeByEmpNo(empNo)
	if err != nil {
		err = errors.New("Invalid EmpNo")
		return nil, err
	}

	oldDeptNo := emp.DeptNo
	emp.EName = eName
	emp.Job = job
	emp.Mgr = mgr
	emp.Salary = salary
	emp.DeptNo = deptNo
	err = db.GetDb().Save(emp).Error
	if err != nil {
		return nil, err
	}

	// If DeptNo is changed then only update emp.Dept
	if oldDeptNo != emp.DeptNo {
		emp.Dept, err = GetDepartmentByDeptNo(emp.DeptNo)
	}

	return emp, err
}

// Delete Employee of given EmployeeNo
func DeleteEmployeeWithEmpNo(empNo int) (*Employee, error) {
	// Check if empNo is valid or not
	emp, err := GetEmployeeByEmpNo(empNo)
	if err != nil {
		err = errors.New("Invalid EmpNo")
		return nil, err
	}

	err = db.GetDb().Delete(&Employee{EmpNo: empNo}).Error
	return emp, err
}
