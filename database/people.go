package database

import (
	"github.com/boltdb/bolt"
)

func GetPerson(key string) string {
	key = "https://swapi.co/api/people/" + key + "/"
	k := []byte(key)
	var val []byte
	db.View(func(tx *bolt.Tx) error {

		b := tx.Bucket(people)
		val = b.Get(k)

		return nil
	})

	str := string(val[:])
	return str
}

func GetPeople() []string {

	var result []string

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(people)
		cur := b.Cursor()

		for k, v := cur.First(); k != nil; k, v = cur.Next() {

			result = append(result, string(v[:]))
		}
		return nil
	})

	return result
}
