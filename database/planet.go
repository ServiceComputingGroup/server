package database

import (
	"github.com/boltdb/bolt"
)

func GetPlanet(key string) string {
	key = "https://swapi.co/api/planets/" + key + "/"
	k := []byte(key)
	var val []byte
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(planet)
		val = b.Get(k)
		return nil
	})
	str := string(val[:])
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
