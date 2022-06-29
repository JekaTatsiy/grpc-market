package server

import (
	"context"

	pb "github.com/JekaTatsiy/grpc-market/suggest_proto"
)

func (s *GServer) Search(ctx context.Context, in *pb.SearchRequest) (*pb.SearchResponse, error) {
	responses := make([]*pb.SuggestResponse, 0)
	responses = append(responses, &pb.SuggestResponse{ID: 1, LinkUrl: "l", Title: "t", Queries: []string{"a", "b"}, Active: true})
	return &pb.SearchResponse{Suggests: responses}, nil
}




func (s *GServer) AddOne(context.Context, *pb.SuggestRequest) (*pb.Status, error) {



	stat := &pb.Status{Msg: "ok"}
	return stat, nil
}

func (s *GServer) AddFile(context.Context, *pb.CSV) (*pb.Status, error) {
	stat := &pb.Status{Msg: "ok"}
	return stat, nil
}

func (s *GServer) GetOne(context.Context, *pb.SuggestRequestIndex) (*pb.SuggestResponse, error) {
	r := &pb.SuggestResponse{ID: 0, LinkUrl: "l", Title: "t", Queries: []string{"a"}, Active: true}

	return r, nil
}

func (s *GServer) Get(context.Context, *pb.Empty) (*pb.SuggestResponseArray, error) {
	r := &pb.SuggestResponseArray{
		Suggests: []*pb.SuggestResponse{{ID: 1, LinkUrl: "l", Title: "t", Queries: []string{"a"}, Active: true}},
	}
	return r, nil
}

func (s *GServer) DeleteOne(context.Context, *pb.SuggestRequestIndex) (*pb.Status, error) {
	stat := &pb.Status{Msg: "ok"}
	return stat, nil
}

func (s *GServer) Delete(context.Context, *pb.Empty) (*pb.Status, error) {
	stat := &pb.Status{Msg: "ok"}
	return stat, nil
}
