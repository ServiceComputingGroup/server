package graph_ql

import (
	"fmt"
	"time"

	"github.com/ServiceComputingGroup/simpleWebServer/database"
	"github.com/graphql-go/graphql"
)

func LoginFieldConfig() *graphql.Field {
	var field = &graphql.Field{
		Type:        graphql.String,
		Description: "用户登录",
		Args: graphql.FieldConfigArgument{
			"username": &graphql.ArgumentConfig{
				Description: "用户名称",
				Type:        graphql.NewNonNull(graphql.String),
			},
			"password": &graphql.ArgumentConfig{
				Description: "用户密码",
				Type:        graphql.NewNonNull(graphql.String),
			},
		},

		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			username := p.Args["username"].(string)
			password := p.Args["password"].(string)
			user, err := database.GetUser(username)
			if err != nil {
				return "No such user.", err
			} else if password != user.Password {
				return "Password incorrect.", nil
			} else {
				userinfo := make(map[string]interface{})
				userinfo["username"] = username
				userinfo["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
				userinfo["iat"] = time.Now().Unix()

				fmt.Println(userinfo)
				key := "UserToken"
				token := createToken(key, userinfo)
				return token, nil
			}
		},
	}
	return field
}
