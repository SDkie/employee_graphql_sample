package gq

import (
	"github.com/SDkie/employee_graphql_sample/gq/resolvers"
	"github.com/SDkie/employee_graphql_sample/gq/types"
	"github.com/graphql-go/graphql"
)

var query = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"getEmployee": &graphql.Field{
			Type:        types.Employee,
			Description: "Gets a Employee record based on the EMPNO",
			Args: graphql.FieldConfigArgument{
				"EMPNO": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: resolvers.GetEmployee,
		},
		"listOfAllEmployees": &graphql.Field{
			Type:        graphql.NewList(types.Employee),
			Description: "Gets a list of all the employees",
			Resolve:     resolvers.ListOfAllEmployees,
		},
	},
})
