package server

import (
	"context"
	"encoding/csv"

	pb "github.com/JekaTatsiy/grpc-market/suggest_proto"
)

func (s *GServer) Search(ctx context.Context, req *pb.SearchRequest) (*pb.SearchResponse, error) {
	suggs := s.ESSearch(req.Query)
	return &pb.SearchResponse{Suggests: suggs}, nil
}

func (s *GServer) AddOne(ctx context.Context, sugg *pb.Suggest) (*pb.Status, error) {
	e := s.ESAdd([]*pb.Suggest{sugg})
	m := "ok"
	if e != nil {
		m = e.Error()
	}
	stat := &pb.Status{Msg: m}
	return stat, nil
}

func (s *GServer) AddFile(ctd context.Context, csv *pb.CSV) (*pb.Status, error) {
	sugg := make([]*pb.Suggest, 0)

	
	e := s.ESAdd(sugg)
	m := "ok"
	if e != nil {
		m = e.Error()
	}
	stat := &pb.Status{Msg: m}
	return stat, nil
}

func (s *GServer) GetOne(ctx context.Context, i *pb.SuggestIndex) (*pb.Suggest, error) {
	r, e := s.ESGetOne(i.Index)
	return r, e
}

func (s *GServer) Get(ctx context.Context, em *pb.Empty) (*pb.SuggestArray, error) {
	r, e := s.ESGet()
	arr := &pb.SuggestArray{Suggests: r}
	return arr, e
}

func (s *GServer) DeleteOne(ctx context.Context, i *pb.SuggestIndex) (*pb.Status, error) {
	e := s.ESDeleteOne(i.Index)
	m := "ok"
	if e != nil {
		m = e.Error()
	}
	return &pb.Status{Msg: m}, nil
}

func (s *GServer) Delete(context.Context, *pb.Empty) (*pb.Status, error) {
	s.ESDeleteIndex()
	s.ESCreateIndex()
	return &pb.Status{Msg: "ok"}, nil
}
