package gq

import (
	"github.com/SDkie/employee_graphql_sample/gq/resolvers"
	"github.com/SDkie/employee_graphql_sample/gq/types"
	"github.com/graphql-go/graphql"
)

var mutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"createEmployee": &graphql.Field{
			Type:        types.Employee,
			Description: "Creates a new Employee record",
			Args: graphql.FieldConfigArgument{
				"ENAME": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"JOB": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"MGR": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
				"SALARY": &graphql.ArgumentConfig{
					Type: graphql.Float,
				},
				"DEPTNO": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: resolvers.CreateEmployee,
		},
	},
})
