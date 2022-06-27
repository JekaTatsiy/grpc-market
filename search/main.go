package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	fmt.Println("search")

	r := mux.NewRouter()
	server := http.Server{Addr: "0.0.0.0:1000", Handler: r}
	server.ListenAndServe()
}
