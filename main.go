package main

import (
	"fmt"
	"net/http"

	"github.com/liping/simpleWebServer/graph_ql"

	"github.com/rs/cors"
)

func main() {
	var h = graph_ql.ApiRegister()
	apiPort := ":9091"
	handler := cors.Default().Handler(h)
	http.Handle("/graphql", handler)
	fmt.Println("server run on port:", apiPort)
	http.ListenAndServe(apiPort, nil)
}
