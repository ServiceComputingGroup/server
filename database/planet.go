package database

import (
	"encoding/json"

	"github.com/ServiceComputingGroup/simpleWebServer/entity"
	"github.com/boltdb/bolt"
)

func GetPlanet(key string) string {
	key = "https://swapi.co/api/planets/" + key + "/"
	var str string
	var data entity.Planet
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(planet)
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

func GetPlanets() []string {
	var numbers []int
	numbers = append(numbers, 2, 3, 4)

	var result []string

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(planet)
		cur := b.Cursor()

		for k, v := cur.First(); k != nil; k, v = cur.Next() {

			result = append(result, string(v[:]))
		}
		return nil
	})

	return result
}
