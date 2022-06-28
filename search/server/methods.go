package server

import (
	"context"
	pb "github.com/JekaTatsiy/grpc-market/suggest_proto"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

func (s *GServer) Search(ctx context.Context, in *pb.SearchRequest) (*pb.SearchResponse, error) {
	responses := make([]*pb.SuggestResponse, 0)
	responses = append(responses, &pb.SuggestResponse{ID: 1, LinkUrl: "l", Title: "t", Queries: []string{"a", "b"}, Active: true})
	return &pb.SearchResponse{Suggests: responses}, nil
}

func (s *GServer) AddOne(context.Context, *pb.SuggestRequest) (*pb.Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddOne not implemented")
}
func (s *GServer) AddFile(context.Context, *pb.CSV) (*pb.Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddFile not implemented")
}
func (s *GServer) GetOne(context.Context, *pb.SuggestRequestIndex) (*pb.SuggestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOne not implemented")
}
func (s *GServer) Get(context.Context, *pb.Empty) (*pb.SuggestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (s *GServer) DeleteOne(context.Context, *pb.SuggestRequestIndex) (*pb.Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteOne not implemented")
}
func (s *GServer) Delete(context.Context, *pb.Empty) (*pb.Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
