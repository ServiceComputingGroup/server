package database

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/ServiceComputingGroup/simpleWebServer/entity"
	"github.com/boltdb/bolt"
)

var (
	db       *bolt.DB
	userB    []byte
	people   []byte
	film     []byte
	planet   []byte
	species  []byte
	starship []byte
	vehicle  []byte
)

const dbname = "./data/module.db"

func init() {
	var err error
	db, err = bolt.Open(dbname, 0600, nil)
	//初始化bucket
	userB = []byte("user")
	people = []byte("people")
	film = []byte("film")
	planet = []byte("planet")
	species = []byte("species")
	starship = []byte("starship")
	vehicle = []byte("vehicle")
	if err != nil {
		fmt.Println("open err:", err)
		return
	}

	//AddInitData()
}
func AddInitData() {

	//创建bucket
	fmt.Println("正在创建bucket")

	db.Update(func(tx *bolt.Tx) error {

		tx.DeleteBucket(userB)
		tx.DeleteBucket(people)
		tx.DeleteBucket(film)
		tx.DeleteBucket(planet)
		tx.DeleteBucket(starship)
		tx.DeleteBucket(species)
		tx.DeleteBucket(vehicle)
		_, err := tx.CreateBucket(userB)
		if err != nil {
			fmt.Println("open err:", err)
		}
		_, err = tx.CreateBucket(people)
		if err != nil {
			fmt.Println("open err:", err)
		}

		_, err = tx.CreateBucket(film)
		if err != nil {
			fmt.Println("open err:", err)
		}
		_, err = tx.CreateBucket(planet)
		if err != nil {
			fmt.Println("open err:", err)
		}
		_, err = tx.CreateBucket(species)
		if err != nil {
			fmt.Println("open err:", err)
		}
		_, err = tx.CreateBucket(starship)
		if err != nil {
			fmt.Println("open err:", err)
		}
		_, err = tx.CreateBucket(vehicle)
		if err != nil {
			fmt.Println("open err:", err)
		}

		return err

	})
	fmt.Println("正在插入数据")
	db.Update(func(tx *bolt.Tx) error {
		fmt.Println("正在插入peoples")
		peoples := initPeoples()
		b := tx.Bucket(people)
		for _, v := range peoples { //range遍历，返回下标，和值
			//encoded, err := json.Marshal(v)
			id, _ := b.NextSequence()
			encoded, err := json.MarshalIndent(v, "", "\t")
			if err != nil {
				fmt.Println("open err:", err)
				return err
			}

			err = b.Put(itob(int(id)), (encoded))
			if err != nil {
				fmt.Println("open err:", err)
				return err
			}
		}
		return nil
	})
	db.Update(func(tx *bolt.Tx) error {
		fmt.Println("正在插入starships")
		starships := initStarships()
		b := tx.Bucket(starship)
		for _, v := range starships { //range遍历，返回下标，和值
			id, _ := b.NextSequence()
			encoded, err := json.MarshalIndent(v, "", "\t")
			if err != nil {
				fmt.Println("open err:", err)
				return err
			}

			err = b.Put(itob(int(id)), (encoded))
			if err != nil {
				fmt.Println("open err:", err)
				return err
			}
		}
		return nil
	})
	db.Update(func(tx *bolt.Tx) error {
		fmt.Println("正在插入planets")
		planets := initPlanets()
		b := tx.Bucket(planet)
		for _, v := range planets { //range遍历，返回下标，和值
			id, _ := b.NextSequence()
			encoded, err := json.MarshalIndent(v, "", "\t")
			if err != nil {
				fmt.Println("open err:", err)
				return err
			}

			err = b.Put(itob(int(id)), (encoded))
			if err != nil {
				fmt.Println("open err:", err)
				return err
			}
		}
		return nil
	})
	db.Update(func(tx *bolt.Tx) error {
		fmt.Println("正在插入films")
		films := initFilms()
		b := tx.Bucket(film)
		for _, v := range films { //range遍历，返回下标，和值
			id, _ := b.NextSequence()
			encoded, err := json.MarshalIndent(v, "", "\t")
			if err != nil {
				fmt.Println("open err:", err)
				return err
			}

			err = b.Put(itob(int(id)), (encoded))
			if err != nil {
				fmt.Println("open err:", err)
				return err
			}
		}
		return nil
	})
	db.Update(func(tx *bolt.Tx) error {
		fmt.Println("正在插入vehicles")
		vehicles := initVehicles()
		b := tx.Bucket(vehicle)
		for _, v := range vehicles { //range遍历，返回下标，和值
			id, _ := b.NextSequence()
			encoded, err := json.MarshalIndent(v, "", "\t")
			if err != nil {
				fmt.Println("open err:", err)
				return err
			}

			err = b.Put(itob(int(id)), (encoded))
			if err != nil {
				fmt.Println("open err:", err)
				return err
			}
		}
		return nil
	})
	db.Update(func(tx *bolt.Tx) error {
		fmt.Println("正在插入species")
		speciesAll := initSpecies()
		b := tx.Bucket(species)
		for _, v := range speciesAll { //range遍历，返回下标，和值
			id, _ := b.NextSequence()
			encoded, err := json.MarshalIndent(v, "", "\t")

			if err != nil {
				fmt.Println("open err:", err)
				return err
			}

			err = b.Put(itob(int(id)), (encoded))
			if err != nil {
				fmt.Println("open err:", err)
				return err
			}
		}
		return nil
	})
	fmt.Println("AddInitData完成")

}
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
func initPeoples() []entity.People {
	Path := "./data/datainit/people.json"
	var datas []entity.People

	jsonStr, err := ioutil.ReadFile(Path)
	if err != nil {
		panic(err)
	}

	//str := strings.Replace(string(jsonStr), "\n", "", 1)

	err = json.Unmarshal(jsonStr, &datas)
	if err != nil {
		panic(err)
	}

	return datas
}
func initFilms() []entity.Film {
	Path := "./data/datainit/films.json"
	var datas []entity.Film

	jsonStr, err := ioutil.ReadFile(Path)
	if err != nil {
		panic(err)
	}

	//str := strings.Replace(string(jsonStr), "\n", "", 1)

	err = json.Unmarshal(jsonStr, &datas)
	if err != nil {
		panic(err)
	}

	return datas
}
func initStarships() []entity.Starship {
	Path := "./data/datainit/starships.json"
	var datas []entity.Starship

	jsonStr, err := ioutil.ReadFile(Path)
	if err != nil {
		panic(err)
	}

	//str := strings.Replace(string(jsonStr), "\n", "", 1)

	err = json.Unmarshal(jsonStr, &datas)
	if err != nil {
		panic(err)
	}

	return datas
}
func initPlanets() []entity.Planet {
	Path := "./data/datainit/planets.json"
	var datas []entity.Planet

	jsonStr, err := ioutil.ReadFile(Path)
	if err != nil {
		panic(err)
	}

	//str := strings.Replace(string(jsonStr), "\n", "", 1)

	err = json.Unmarshal(jsonStr, &datas)
	if err != nil {
		panic(err)
	}

	return datas
}
func initVehicles() []entity.Vehicle {
	Path := "./data/datainit/vehicles.json"
	var datas []entity.Vehicle

	jsonStr, err := ioutil.ReadFile(Path)
	if err != nil {
		panic(err)
	}

	//str := strings.Replace(string(jsonStr), "\n", "", 1)

	err = json.Unmarshal(jsonStr, &datas)
	if err != nil {
		panic(err)
	}

	return datas
}
func initSpecies() []entity.Species {
	Path := "./data/datainit/species.json"
	var datas []entity.Species

	jsonStr, err := ioutil.ReadFile(Path)
	if err != nil {
		panic(err)
	}

	//str := strings.Replace(string(jsonStr), "\n", "", 1)

	err = json.Unmarshal(jsonStr, &datas)
	if err != nil {
		panic(err)
	}

	return datas
}
