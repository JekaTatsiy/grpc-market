package search

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	pb "github.com/JekaTatsiy/grpc-market/suggest_proto"

	"github.com/JekaTatsiy/grpc-market/http/suggest"
	"github.com/gorilla/mux"
)

func Find(grpcClient pb.SuggestServiceClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		v := mux.Vars(r)

		query, ok := v["q"]
		if !ok {
			json.NewEncoder(w).Encode(suggest.Status{Status: "query not found"})
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		suggs, e := grpcClient.Search(ctx, &pb.SearchRequest{Query: query})
		if e != nil {
			json.NewEncoder(w).Encode(suggest.Status{Status: e.Error()})
			return
		}
		json.NewEncoder(w).Encode(suggs)
	}
}

func GenRouting(r *mux.Router, grpcClient pb.SuggestServiceClient) {
	r.HandleFunc("/find/{q:.+}", Find(grpcClient)).Methods(http.MethodGet)
}
