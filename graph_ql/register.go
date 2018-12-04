package graph_ql

import (
	"fmt"
	"time"

	"github.com/ServiceComputingGroup/simpleWebServer/database"
	"github.com/ServiceComputingGroup/simpleWebServer/entity"
	"github.com/graphql-go/graphql"
)

func RegisterFieldConfig() *graphql.Field {
	var field = &graphql.Field{
		Type:        graphql.String,
		Description: "用户注册",
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

			if isInsert := database.InsertUser(user); !isInsert {
				return "", nil
			}

			userinfo := make(map[string]interface{})
			userinfo["username"] = user.UserName
			userinfo["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
			userinfo["iat"] = time.Now().Unix()
			fmt.Println(userinfo["username"])
			fmt.Println(userinfo["exp"])
			fmt.Println(userinfo["iat"])

			key := "UserToken"
			token := createToken(key, userinfo)

			return token, nil
		},
	}
	return field
}
