package database

import (
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

	db, err := bolt.Open(dbname, 0600, nil)
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

	db.Close()

	//AddInitData()
}
func AddInitData() {

	db, _ := bolt.Open(dbname, 0600, nil)
	//创建bucket
	fmt.Println("正在创建bucket")
	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket(userB)
		if err != nil {
			fmt.Println("open err:", err)
			return err
		}
		_, err = tx.CreateBucket(people)
		if err != nil {
			fmt.Println("open err:", err)
			return err
		}

		_, err = tx.CreateBucket(film)
		if err != nil {
			fmt.Println("open err:", err)
			return err
		}
		_, err = tx.CreateBucket(planet)
		if err != nil {
			fmt.Println("open err:", err)
			return err
		}
		_, err = tx.CreateBucket(species)
		if err != nil {
			fmt.Println("open err:", err)
			return err
		}
		_, err = tx.CreateBucket(starship)
		if err != nil {
			fmt.Println("open err:", err)
			return err
		}
		_, err = tx.CreateBucket(vehicle)
		if err != nil {
			fmt.Println("open err:", err)
			return err
		}
		b := tx.Bucket(people)
		fmt.Println("正在插入数据")
		peoples := initPeoples()
		films := initFilms()
		starships := initStarships()
		planets := initPlanets()
		vehicles := initVehicles()
		species := initSpecies()

		fmt.Println("正在插入peoples")
		for i, v := range peoples { //range遍历，返回下标，和值
			encoded, err := json.Marshal(v)
			if err != nil {
				fmt.Println("open err:", err)
				return err
			}
			err = b.Put([]byte(string(i+1)), []byte(encoded))
			if err != nil {
				fmt.Println("open err:", err)
				return err
			}
		}

		fmt.Println("正在插入films")
		for i, v := range films { //range遍历，返回下标，和值
			encoded, err := json.Marshal(v)
			if err != nil {
				fmt.Println("open err:", err)
				return err
			}
			err = b.Put([]byte(string(i+1)), []byte(encoded))
			if err != nil {
				fmt.Println("open err:", err)
				return err
			}
		}

		fmt.Println("正在插入starships")
		for i, v := range starships { //range遍历，返回下标，和值
			encoded, err := json.Marshal(v)
			if err != nil {
				fmt.Println("open err:", err)
				return err
			}
			err = b.Put([]byte(string(i+1)), []byte(encoded))
			if err != nil {
				fmt.Println("open err:", err)
				return err
			}
		}

		fmt.Println("正在插入planets")
		for i, v := range planets { //range遍历，返回下标，和值
			encoded, err := json.Marshal(v)
			if err != nil {
				fmt.Println("open err:", err)
				return err
			}
			err = b.Put([]byte(string(i+1)), []byte(encoded))
			if err != nil {
				fmt.Println("open err:", err)
				return err
			}
		}

		fmt.Println("正在插入vehicles")
		for i, v := range vehicles { //range遍历，返回下标，和值
			encoded, err := json.Marshal(v)
			if err != nil {
				fmt.Println("open err:", err)
				return err
			}
			err = b.Put([]byte(string(i+1)), []byte(encoded))
			if err != nil {
				fmt.Println("open err:", err)
				return err
			}
		}

		fmt.Println("正在插入species")
		for i, v := range species { //range遍历，返回下标，和值
			encoded, err := json.Marshal(v)
			if err != nil {
				fmt.Println("open err:", err)
				return err
			}
			err = b.Put([]byte(string(i+1)), []byte(encoded))
			if err != nil {
				fmt.Println("open err:", err)
				return err
			}
		}

		return err

	})

}
func initPeoples() []entity.People {
	Path := "./data/datainit/people.json"
	var datas []entity.People

	josnStr, err := ioutil.ReadFile(Path)
	if err != nil {
		panic(err)
	}

	//str := strings.Replace(string(josnStr), "\n", "", 1)

	err = json.Unmarshal(josnStr, &datas)
	if err != nil {
		panic(err)
	}

	return datas
}
func initFilms() []entity.Film {
	Path := "./data/datainit/films.json"
	var datas []entity.Film

	josnStr, err := ioutil.ReadFile(Path)
	if err != nil {
		panic(err)
	}

	//str := strings.Replace(string(josnStr), "\n", "", 1)

	err = json.Unmarshal(josnStr, &datas)
	if err != nil {
		panic(err)
	}

	return datas
}
func initStarships() []entity.Starship {
	Path := "./data/datainit/starships.json"
	var datas []entity.Starship

	josnStr, err := ioutil.ReadFile(Path)
	if err != nil {
		panic(err)
	}

	//str := strings.Replace(string(josnStr), "\n", "", 1)

	err = json.Unmarshal(josnStr, &datas)
	if err != nil {
		panic(err)
	}

	return datas
}
func initPlanets() []entity.Planet {
	Path := "./data/datainit/planets.json"
	var datas []entity.Planet

	josnStr, err := ioutil.ReadFile(Path)
	if err != nil {
		panic(err)
	}

	//str := strings.Replace(string(josnStr), "\n", "", 1)

	err = json.Unmarshal(josnStr, &datas)
	if err != nil {
		panic(err)
	}

	return datas
}
func initVehicles() []entity.Vehicle {
	Path := "./data/datainit/vehicles.json"
	var datas []entity.Vehicle

	josnStr, err := ioutil.ReadFile(Path)
	if err != nil {
		panic(err)
	}

	//str := strings.Replace(string(josnStr), "\n", "", 1)

	err = json.Unmarshal(josnStr, &datas)
	if err != nil {
		panic(err)
	}

	return datas
}
func initSpecies() []entity.Species {
	Path := "./data/datainit/species.json"
	var datas []entity.Species

	josnStr, err := ioutil.ReadFile(Path)
	if err != nil {
		panic(err)
	}

	//str := strings.Replace(string(josnStr), "\n", "", 1)

	err = json.Unmarshal(josnStr, &datas)
	if err != nil {
		panic(err)
	}

	return datas
}
