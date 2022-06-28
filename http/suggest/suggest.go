package suggest

import (
	"net/http"

	pb "github.com/JekaTatsiy/grpc-market/suggest_proto"
	"github.com/gorilla/mux"
)

type Suggest struct {
	ID      int
	LinkURL string
	Title   string
	Queries []string
	Active  bool
}

type Status struct {
	Status string `json:"string"`
}

const searchService = "172.20.0.1:1000"

func GetAll(grpcClient pb.SuggestServiceClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("get all"))
	}
}

func Get(grpcClient pb.SuggestServiceClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		v := mux.Vars(r)
		w.Write([]byte("get " + v["id"]))
	}
}

func Post(grpcClient pb.SuggestServiceClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("post"))
	}
}

func Delete(grpcClient pb.SuggestServiceClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("delete"))
	}
}

func Import(grpcClient pb.SuggestServiceClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("import"))
	}
}

func GenRouting(r *mux.Router, grpcClient pb.SuggestServiceClient) {
	r.HandleFunc("/suggest", GetAll(grpcClient)).Methods(http.MethodGet)
	r.HandleFunc("/suggest/{id:[/d]+}", Get(grpcClient)).Methods(http.MethodGet)
	r.HandleFunc("/suggest", Post(grpcClient)).Methods(http.MethodPost)
	r.HandleFunc("/suggest", Delete(grpcClient)).Methods(http.MethodDelete)
	r.HandleFunc("/suggest/import", Import(grpcClient)).Methods(http.MethodPost)
}
