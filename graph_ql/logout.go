package graph_ql

import (
	"time"

	"github.com/ServiceComputingGroup/simpleWebServer/database"
	"github.com/graphql-go/graphql"
)

func LogoutFieldConfig() *graphql.Field {
	var field = &graphql.Field{
		Type:        graphql.String,
		Description: "用户登出",
		Args: graphql.FieldConfigArgument{
			"username": &graphql.ArgumentConfig{
				Description: "用户名称",
				Type:        graphql.NewNonNull(graphql.String),
			},
		},

		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			username := p.Args["username"].(string)

			if _, err := database.GetUser(username); err != nil {
				return "", err
			} else {
				var userinfo map[string]interface{}
				userinfo["username"] = username
				userinfo["exp"] = time.Now().Unix()
				userinfo["iat"] = time.Now().Unix()
				key := "UserToken"
				token := createToken(key, userinfo)
				return token, nil
			}
		},
	}
	return field
}
