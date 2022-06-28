package server

import (
	"fmt"
	"net/http"

	pb "github.com/JekaTatsiy/grpc-market/suggest_proto"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Server struct {
	HTTPServer *http.Server
	GRPCClient pb.SuggestServiceClient
}

func NewGrpcClient(port string) pb.SuggestServiceClient {
	conn, err := grpc.Dial(":"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("did not connect: %v", err)
		return nil
	}
	return pb.NewSuggestServiceClient(conn)
}

func NewServer(port string, routing *mux.Router, grpcClient pb.SuggestServiceClient) *Server {
	s := &Server{}
	s.HTTPServer = &http.Server{Addr: "0.0.0.0:" + port, Handler: routing}
	s.GRPCClient = grpcClient
	return s
}
