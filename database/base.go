package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/ServiceComputingGroup/simpleWebServer/entity"
	"github.com/boltdb/bolt"
	_ "github.com/go-sql-driver/mysql"
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
	InitDB()
	//AddInitData()
}

const (
	userName = "docker"
	password = "123456"
	ip       = "localhost"
	port     = "3306"
	dbName   = "docker_mysql"
)

//Db数据库连接池
var DB *sql.DB

func InitDB() {
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")

	//打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
	DB, _ = sql.Open("mysql", path)
	//设置数据库最大连接数
	DB.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(10)
	//验证连接
	if err := DB.Ping(); err != nil {
		fmt.Println("opon database fail")
		return
	}

}

func AddInitData() {
	InitDB()
	/*mysql_film()
	mysql_people()
	mysql_starship()
	mysql_planet()
	mysql_vehicle()
	mysql_species()*/
}
func mysql_film() {
	f, _ := os.OpenFile("data/datamysql/film.txt", os.O_WRONLY|os.O_TRUNC, 0600)
	data := initFilms()
	size := len(data)
	for index, v := range data {
		encoded, _ := json.MarshalIndent(v, "", "\t")
		var new []byte
		for _, v := range encoded {
			if v == '\'' {
				new = append(new, '\\', '\'')
			} else {
				new = append(new, v)
			}
		}

		str := "(" + string(new) + "')"
		f.Write([]byte(str))
		if size != index+1 {
			f.Write([]byte(","))
		}

	}
	f.Write([]byte(";"))
	f.Close()
}
func mysql_people() {
	f, _ := os.OpenFile("data/datamysql/people.txt", os.O_WRONLY|os.O_TRUNC, 0600)
	data := initPeoples()
	size := len(data)
	for index, v := range data {
		encoded, _ := json.MarshalIndent(v, "", "\t")
		var new []byte
		for _, v := range encoded {
			if v == '\'' {
				new = append(new, '\\', '\'')
			} else {
				new = append(new, v)
			}
		}
		str := "(" + string(new) + "')"
		f.Write([]byte(str))
		if size != index+1 {
			f.Write([]byte(","))
		}

	}
	f.Write([]byte(";"))
	f.Close()
}
func mysql_starship() {
	f, _ := os.OpenFile("data/datamysql/starship.txt", os.O_WRONLY|os.O_TRUNC, 0600)
	data := initStarships()
	size := len(data)
	for index, v := range data {
		encoded, _ := json.MarshalIndent(v, "", "\t")
		var new []byte
		for _, v := range encoded {
			if v == '\'' {
				new = append(new, '\\', '\'')
			} else {
				new = append(new, v)
			}
		}
		str := "(" + string(new) + "')"
		f.Write([]byte(str))
		if size != index+1 {
			f.Write([]byte(","))
		}

	}
	f.Write([]byte(";"))
	f.Close()
}
func mysql_planet() {
	f, _ := os.OpenFile("data/datamysql/planet.txt", os.O_WRONLY|os.O_TRUNC, 0600)
	data := initPlanets()
	size := len(data)
	for index, v := range data {
		encoded, _ := json.MarshalIndent(v, "", "\t")
		var new []byte
		for _, v := range encoded {
			if v == '\'' {
				new = append(new, '\\', '\'')
			} else {
				new = append(new, v)
			}
		}
		str := "(" + string(new) + "')"
		f.Write([]byte(str))
		if size != index+1 {
			f.Write([]byte(","))
		}

	}
	f.Write([]byte(";"))
	f.Close()
}
func mysql_vehicle() {
	f, _ := os.OpenFile("data/datamysql/vehicle.txt", os.O_WRONLY|os.O_TRUNC, 0600)
	data := initVehicles()
	size := len(data)
	for index, v := range data {
		encoded, _ := json.MarshalIndent(v, "", "\t")
		var new []byte
		for _, v := range encoded {
			if v == '\'' {
				new = append(new, '\\', '\'')
			} else {
				new = append(new, v)
			}
		}
		str := "(" + string(new) + "')"
		f.Write([]byte(str))
		if size != index+1 {
			f.Write([]byte(","))
		}

	}
	f.Write([]byte(";"))
	f.Close()
}
func mysql_species() {
	f, _ := os.OpenFile("data/datamysql/species.txt", os.O_WRONLY|os.O_TRUNC, 0600)
	data := initSpecies()
	size := len(data)
	for index, v := range data {
		encoded, _ := json.MarshalIndent(v, "", "\t")
		var new []byte
		for _, v := range encoded {
			if v == '\'' {
				new = append(new, '\\', '\'')
			} else {
				new = append(new, v)
			}
		}
		str := "(" + string(new) + "')"
		f.Write([]byte(str))
		if size != index+1 {
			f.Write([]byte(","))
		}

	}
	f.Write([]byte(";"))
	f.Close()
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
