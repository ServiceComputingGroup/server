package graph_ql

import (
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/graphql-go/graphql"
)

/*
query user message
*/
var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"judgestate": Judgestate(),
			"query":      Queryinfo(),
		},
	},
)

func createToken(key string, m map[string]interface{}) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)

	for index, val := range m {
		claims[index] = val
	}
	token.Claims = claims
	tokenString, _ := token.SignedString([]byte(key))
	return tokenString
}

func parseToken(tokenString string, key string) (interface{}, bool) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("UserToken"), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	} else {
		fmt.Println(err)
		return "", false
	}
}
