package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	fmt.Println("search")

	r := mux.NewRouter() //.PathPrefix("/api").Subrouter()
	r.HandleFunc("ping", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("ok")) }).Methods("GET")
	server := http.Server{Addr: "0.0.0.0:1000", Handler: r}
	server.ListenAndServe()
}
