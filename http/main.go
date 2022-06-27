package main

import (
	"fmt"
	"net/http"

	suggest "github.com/JekaTatsiy/grpc-market/http/suggest"
	"github.com/gorilla/mux"
)

func main() {

	fmt.Println("http")

	r := mux.NewRouter()
	suggest.GenRouting(r)

	server := http.Server{Addr: "0.0.0.0:3000", Handler: r}
	server.ListenAndServe()
}
