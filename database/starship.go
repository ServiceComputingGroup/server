package database

import (
	"encoding/json"

	"github.com/ServiceComputingGroup/simpleWebServer/entity"
	"github.com/boltdb/bolt"
)

func GetStartship(key string) string {
	key = "https://swapi.co/api/starships/" + key + "/"
	var str string
	var data entity.Starship
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(starship)
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

func GetStartships() []string {

	var result []string

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(starship)
		cur := b.Cursor()

		for k, v := cur.First(); k != nil; k, v = cur.Next() {

			result = append(result, string(v[:]))
		}
		return nil
	})

	return result
}
