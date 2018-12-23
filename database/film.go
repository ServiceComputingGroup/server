package database

import (
	"encoding/json"
	"fmt"

	"github.com/ServiceComputingGroup/simpleWebServer/entity"
)

func GetFilm(key string) string {
	key = "https://swapi.co/api/films/" + key + "/"
	var str string

	datas := GetFilms()
	for _, v := range datas {

		if v.Url == key {
			encoded, _ := json.MarshalIndent(v, "", "\t")
			str = string(encoded)

		}
	}
	return str
}
func GetFilms() []entity.Film {
	//执行查询语句
	rows, err := DB.Query("SELECT * from film")
	if err != nil {
		fmt.Println("查询出错了")
	}
	var datas []entity.Film
	//循环读取结果
	for rows.Next() {
		var data entity.Film
		var str string
		var id int
		err := rows.Scan(&id, &str)
		json.Unmarshal([]byte(str), &data)
		if err != nil {
			fmt.Println("rows fail")
		}
		//将user追加到users的这个数组中
		datas = append(datas, data)
	}
	return datas
}
