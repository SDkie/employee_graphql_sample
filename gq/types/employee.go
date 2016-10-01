package types

import "github.com/graphql-go/graphql"

var Employee = graphql.NewObject(graphql.ObjectConfig{
	Name: "Employee",
	Fields: graphql.Fields{
		"EMPNO": &graphql.Field{
			Type:        graphql.ID,
			Description: "Employee Number",
		},
		"ENAME": &graphql.Field{
			Type:        graphql.String,
			Description: "Employee Name",
		},
		"JOB": &graphql.Field{
			Type:        graphql.String,
			Description: "Employee Job Title",
		},
		"MGR": &graphql.Field{
			Type:        graphql.Int,
			Description: "Employee Manager Id",
		},
		"SALARY": &graphql.Field{
			Type:        graphql.Float,
			Description: "Employee Salary",
		},
	}})
