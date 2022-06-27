package search

import (
	"encoding/json"
	"net/http"

	"github.com/JekaTatsiy/grpc-market/http/suggest"
	"github.com/gorilla/mux"
)

func Find() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s := suggest.Suggest{}
		json.NewEncoder(w).Encode(s)
	}
}

func GenRouting(r *mux.Router) {
	r.HandleFunc("find/{q:[/d/w]+}", Find()).Methods(http.MethodGet)
}
