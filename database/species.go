package database

import (
	"github.com/boltdb/bolt"
)

func GetSpecies(key string) string {
	key = "https://swapi.co/api/species/" + key + "/"
	k := []byte(key)
	var val []byte
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(species)
		val = b.Get(k)
		return nil
	})
	str := string(val[:])
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
