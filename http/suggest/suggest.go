package suggest

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"

	pb "github.com/JekaTatsiy/grpc-market/suggest_proto"
	"github.com/gorilla/mux"
)

type Suggest struct {
	ID      int      `json:"id"`
	LinkUrl string   `json:"linkUrl"`
	Title   string   `json:"title"`
	Queries []string `json:"queries"`
}

type Status struct {
	Status string `json:"Msg"`
}

func GetAll(grpcClient pb.SuggestServiceClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		suggs, e := grpcClient.Get(ctx, &pb.Empty{})
		if e != nil {
			json.NewEncoder(w).Encode(Status{Status: e.Error()})
			return
		}
		if suggs == nil {
			suggs = &pb.SuggestArray{}
		}
		json.NewEncoder(w).Encode(suggs.Suggests)
	}
}

func Get(grpcClient pb.SuggestServiceClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		v := mux.Vars(r)

		ind_s, ok := v["id"]
		if !ok {
			json.NewEncoder(w).Encode(Status{Status: "id not found"})
			return
		}

		ind, e := strconv.Atoi(ind_s)
		if !ok {
			json.NewEncoder(w).Encode(Status{Status: e.Error()})
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		suggs, e := grpcClient.GetOne(ctx, &pb.SuggestIndex{Index: int32(ind)})
		if e != nil {
			json.NewEncoder(w).Encode(Status{Status: e.Error()})
			return
		}
		json.NewEncoder(w).Encode(suggs)
	}
}

func Post(grpcClient pb.SuggestServiceClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		id_str := r.FormValue("id")
		if id_str == "" {
			json.NewEncoder(w).Encode(Status{Status: "expect id"})
			return
		}
		id, e := strconv.Atoi(id_str)
		if e != nil {
			json.NewEncoder(w).Encode(Status{Status: "id is not a number"})
			return
		}

		url := r.FormValue("link")
		if url == "" {
			json.NewEncoder(w).Encode(Status{Status: "expect url"})
			return
		}
		title := r.FormValue("title")
		if title == "" {
			json.NewEncoder(w).Encode(Status{Status: "expect title"})
			return
		}
		queries := r.Form["query"]
		if len(queries) == 0 {
			json.NewEncoder(w).Encode(Status{Status: "expect queries"})
			return
		}
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		status, e := grpcClient.AddOne(ctx, &pb.Suggest{ID: int32(id), LinkUrl: url, Title: title, Queries: queries})
		if e != nil {
			json.NewEncoder(w).Encode(Status{Status: e.Error()})
			return
		}
		json.NewEncoder(w).Encode(status)
	}
}

func Import(grpcClient pb.SuggestServiceClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		b, e := io.ReadAll(r.Body)
		if e != nil {
			json.NewEncoder(w).Encode(Status{Status: e.Error()})
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		status, e := grpcClient.Delete(ctx, &pb.Empty{})
		if e != nil {
			json.NewEncoder(w).Encode(Status{Status: e.Error()})
			return
		}
		ctx, cancel = context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		status, e = grpcClient.AddFile(ctx, &pb.CSV{Text: b})
		if e != nil {
			json.NewEncoder(w).Encode(Status{Status: e.Error()})
			return
		}
		json.NewEncoder(w).Encode(status)
	}
}

func DeleteOne(grpcClient pb.SuggestServiceClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		v := mux.Vars(r)

		ind_s, ok := v["id"]
		if !ok {
			json.NewEncoder(w).Encode(Status{Status: "id not found"})
			return
		}

		ind, e := strconv.Atoi(ind_s)
		if !ok {
			json.NewEncoder(w).Encode(Status{Status: e.Error()})
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		status, e := grpcClient.DeleteOne(ctx, &pb.SuggestIndex{Index: int32(ind)})
		if e != nil {
			json.NewEncoder(w).Encode(Status{Status: e.Error()})
			return
		}
		json.NewEncoder(w).Encode(status)
	}
}
func Delete(grpcClient pb.SuggestServiceClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		status, e := grpcClient.Delete(ctx, &pb.Empty{})
		if e != nil {
			json.NewEncoder(w).Encode(Status{Status: e.Error()})
			return
		}
		json.NewEncoder(w).Encode(status)
	}
}
func GenRouting(r *mux.Router, grpcClient pb.SuggestServiceClient) {
	r.HandleFunc("/suggest", GetAll(grpcClient)).Methods(http.MethodGet)
	r.HandleFunc("/suggest/{id:[0-9]+}", Get(grpcClient)).Methods(http.MethodGet)
	r.HandleFunc("/suggest", Post(grpcClient)).Methods(http.MethodPost)
	r.HandleFunc("/suggest/import", Import(grpcClient)).Methods(http.MethodPost)
	r.HandleFunc("/suggest", Delete(grpcClient)).Methods(http.MethodDelete)
	r.HandleFunc("/suggest/{id:[0-9]+}", DeleteOne(grpcClient)).Methods(http.MethodDelete)
}
