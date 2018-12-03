package database

import (
	"fmt"

	"github.com/boltdb/bolt"
)

func GetFilm(key string) string {
	fmt.Println("GetFilm")
	fmt.Println(key)
	k := []byte(key)
	fmt.Println(k)
	var val []byte
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(film)
		val = b.Get(k)
		fmt.Println(val)
		return nil
	})
	str := string(val)
	return str
}

func GetFilms() []string {

	var result []string

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(film)
		cur := b.Cursor()

		for k, v := cur.First(); k != nil; k, v = cur.Next() {

			result = append(result, string(v[:]))
		}
		return nil
	})

	return result
}
