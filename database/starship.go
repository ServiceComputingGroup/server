package database

import (
	"github.com/boltdb/bolt"
)

func GetStartship(key string) string {
	key = "https://swapi.co/api/starships/" + key + "/"
	k := []byte(key)
	var val []byte
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(starship)
		val = b.Get(k)
		return nil
	})
	str := string(val[:])
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
