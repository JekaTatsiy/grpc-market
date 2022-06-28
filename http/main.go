package main

import (
	"fmt"
	suggest "github.com/JekaTatsiy/grpc-market/http/suggest"
	search "github.com/JekaTatsiy/grpc-market/http/search"

	"github.com/JekaTatsiy/grpc-market/http/server"
	"github.com/gorilla/mux"
)

var httpport string = "3000"
var searchport string = "1000"

func main() {
	fmt.Println("http server at " + httpport)

	r := mux.NewRouter()

	g := server.NewGrpcClient(searchport)
	suggest.GenRouting(r, g)
	search.GenRouting(r,g)

	s := server.NewServer(httpport, r, g)
	s.HTTPServer.ListenAndServe()
}
