package search_test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"encoding/json"

	server "github.com/JekaTatsiy/grpc-market/http/server"
	"github.com/gorilla/mux"

	"flag"

	repo "github.com/JekaTatsiy/grpc-market/http/search"
	sugg "github.com/JekaTatsiy/grpc-market/http/suggest"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/format"
)

var searchaddr = flag.String("s", "grpc-search:1000", "adres grpc-search service")

func TestHTTPSearch(t *testing.T) {
	format.MaxLength = 0

	RegisterFailHandler(Fail)
	RunSpecs(t, "HTTPSearch")
}

var _ = Describe("HTTPSearch", func() {
	flag.Parse()
	g := server.NewGrpcClient("0.0.0.0:1000")
	//g := server.NewGrpcClient(*searchaddr)

	find := repo.Find(g)

	BeforeSuite(func() {
		post := sugg.Post(g)
		r := httptest.NewRequest(http.MethodGet, "/suggest", nil)
		r.Form = url.Values{
			"id":    []string{"1"},
			"link":  []string{"abc"},
			"title": []string{"NAME"},
			"query": []string{"автомобиль"}}
		r.Header.Set("Content-Type", "maultipart/form-data")
		w := httptest.NewRecorder()
		post(w, r)
		time.Sleep(1000 * time.Millisecond)

	})

	AfterSuite(func() {
		del := sugg.Delete(g)
		r := httptest.NewRequest(http.MethodDelete, "/suggest", nil)
		w := httptest.NewRecorder()
		del(w, r)
	})

	Context("Public functions", func() {
		When("find", func() {
			It("Success", func() {

				r := httptest.NewRequest(http.MethodGet, "/find/автомобиль", nil)
				r = mux.SetURLVars(r, map[string]string{"q": "автомобиль"})
				w := httptest.NewRecorder()

				find(w, r)
				suggests := &struct {
					Suggests []sugg.Suggest `json:"suggests"`
				}{}
				e := json.Unmarshal(w.Body.Bytes(), &suggests)
				Expect(e).ShouldNot(HaveOccurred())
				Expect(len(suggests.Suggests)).Should(Equal(1))

			})
		})
	})
})
