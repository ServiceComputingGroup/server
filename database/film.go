package database

import (
	"encoding/json"

	"github.com/ServiceComputingGroup/simpleWebServer/entity"
	"github.com/boltdb/bolt"
)

func GetFilm(key string) string {

	key = "https://swapi.co/api/films/" + key + "/"
	var str string
	var data entity.Film
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(film)
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

func GetFilms() []entity.Film {

	var result []entity.Film
	var data entity.Film
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(film)
		cur := b.Cursor()

		for k, v := cur.First(); k != nil; k, v = cur.Next() {
			json.Unmarshal(v, &data)
			result = append(result, data)
		}
		return nil
	})

	return result
}
