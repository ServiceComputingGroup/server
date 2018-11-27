package graph_ql

import "github.com/graphql-go/graphql"

var peopleType = graphql.NewObject(
	graphql.ObjectConfig{
		Fields: graphql.Fields{},
	},
)

var peopleListType = graphql.NewList(peopleType)
