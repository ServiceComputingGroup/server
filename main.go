package main

import (
	"encoding/json"
	"fmt"

	"github.com/ServiceComputingGroup/simpleWebServer/database"
	"github.com/ServiceComputingGroup/simpleWebServer/entity"
)

/*"fmt"
"net/http"

"github.com/ServiceComputingGroup/simpleWebServer/graph_ql"

"github.com/rs/cors"*/
func main() {
	fmt.Println([]byte("\n"))
	fmt.Println("h'h")
	database.AddInitData()
	fmt.Println("User")
	var user entity.User
	user.Password = "4"
	user.UserName = "45"
	user.Email = "456"
	user.Phone = "456"

	fmt.Println("InsertUser")
	database.InsertUser(&user)
	database.DeleteUser(user.UserName)

	fmt.Println("GetUser")

	var user2 *entity.User
	user2, _ = (database.GetUser("45"))
	encoded, _ := json.Marshal(&user2)

	fmt.Println(string(encoded))

	//fmt.Println("GetFilm")
	//fmt.Println(database.GetFilm("1"))
	fmt.Println("GetPerson")
	fmt.Println(database.GetPerson("1"))
	//fmt.Println("GetStartship")
	//fmt.Println(database.GetStartship("1"))
	/*var h = graph_ql.ApiRegister()
	apiPort := ":9091"
	handler := cors.Default().Handler(h)
	http.Handle("/graphql", handler)
	fmt.Println("server run on port:", apiPort)
	http.ListenAndServe(apiPort, nil)*/
}
