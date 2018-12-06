package graph_ql

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/ServiceComputingGroup/simpleWebServer/database"
	"github.com/ServiceComputingGroup/simpleWebServer/entity"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/graphql-go/graphql"
)

func Queryinfo() *graphql.Field {
	var field = &graphql.Field{
		Type:        graphql.String,
		Description: "查询信息",
		Args: graphql.FieldConfigArgument{
			"type": &graphql.ArgumentConfig{
				Description: "查询类型",
				Type:        graphql.NewNonNull(graphql.String),
			},
			"page": &graphql.ArgumentConfig{
				Description: "查询页数",
				Type:        graphql.String,
			},
			"index": &graphql.ArgumentConfig{
				Description: "查询编号",
				Type:        graphql.String,
			},
			"token": &graphql.ArgumentConfig{
				Description: "查询权限",
				Type:        graphql.NewNonNull(graphql.String),
			},
		},

		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			type_name := p.Args["type"].(string)
			page := p.Args["page"]
			index := p.Args["index"]

			tokenstring := p.Args["token"].(string)
			key := "UserToken"
			isPermitted := false
			claims, ok := parseToken(tokenstring, key)
			if ok {
				exp := (int64)(claims.(jwt.MapClaims)["exp"].(float64))
				if time.Now().UTC().Unix() > exp {
					isPermitted = false
				} else {
					isPermitted = true
				}
			} else {
				isPermitted = false
			}

			if !isPermitted {
				return "Please log in / register first", nil
			}

			if index == nil {
				if page != nil {
					pagenum, err := strconv.Atoi(page.(string))
					switch type_name {
					case "films":
						all_data := database.GetFilms()

						if (pagenum-1)*10 >= len(all_data) || err != nil || pagenum <= 0 {
							return "404 not found", nil
						} else {
							//get count, next_page, previous_page
							count := len(all_data)
							var next string
							var previous string
							if pagenum*10 >= len(all_data) {
								next = "null"
							} else {
								next = "https://swapi.co/api/" + type_name + "/?page=" + strconv.Itoa(pagenum+1)
							}
							if pagenum == 1 {
								previous = "null"
							} else {
								previous = "https://swapi.co/api/" + type_name + "/?page=" + strconv.Itoa(pagenum-1)
							}
							var max_num int
							if pagenum*10 > len(all_data) {
								max_num = len(all_data)
							} else {
								max_num = pagenum * 10
							}
							var result []entity.Film
							for i := (pagenum - 1) * 10; i < max_num; i++ {
								result = append(result, all_data[i])
							}
							result_information := &entity.QueryFilm{
								Count:    strconv.Itoa(count),
								Next:     next,
								Previous: previous,
								Result:   result,
							}
							json_string, _ := json.MarshalIndent(result_information, "", "\t")
							return string(json_string), nil
						}
					case "people":
						all_data := database.GetPeople()

						if (pagenum-1)*10 >= len(all_data) || err != nil || pagenum <= 0 {
							return "404 not found", nil
						} else {
							//get count, next_page, previous_page
							count := len(all_data)
							var next string
							var previous string
							if pagenum*10 >= len(all_data) {
								next = "null"
							} else {
								next = "https://swapi.co/api/" + type_name + "/?page=" + strconv.Itoa(pagenum+1)
							}
							if pagenum == 1 {
								previous = "null"
							} else {
								previous = "https://swapi.co/api/" + type_name + "/?page=" + strconv.Itoa(pagenum-1)
							}
							var max_num int
							if pagenum*10 > len(all_data) {
								max_num = len(all_data)
							} else {
								max_num = pagenum * 10
							}
							var result []entity.People
							for i := (pagenum - 1) * 10; i < max_num; i++ {
								result = append(result, all_data[i])
							}
							result_information := &entity.QueryPeople{
								Count:    strconv.Itoa(count),
								Next:     next,
								Previous: previous,
								Result:   result,
							}
							json_string, _ := json.MarshalIndent(result_information, "", "\t")
							return string(json_string), nil
						}
					case "planets":
						all_data := database.GetPlanets()

						if (pagenum-1)*10 >= len(all_data) || err != nil || pagenum <= 0 {
							return "404 not found", nil
						} else {
							//get count, next_page, previous_page
							count := len(all_data)
							var next string
							var previous string
							if pagenum*10 >= len(all_data) {
								next = "null"
							} else {
								next = "https://swapi.co/api/" + type_name + "/?page=" + strconv.Itoa(pagenum+1)
							}
							if pagenum == 1 {
								previous = "null"
							} else {
								previous = "https://swapi.co/api/" + type_name + "/?page=" + strconv.Itoa(pagenum-1)
							}
							var max_num int
							if pagenum*10 > len(all_data) {
								max_num = len(all_data)
							} else {
								max_num = pagenum * 10
							}
							var result []entity.Planet
							for i := (pagenum - 1) * 10; i < max_num; i++ {
								result = append(result, all_data[i])
							}
							result_information := &entity.QueryPlanet{
								Count:    strconv.Itoa(count),
								Next:     next,
								Previous: previous,
								Result:   result,
							}
							json_string, _ := json.MarshalIndent(result_information, "", "\t")
							return string(json_string), nil
						}
					case "species":
						all_data := database.GetAllSpecies()

						if (pagenum-1)*10 >= len(all_data) || err != nil || pagenum <= 0 {
							return "404 not found", nil
						} else {
							//get count, next_page, previous_page
							count := len(all_data)
							var next string
							var previous string
							if pagenum*10 >= len(all_data) {
								next = "null"
							} else {
								next = "https://swapi.co/api/" + type_name + "/?page=" + strconv.Itoa(pagenum+1)
							}
							if pagenum == 1 {
								previous = "null"
							} else {
								previous = "https://swapi.co/api/" + type_name + "/?page=" + strconv.Itoa(pagenum-1)
							}
							var max_num int
							if pagenum*10 > len(all_data) {
								max_num = len(all_data)
							} else {
								max_num = pagenum * 10
							}
							var result []entity.Species
							for i := (pagenum - 1) * 10; i < max_num; i++ {
								result = append(result, all_data[i])
							}
							result_information := &entity.QuerySpecies{
								Count:    strconv.Itoa(count),
								Next:     next,
								Previous: previous,
								Result:   result,
							}
							json_string, _ := json.MarshalIndent(result_information, "", "\t")
							return string(json_string), nil
						}
					case "starships":
						all_data := database.GetStarships()

						if (pagenum-1)*10 >= len(all_data) || err != nil || pagenum <= 0 {
							return "404 not found", nil
						} else {
							//get count, next_page, previous_page
							count := len(all_data)
							var next string
							var previous string
							if pagenum*10 >= len(all_data) {
								next = "null"
							} else {
								next = "https://swapi.co/api/" + type_name + "/?page=" + strconv.Itoa(pagenum+1)
							}
							if pagenum == 1 {
								previous = "null"
							} else {
								previous = "https://swapi.co/api/" + type_name + "/?page=" + strconv.Itoa(pagenum-1)
							}
							var max_num int
							if pagenum*10 > len(all_data) {
								max_num = len(all_data)
							} else {
								max_num = pagenum * 10
							}
							var result []entity.Starship
							for i := (pagenum - 1) * 10; i < max_num; i++ {
								result = append(result, all_data[i])
							}
							result_information := &entity.QueryStarship{
								Count:    strconv.Itoa(count),
								Next:     next,
								Previous: previous,
								Result:   result,
							}
							json_string, _ := json.MarshalIndent(result_information, "", "\t")
							return string(json_string), nil
						}
					case "vehicles":
						all_data := database.GetVehicles()

						if (pagenum-1)*10 >= len(all_data) || err != nil || pagenum <= 0 {
							return "404 not found", nil
						} else {
							//get count, next_page, previous_page
							count := len(all_data)
							var next string
							var previous string
							if pagenum*10 >= len(all_data) {
								next = "null"
							} else {
								next = "https://swapi.co/api/" + type_name + "/?page=" + strconv.Itoa(pagenum+1)
							}
							if pagenum == 1 {
								previous = "null"
							} else {
								previous = "https://swapi.co/api/" + type_name + "/?page=" + strconv.Itoa(pagenum-1)
							}
							var max_num int
							if pagenum*10 > len(all_data) {
								max_num = len(all_data)
							} else {
								max_num = pagenum * 10
							}
							var result []entity.Vehicle
							for i := (pagenum - 1) * 10; i < max_num; i++ {
								result = append(result, all_data[i])
							}
							result_information := &entity.QueryVehicle{
								Count:    strconv.Itoa(count),
								Next:     next,
								Previous: previous,
								Result:   result,
							}
							json_string, _ := json.MarshalIndent(result_information, "", "\t")
							return string(json_string), nil
						}
					}
				} else {
					switch type_name {
					case "films":
						var result []entity.Film
						var max_num int
						all_data := database.GetFilms()
						if 10 > len(all_data) {
							max_num = len(all_data)
						} else {
							max_num = 10
						}
						for i := 0; i < max_num; i++ {
							result = append(result, all_data[i])
						}
						count := len(all_data)
						next := "https://swapi.co/api/" + type_name + "/?page=2"
						previous := "null"
						result_information := &entity.QueryFilm{
							Count:    strconv.Itoa(count),
							Next:     next,
							Previous: previous,
							Result:   result,
						}
						json_string, _ := json.MarshalIndent(result_information, "", "\t")
						fmt.Println(string(json_string))
						return string(json_string), nil
					case "people":
						var result []entity.People
						var max_num int
						all_data := database.GetPeople()
						if 10 > len(all_data) {
							max_num = len(all_data)
						} else {
							max_num = 10
						}
						for i := 0; i < max_num; i++ {
							result = append(result, all_data[i])
						}
						count := len(all_data)
						next := "https://swapi.co/api/" + type_name + "/?page=2"
						previous := "null"
						result_information := &entity.QueryPeople{
							Count:    strconv.Itoa(count),
							Next:     next,
							Previous: previous,
							Result:   result,
						}
						json_string, _ := json.MarshalIndent(result_information, "", "\t")
						fmt.Println(string(json_string))
						return string(json_string), nil
					case "planets":
						var result []entity.Planet
						var max_num int
						all_data := database.GetPlanets()
						if 10 > len(all_data) {
							max_num = len(all_data)
						} else {
							max_num = 10
						}
						for i := 0; i < max_num; i++ {
							result = append(result, all_data[i])
						}
						count := len(all_data)
						next := "https://swapi.co/api/" + type_name + "/?page=2"
						previous := "null"
						result_information := &entity.QueryPlanet{
							Count:    strconv.Itoa(count),
							Next:     next,
							Previous: previous,
							Result:   result,
						}
						json_string, _ := json.MarshalIndent(result_information, "", "\t")
						fmt.Println(string(json_string))
						return string(json_string), nil
					case "species":
						var result []entity.Species
						var max_num int
						all_data := database.GetAllSpecies()
						if 10 > len(all_data) {
							max_num = len(all_data)
						} else {
							max_num = 10
						}
						for i := 0; i < max_num; i++ {
							result = append(result, all_data[i])
						}
						count := len(all_data)
						next := "https://swapi.co/api/" + type_name + "/?page=2"
						previous := "null"
						result_information := &entity.QuerySpecies{
							Count:    strconv.Itoa(count),
							Next:     next,
							Previous: previous,
							Result:   result,
						}
						json_string, _ := json.MarshalIndent(result_information, "", "\t")
						fmt.Println(string(json_string))
						return string(json_string), nil
					case "starships":
						var result []entity.Starship
						var max_num int
						all_data := database.GetStarships()
						if 10 > len(all_data) {
							max_num = len(all_data)
						} else {
							max_num = 10
						}
						for i := 0; i < max_num; i++ {
							result = append(result, all_data[i])
						}
						count := len(all_data)
						next := "https://swapi.co/api/" + type_name + "/?page=2"
						previous := "null"
						result_information := &entity.QueryStarship{
							Count:    strconv.Itoa(count),
							Next:     next,
							Previous: previous,
							Result:   result,
						}
						json_string, _ := json.MarshalIndent(result_information, "", "\t")
						fmt.Println(string(json_string))
						return string(json_string), nil
					case "vehicles":
						var result []entity.Vehicle
						var max_num int
						all_data := database.GetVehicles()
						if 10 > len(all_data) {
							max_num = len(all_data)
						} else {
							max_num = 10
						}
						for i := 0; i < max_num; i++ {
							result = append(result, all_data[i])
						}
						count := len(all_data)
						next := "https://swapi.co/api/" + type_name + "/?page=2"
						previous := "null"
						result_information := &entity.QueryVehicle{
							Count:    strconv.Itoa(count),
							Next:     next,
							Previous: previous,
							Result:   result,
						}
						json_string, _ := json.MarshalIndent(result_information, "", "\t")
						fmt.Println(string(json_string))
						return string(json_string), nil
					}
				}
			} else {
				switch type_name {
				case "films":
					var someFilm entity.Film
					if err := json.Unmarshal([]byte(database.GetFilm(index.(string))), &someFilm); err == nil {
						json_string, err := json.MarshalIndent(someFilm, "", "   ")
						if err != nil {
							fmt.Println("json err:", err)
						}
						fmt.Printf(string(json_string))
						return string(json_string), nil
					} else {
						return "404 Not Found", nil
					}
				case "people":
					var somePeople entity.People
					if err := json.Unmarshal([]byte(database.GetPerson(index.(string))), &somePeople); err == nil {
						json_string, err := json.MarshalIndent(somePeople, "", "   ")
						if err != nil {
							fmt.Println("json err:", err)
						}
						fmt.Printf(string(json_string))
						return string(json_string), nil
					} else {
						return "404 Not Found", nil
					}
				case "planets":
					var somePlanet entity.Planet
					if err := json.Unmarshal([]byte(database.GetPlanet(index.(string))), &somePlanet); err == nil {
						json_string, err := json.MarshalIndent(somePlanet, "", "   ")
						if err != nil {
							fmt.Println("json err:", err)
						}
						fmt.Printf(string(json_string))
						return string(json_string), nil
					} else {
						return "404 Not Found", nil
					}
				case "species":
					var someSpecies entity.Species
					if err := json.Unmarshal([]byte(database.GetSpecies(index.(string))), &someSpecies); err == nil {
						json_string, err := json.MarshalIndent(someSpecies, "", "   ")
						if err != nil {
							fmt.Println("json err:", err)
						}
						fmt.Printf(string(json_string))
						return string(json_string), nil
					} else {
						return "404 Not Found", nil
					}
				case "starships":
					var someStarship entity.Starship
					if err := json.Unmarshal([]byte(database.GetStarship(index.(string))), &someStarship); err == nil {
						json_string, err := json.MarshalIndent(someStarship, "", "   ")
						if err != nil {
							fmt.Println("json err:", err)
						}
						fmt.Printf(string(json_string))
						return string(json_string), nil
					} else {
						return "404 Not Found", nil
					}
				case "vehicles":
					var someVehicle entity.Vehicle
					if err := json.Unmarshal([]byte(database.GetVehicle(index.(string))), &someVehicle); err == nil {
						json_string, err := json.MarshalIndent(someVehicle, "", "   ")
						if err != nil {
							fmt.Println("json err:", err)
						}
						fmt.Printf(string(json_string))
						return string(json_string), nil
					} else {
						return "404 Not Found", nil
					}
				}
			}
			return "", nil
		},
	}
	return field
}
