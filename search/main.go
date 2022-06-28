package main

import "github.com/JekaTatsiy/grpc-market/search/server"

func main() {
	_ = server.NewServer("1000")
}
