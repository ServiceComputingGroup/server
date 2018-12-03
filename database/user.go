package database

import (
	"encoding/json"
	"fmt"

	"github.com/ServiceComputingGroup/simpleWebServer/entity"
	"github.com/boltdb/bolt"
)

func InsertUser(user *entity.User) bool {
	_, err := GetUser(user.UserName)
	if err == nil {
		return false
	}
	if UpdateUser(user) == nil {
		return true
	}
	return false
}
func DeleteUser(username string) error {
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(userB)

		return b.Delete([]byte(username))
	})
	return err

}
func UpdateUser(user *entity.User) error {
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(userB)

		encoded, err := json.Marshal(user)
		if err != nil {
			return err
		}
		return b.Put([]byte(user.UserName), encoded)
	})
	return err
}
func GetUser(username string) (*entity.User, error) {
	fmt.Println("GetUser")
	var user *entity.User

	err := db.View(func(tx *bolt.Tx) error {

		b := tx.Bucket(userB)

		val := b.Get([]byte(username))
		err := json.Unmarshal(val, &user)
		return err
	})
	return user, err

}
