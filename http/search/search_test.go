package search_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"encoding/json"

	repo "github.com/JekaTatsiy/grpc-market/http/search"
	sugg "github.com/JekaTatsiy/grpc-market/http/suggest"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/format"
)

func TestAdvProductRepository(t *testing.T) {
	format.MaxLength = 0

	RegisterFailHandler(Fail)
	RunSpecs(t, "AdvProductRepository")
}

func NewSuggest(link, title string, queries []string) {
	post := sugg.Post()

	payload := &bytes.Buffer{}
	payload.WriteString("link=abc&title=NAME&query=a&query=b&query=c")
	r := httptest.NewRequest(http.MethodGet, "/suggest", payload)
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	post(w, r)
}

var _ = Describe("AdvProductRepository", func() {

	find := repo.Find()

	BeforeSuite(func() {
		NewSuggest("abc", "name", []string{"a", "b", "c"})
	})

	Context("Public functions", func() {
		When("find", func() {
			It("Success", func() {

				r := httptest.NewRequest(http.MethodGet, "/find/автомобиль", nil)
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
