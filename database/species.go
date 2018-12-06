package database

import (
	"encoding/json"

	"github.com/ServiceComputingGroup/simpleWebServer/entity"
	"github.com/boltdb/bolt"
)

func GetSpecies(key string) string {
	key = "https://swapi.co/api/species/" + key + "/"
	var str string
	var data entity.Species
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(species)
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

func GetAllSpecies() []string {

	var result []string

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(species)
		cur := b.Cursor()

		for k, v := cur.First(); k != nil; k, v = cur.Next() {

			result = append(result, string(v[:]))
		}
		return nil
	})

	return result
}
