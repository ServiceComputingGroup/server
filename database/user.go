package database

import (
	"encoding/json"
	"fmt"

	"github.com/ServiceComputingGroup/simpleWebServer/entity"
	"github.com/boltdb/bolt"
)

func InsertUser(user *entity.User) bool {

	_, has := GetUser(user.UserName)

	if has == true {
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
	fmt.Println("UpdateUser")
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
func GetUser(username string) (*entity.User, bool) {

	var user *entity.User
	var val []byte
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(userB)
		val = b.Get([]byte(username))

		//val := b.Get([]byte(username))
		return nil
	})
	if val == nil {
		return user, false
	} else {
		json.Unmarshal(val, &user)
		return user, true
	}
}
