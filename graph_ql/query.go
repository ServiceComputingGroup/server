package graph_ql

import "github.com/graphql-go/graphql"

/*
query user message
*/
var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name:   "Query",
		Fields: graphql.Fields{},
	},
)
