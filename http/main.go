package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	fmt.Println("http")

	r := mux.NewRouter()
	server := http.Server{Addr: "0.0.0.0:3000", Handler: r}
	server.ListenAndServe()
}
