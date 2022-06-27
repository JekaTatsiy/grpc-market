package suggest

import (
	"net/http"

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

func GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("get all"))
	}
}

func Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		v := mux.Vars(r)
		w.Write([]byte("get " + v["id"]))
	}
}

func Post() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("post"))
	}
}

func Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("delete"))
	}
}

func Import() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("import"))
	}
}

func GenRouting(r *mux.Router) {
	r.HandleFunc("/suggest", GetAll()).Methods(http.MethodGet)
	r.HandleFunc("/suggest/{id:[/d]+}", Get()).Methods(http.MethodGet)
	r.HandleFunc("/suggest", Post()).Methods(http.MethodPost)
	r.HandleFunc("/suggest", Delete()).Methods(http.MethodDelete)
	r.HandleFunc("/suggest/import", Import()).Methods(http.MethodPost)
}
