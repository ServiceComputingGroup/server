package database

import (
	"encoding/json"

	"github.com/ServiceComputingGroup/simpleWebServer/entity"
	"github.com/boltdb/bolt"
)

func GetVehicle(key string) string {
	key = "https://swapi.co/api/vehicles/" + key + "/"
	var str string
	var data entity.Vehicle
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(vehicle)
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

func GetVehicles() []string {

	var result []string

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(vehicle)
		cur := b.Cursor()

		for k, v := cur.First(); k != nil; k, v = cur.Next() {

			result = append(result, string(v[:]))
		}
		return nil
	})

	return result
}
