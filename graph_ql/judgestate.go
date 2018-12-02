package graph_ql

import (
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/graphql-go/graphql"
)

func Judgestate() *graphql.Field {
	var field = &graphql.Field{
		Type:        graphql.String,
		Description: "Token判断",
		Args: graphql.FieldConfigArgument{
			"token": &graphql.ArgumentConfig{
				Description: "用户Token",
				Type:        graphql.NewNonNull(graphql.String),
			},
		},

		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			key := "UserToken"
			message := "Token Valid."
			tokenstring := p.Args["token"].(string)
			claims, ok := parseToken(tokenstring, key)
			if ok {
				exp, _ := strconv.ParseInt(claims.(jwt.MapClaims)["exp"].(string), 10, 64)
				if time.Now().UTC().Unix() > exp {
					message = "Token Out of Data."
				} else {
					message = "Token Valid."
				}
			} else {
				message = "Token Invalid."
			}
			return message, nil
		},
	}
	return field
}
