package graph_ql

import (
	"github.com/ServiceComputingGroup/simpleWebServer/database"
	"github.com/graphql-go/graphql"
)

func Queryuser() *graphql.Field {
	var field = &graphql.Field{
		Type:        graphql.String,
		Description: "用户查询",
		Args: graphql.FieldConfigArgument{
			"username": &graphql.ArgumentConfig{
				Description: "用户名称",
				Type:        graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			username := p.Args["username"].(string)
			user, err := database.GetUser(username)
			if err != true {
				return "Fail to find user", nil
			} else {
				return user.Password + "," + user.Phone + "," + user.Email, nil
			}
		},
	}
	return field
}
