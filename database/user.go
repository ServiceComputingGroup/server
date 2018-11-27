package database

import (
	"github.com/liping/simpleWebServer/entity"
)

func SaveUser(usrs []entity.User) {
	return
}
func LoadUser() []entity.User {
	var res = make([]entity.User, 0)
	return res
}

func SaveToken(tokens []string) {

}

func LoadToken() []string {
	return make([]string, 0)
}
