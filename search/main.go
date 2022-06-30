package main

import "github.com/JekaTatsiy/grpc-market/search/server"

var searchport string = "1000"
var esaddr string = "grpc-es:9200"

func main() {
	_ = server.NewServer(searchport, esaddr)
}
