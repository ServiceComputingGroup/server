package graph_ql

import (
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
			if ok && claims != nil {
				exp, isExist := claims.(jwt.MapClaims)["exp"]
				if !isExist {
					message = "Token Invalid"
				} else if time.Now().UTC().Unix() > (int64)(exp.(float64)) {
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
