package graph_ql

import (
	"fmt"

	"github.com/ServiceComputingGroup/simpleWebServer/database"
	"github.com/ServiceComputingGroup/simpleWebServer/entity"
	"github.com/graphql-go/graphql"
)

func ModifyuserFieldConfig() *graphql.Field {
	var field = &graphql.Field{
		Type:        graphql.Boolean,
		Description: "用户更改",
		Args: graphql.FieldConfigArgument{
			"username": &graphql.ArgumentConfig{
				Description: "用户名称",
				Type:        graphql.NewNonNull(graphql.String),
			},
			"email": &graphql.ArgumentConfig{
				Description: "用户邮箱",
				Type:        graphql.NewNonNull(graphql.String),
			},
			"password": &graphql.ArgumentConfig{
				Description: "用户密码",
				Type:        graphql.NewNonNull(graphql.String),
			},
			"phone": &graphql.ArgumentConfig{
				Description: "用户电话",
				Type:        graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			user := &entity.User{
				UserName: p.Args["username"].(string),
				Email:    p.Args["email"].(string),
				Password: p.Args["password"].(string),
				Phone:    p.Args["phone"].(string),
			}
			if err := database.UpdateUser(user); err != nil {
				return false, err
			}
			fmt.Println(user)
			return true, nil

		},
	}
	return field
}
