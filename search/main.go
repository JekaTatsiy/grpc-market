package main

import (
	"flag"
	"fmt"

	"github.com/JekaTatsiy/grpc-market/search/server"
)

var searchport string = "1000"
var esaddr = flag.String("s", "grpc-es:9200", "adres es-search service")

func main() {
	flag.Parse()
	fmt.Println("elastic adres ", *esaddr)
	_ = server.NewServer(searchport, *esaddr)
}
