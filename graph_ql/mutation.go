package graph_ql

import (
	"github.com/graphql-go/graphql"
)

var mutationType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"register":   RegisterFieldConfig(),
			"login":      LoginFieldConfig(),
			"logout":     LogoutFieldConfig(),
			"deleteuser": DeleteuserFieldConfig(),
			"modifyuser": ModifyuserFieldConfig(),
		},
	},
)
