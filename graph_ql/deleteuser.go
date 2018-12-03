package graph_ql

import (
	"github.com/ServiceComputingGroup/simpleWebServer/database"
	"github.com/graphql-go/graphql"
)

func DeleteuserFieldConfig() *graphql.Field {
	var field = &graphql.Field{
		Type:        graphql.Boolean,
		Description: "用户删除",
		Args: graphql.FieldConfigArgument{
			"username": &graphql.ArgumentConfig{
				Description: "用户名称",
				Type:        graphql.NewNonNull(graphql.String),
			},
		},

		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			username := p.Args["username"].(string)
			if err := database.DeleteUser(username); err != nil {
				return false, err
			}
			return true, nil

		},
	}
	return field
}
