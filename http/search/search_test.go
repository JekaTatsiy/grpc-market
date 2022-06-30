package search_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

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
	//g := server.NewGrpcClient(*searchaddr)
	g := server.NewGrpcClient("0.0.0.0:9200")

	find := repo.Find(g)

	BeforeSuite(func() {
		post := sugg.Post(g)
		payload := &bytes.Buffer{}
		payload.WriteString("link=abc&title=NAME&query=автомобиль")
		r := httptest.NewRequest(http.MethodGet, "/suggest", payload)
		r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		w := httptest.NewRecorder()
		post(w, r)

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

				suggests := make([]sugg.Suggest, 0)
				e := json.Unmarshal(w.Body.Bytes(), &suggests)
				Expect(e).ShouldNot(HaveOccurred())
				Expect(len(suggests) > 0).Should(BeTrue())

			})
		})
	})
})
