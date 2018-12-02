package database

import (
	"github.com/boltdb/bolt"
)

func GetVehicle(key string) string {
	k := []byte(key)
	var val []byte
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(vehicle)
		val = b.Get(k)
		return nil
	})
	str := string(val[:])
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
