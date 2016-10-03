package types

import "github.com/graphql-go/graphql"

var Department = graphql.NewObject(graphql.ObjectConfig{
	Name: "Department",
	Fields: graphql.Fields{
		"DEPTNO": &graphql.Field{
			Type:        graphql.Int,
			Description: "Department Number",
		},
		"DNAME": &graphql.Field{
			Type:        graphql.String,
			Description: "Department Name",
		},
		"LOC": &graphql.Field{
			Type:        graphql.String,
			Description: "Department Location",
		},
	},
})
