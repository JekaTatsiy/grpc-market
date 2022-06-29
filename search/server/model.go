package server

import (
	"fmt"
	"net"

	pb "github.com/JekaTatsiy/grpc-market/suggest_proto"

	"google.golang.org/grpc"
)

type GServer struct {
	pb.UnimplementedSuggestServiceServer
	ESaddr string
}

func NewServer(port, esaddr string) *GServer {

	s := &GServer{}

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Printf("failed to listen: %v\n", err)
		return nil
	}
	s.ESaddr = esaddr
	go s.IndexCreateIfNotExist()

	gs := grpc.NewServer()
	pb.RegisterSuggestServiceServer(gs, s)
	fmt.Printf("server listening at %v", lis.Addr())
	if err := gs.Serve(lis); err != nil {
		fmt.Printf("failed to serve: %v\n", err)
		return nil
	}

	return s
}
