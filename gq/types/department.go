package types

import "github.com/graphql-go/graphql"

var Department = graphql.NewObject(graphql.ObjectConfig{
	Name: "Department",
	Fields: graphql.Fields{
		"DEPTNO": &graphql.Field{
			Type: graphql.Int,
		},
		"DNAME": &graphql.Field{
			Type: graphql.String,
		},
		"LOC": &graphql.Field{
			Type: graphql.String,
		},
	},
})
