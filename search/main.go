package main

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"

	pb "github.com/JekaTatsiy/grpc-market/suggest_proto"
)

type server struct {
	pb.UnimplementedSuggestServiceServer
}

func (s *server) Search(ctx context.Context, in *pb.SearchRequest) (*pb.SearchResponse, error) {
	responses := make([]*pb.SuggestResponse, 0)
	responses = append(responses, &pb.SuggestResponse{ID: 1, LinkUrl: "l", Title: "t", Queries: []string{"a", "b"}, Active: true})
	return &pb.SearchResponse{Suggests: responses}, nil
}

func main() {

	lis, err := net.Listen("tcp", "1000")
	if err != nil {
		fmt.Printf("failed to listen: %v\n", err)
		return
	}

	s := grpc.NewServer()
	pb.RegisterSuggestServiceServer(s, &server{})
	fmt.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		fmt.Printf("failed to serve: %v\n", err)
		return
	}
}
