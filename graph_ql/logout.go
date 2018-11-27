package graph_ql

import (
	"github.com/graphql-go/graphql"
)

func LogoutFieldConfig() *graphql.Field {
	var field = &graphql.Field{
		//Name: "logout",
		Type: graphql.String,
		Args: graphql.FieldConfigArgument{
			"username": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},

		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return p.Args["username"].(string) + "logout successfully", nil
		},
	}
	return field
}
