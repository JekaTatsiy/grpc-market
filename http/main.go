package main

import (
	"fmt"
	search "github.com/JekaTatsiy/grpc-market/http/search"
	suggest "github.com/JekaTatsiy/grpc-market/http/suggest"

	"github.com/JekaTatsiy/grpc-market/http/server"
	"github.com/gorilla/mux"
)

var httpport string = "3000"
var searchaddr string = "grpc-search:1000"

func main() {
	fmt.Println("http server at " + httpport)

	r := mux.NewRouter()

	g := server.NewGrpcClient(searchaddr)
	suggest.GenRouting(r, g)
	search.GenRouting(r, g)

	s := server.NewServer(httpport, r, g)
	s.HTTPServer.ListenAndServe()
}
