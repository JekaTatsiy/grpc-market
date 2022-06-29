package server

import (
	"fmt"
	"net"

	pb "github.com/JekaTatsiy/grpc-market/suggest_proto"
	"google.golang.org/grpc"
)

type Server struct {
	
}

type GServer struct {
	pb.UnimplementedSuggestServiceServer
}

func NewServer(port, esaddr string) *Server {
	s := &Server{}
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Printf("failed to listen: %v\n", err)
		return nil
	}

	gs := grpc.NewServer()
	pb.RegisterSuggestServiceServer(gs, &GServer{})
	fmt.Printf("server listening at %v", lis.Addr())
	if err := gs.Serve(lis); err != nil {
		fmt.Printf("failed to serve: %v\n", err)
		return nil
	}

	return s
}
