package database

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"

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

		fmt.Println("正在插入数据")
		peoples := initPeoples()
		films := initFilms()
		starships := initStarships()
		planets := initPlanets()
		vehicles := initVehicles()
		speciesAll := initSpecies()

		fmt.Println("正在插入peoples")
		b := tx.Bucket(people)
		for i, v := range peoples { //range遍历，返回下标，和值
			encoded, err := json.Marshal(v)
			if err != nil {
				fmt.Println("open err:", err)
				return err
			}

			err = b.Put([]byte(strconv.Itoa(i+1)), []byte(encoded))
			if err != nil {
				fmt.Println("open err:", err)
				return err
			}
		}

		fmt.Println("正在插入films")
		b = tx.Bucket(film)
		for i, v := range films { //range遍历，返回下标，和值
			encoded, err := json.Marshal(v)
			if err != nil {
				fmt.Println("open err:", err)
				return err
			}

			fmt.Println([]byte(strconv.Itoa(i + 1)))

			err = b.Put([]byte(strconv.Itoa(i+1)), []byte(encoded))
			val := b.Get([]byte(strconv.Itoa(i + 1)))
			fmt.Println(val)
			if err != nil {
				fmt.Println("open err:", err)
				return err
			}
		}

		fmt.Println("正在插入starships")
		b = tx.Bucket(starship)
		for i, v := range starships { //range遍历，返回下标，和值
			encoded, err := json.Marshal(v)
			if err != nil {
				fmt.Println("open err:", err)
				return err
			}
			err = b.Put([]byte(strconv.Itoa(i+1)), []byte(encoded))
			if err != nil {
				fmt.Println("open err:", err)
				return err
			}
		}

		fmt.Println("正在插入planets")
		b = tx.Bucket(planet)
		for i, v := range planets { //range遍历，返回下标，和值
			encoded, err := json.Marshal(v)
			if err != nil {
				fmt.Println("open err:", err)
				return err
			}
			err = b.Put([]byte(strconv.Itoa(i+1)), []byte(encoded))
			if err != nil {
				fmt.Println("open err:", err)
				return err
			}
		}

		fmt.Println("正在插入vehicles")
		b = tx.Bucket(vehicle)
		for i, v := range vehicles { //range遍历，返回下标，和值
			encoded, err := json.Marshal(v)
			if err != nil {
				fmt.Println("open err:", err)
				return err
			}
			err = b.Put([]byte(strconv.Itoa(i+1)), []byte(encoded))
			if err != nil {
				fmt.Println("open err:", err)
				return err
			}
		}

		fmt.Println("正在插入species")
		b = tx.Bucket(species)
		for i, v := range speciesAll { //range遍历，返回下标，和值
			encoded, err := json.Marshal(v)
			if err != nil {
				fmt.Println("open err:", err)
				return err
			}
			err = b.Put([]byte(strconv.Itoa(i+1)), []byte(encoded))
			if err != nil {
				fmt.Println("open err:", err)
				return err
			}
		}

		return err

	})
	fmt.Println("AddInitData完成")

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
