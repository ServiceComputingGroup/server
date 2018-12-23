package database

import (
	"encoding/json"
	"fmt"

	"github.com/ServiceComputingGroup/simpleWebServer/entity"
)

func GetSpecies(key string) string {
	key = "https://swapi.co/api/species/" + key + "/"
	var str string

	datas := GetAllSpecies()
	for _, v := range datas {

		if v.Url == key {
			encoded, _ := json.MarshalIndent(v, "", "\t")
			str = string(encoded)

		}
	}
	return str
}
func GetAllSpecies() []entity.Species {
	//执行查询语句
	rows, err := DB.Query("SELECT * from species")
	if err != nil {
		fmt.Println("查询出错了")
	}
	var datas []entity.Species
	//循环读取结果
	for rows.Next() {
		var data entity.Species
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
