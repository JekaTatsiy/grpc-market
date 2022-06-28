package main

import (
	"fmt"
	suggest "github.com/JekaTatsiy/grpc-market/http/suggest"

	"github.com/JekaTatsiy/grpc-market/http/server"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("start http server")

	r := mux.NewRouter()

	g := server.NewGrpcClient("1000")
	suggest.GenRouting(r, g)

	s := server.NewServer("3000", r, g)
	s.HTTPServer.ListenAndServe()

}
