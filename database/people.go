package database

import (
	"encoding/json"

	"github.com/ServiceComputingGroup/simpleWebServer/entity"
	"github.com/boltdb/bolt"
)

func GetPerson(key string) string {
	key = "https://swapi.co/api/people/" + key + "/"
	var str string
	var data entity.People
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(people)
		cur := b.Cursor()

		for k, v := cur.First(); k != nil; k, v = cur.Next() {
			json.Unmarshal(v, &data)
			if data.Url == key {
				str = string(v)
				return nil
			}
		}
		return nil
	})
	return str
}

func GetPeople() []entity.People {

	var result []entity.People
	var data entity.People
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(people)
		cur := b.Cursor()

		for k, v := cur.First(); k != nil; k, v = cur.Next() {
			json.Unmarshal(v, &data)
			result = append(result, data)
		}
		return nil
	})

	return result
}
