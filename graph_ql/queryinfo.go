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
				var all_data []string
				switch type_name {
				case "film":
					all_data = database.GetFilms()
				case "people":
					all_data = database.GetPeople()
				case "planet":
					all_data = database.GetPlanets()
				case "species":
					all_data = database.GetAllSpecies()
				case "starship":
					all_data = database.GetStartships()
				case "vehicle":
					all_data = database.GetVehicles()
				}
				if page != nil {
					pagenum, err := strconv.Atoi(page.(string))
					if (pagenum-1)*10 >= len(all_data) || err != nil || page <= 0{
						return "404 not found", nil
					} else {
						var max_num int
						if pagenum*10 > len(all_data) {
							max_num = len(all_data)
						} else {
							max_num = pagenum * 10
						}
						var result []string
						for i := (pagenum - 1) * 10; i < max_num; i++ {
							var json_result string
							switch type_name {
							case "film":
								var someFilm entity.Film
								if err := json.Unmarshal([]byte(all_data[i]), &someFilm); err == nil {
									json_string, err := json.MarshalIndent(someFilm, "", "   ")
									if err != nil {
										fmt.Println("json err:", err)
									}
									json_result = string(json_string)
								} else {
									json_result = "404 Not Found"
								}
							case "people":
								var somePeople entity.People
								if err := json.Unmarshal([]byte(all_data[i]), &somePeople); err == nil {
									json_string, err := json.MarshalIndent(somePeople, "", "   ")
									if err != nil {
										fmt.Println("json err:", err)
									}
									json_result = string(json_string)
								} else {
									json_result = "404 Not Found"
								}
							case "planet":
								var somePlanet entity.Planet
								if err := json.Unmarshal([]byte(all_data[i]), &somePlanet); err == nil {
									json_string, err := json.MarshalIndent(somePlanet, "", "   ")
									if err != nil {
										fmt.Println("json err:", err)
									}
									json_result = string(json_string)
								} else {
									json_result = "404 Not Found"
								}
							case "species":
								var someSpecies entity.Species
								if err := json.Unmarshal([]byte(all_data[i]), &someSpecies); err == nil {
									json_string, err := json.MarshalIndent(someSpecies, "", "   ")
									if err != nil {
										fmt.Println("json err:", err)
									}
									json_result = string(json_string)
								} else {
									json_result = "404 Not Found"
								}
							case "starship":
								var someStarship entity.Starship
								if err := json.Unmarshal([]byte(all_data[i]), &someStarship); err == nil {
									json_string, err := json.MarshalIndent(someStarship, "", "   ")
									if err != nil {
										fmt.Println("json err:", err)
									}
									json_result = string(json_string)
								} else {
									json_result = "404 Not Found"
								}
							case "vehicle":
								var someVehicle entity.Vehicle
								if err := json.Unmarshal([]byte(all_data[i]), &someVehicle); err == nil {
									json_string, err := json.MarshalIndent(someVehicle, "", "   ")
									if err != nil {
										fmt.Println("json err:", err)
									}
									json_result = string(json_string)
								} else {
									json_result = "404 Not Found"
								}
							}
							result = append(result, json_result+"\n")
						}
						fmt.Println(result)
						return result, nil
					}
				} else {
					var result []string
					for i := 0; i < 10; i++ {
						var json_result string
						switch type_name {
						case "film":
							var someFilm entity.Film
							if err := json.Unmarshal([]byte(all_data[i]), &someFilm); err == nil {
								json_string, err := json.MarshalIndent(someFilm, "", "   ")
								if err != nil {
									fmt.Println("json err:", err)
								}
								json_result = string(json_string)
							} else {
								json_result = "404 Not Found"
							}
						case "people":
							var somePeople entity.People
							if err := json.Unmarshal([]byte(all_data[i]), &somePeople); err == nil {
								json_string, err := json.MarshalIndent(somePeople, "", "   ")
								if err != nil {
									fmt.Println("json err:", err)
								}
								json_result = string(json_string)
							} else {
								json_result = "404 Not Found"
							}
						case "planet":
							var somePlanet entity.Planet
							if err := json.Unmarshal([]byte(all_data[i]), &somePlanet); err == nil {
								json_string, err := json.MarshalIndent(somePlanet, "", "   ")
								if err != nil {
									fmt.Println("json err:", err)
								}
								json_result = string(json_string)
							} else {
								json_result = "404 Not Found"
							}
						case "species":
							var someSpecies entity.Species
							if err := json.Unmarshal([]byte(all_data[i]), &someSpecies); err == nil {
								json_string, err := json.MarshalIndent(someSpecies, "", "   ")
								if err != nil {
									fmt.Println("json err:", err)
								}
								json_result = string(json_string)
							} else {
								json_result = "404 Not Found"
							}
						case "starship":
							var someStarship entity.Starship
							if err := json.Unmarshal([]byte(all_data[i]), &someStarship); err == nil {
								json_string, err := json.MarshalIndent(someStarship, "", "   ")
								if err != nil {
									fmt.Println("json err:", err)
								}
								json_result = string(json_string)
							} else {
								json_result = "404 Not Found"
							}
						case "vehicle":
							var someVehicle entity.Vehicle
							if err := json.Unmarshal([]byte(all_data[i]), &someVehicle); err == nil {
								json_string, err := json.MarshalIndent(someVehicle, "", "   ")
								if err != nil {
									fmt.Println("json err:", err)
								}
								json_result = string(json_string)
							} else {
								json_result = "404 Not Found"
							}
						}
						result = append(result, json_result+"\n")
					}
					fmt.Println(result)
					return result, nil
				}
			} else {
				switch type_name {
				case "film":
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
				case "planet":
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
				case "starship":
					var someStarship entity.Starship
					if err := json.Unmarshal([]byte(database.GetStartship(index.(string))), &someStarship); err == nil {
						json_string, err := json.MarshalIndent(someStarship, "", "   ")
						if err != nil {
							fmt.Println("json err:", err)
						}
						fmt.Printf(string(json_string))
						return string(json_string), nil
					} else {
						return "404 Not Found", nil
					}
				case "vehicle":
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
