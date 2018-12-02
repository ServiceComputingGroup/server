package graph_ql

import (
	"strconv"
	"time"

	"github.com/ServiceComputingGroup/simpleWebServer/database"
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
				exp, _ := strconv.ParseInt(claims.(jwt.MapClaims)["exp"].(string), 10, 64)
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
					if (pagenum-1)*10 >= len(all_data) || err != nil {
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
							result = append(result, all_data[i])
						}
						return result, nil
					}
				} else {
					var max_num int
					if len(all_data) < 10 {
						max_num = len(all_data)
					} else {
						max_num = 10
					}
					var result []string
					for i := 0; i < max_num; i++ {
						result = append(result, all_data[i])
					}
					return result, nil
				}
			} else {
				switch type_name {
				case "film":
					return database.GetFilm(index.(string)), nil
				case "people":
					return database.GetPerson(index.(string)), nil
				case "planet":
					return database.GetPlanet(index.(string)), nil
				case "species":
					return database.GetSpecies(index.(string)), nil
				case "starship":
					return database.GetStartship(index.(string)), nil
				case "vehicle":
					return database.GetVehicle(index.(string)), nil
				}
			}
			return "", nil
		},
	}
	return field
}
