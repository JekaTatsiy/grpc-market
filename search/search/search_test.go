package search_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/JekaTatsiy/grpc-market/http/server"
	sugg "github.com/JekaTatsiy/grpc-market/http/suggest"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/format"
)

func NewSuggest(link, title string, queries []string) {
	post := sugg.Post(server.NewGrpcClient("1000"))

	payload := &bytes.Buffer{}
	payload.WriteString("link=abc&title=NAME&query=a&query=b&query=c")
	r := httptest.NewRequest(http.MethodGet, "/suggest", payload)
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	post(w, r)
}

func TestSearch(t *testing.T) {
	format.MaxLength = 0

	RegisterFailHandler(Fail)
	RunSpecs(t, "Search")
}

var _ = Describe("Search", func() {

	//es, err := elasticsearch.NewDefaultClient()
	//Expect(err).ShouldNot(HaveOccurred())

	BeforeSuite(func() {
		NewSuggest("abc", "name", []string{"a", "b", "c"})
		NewSuggest("abc", "name", []string{"a", "b", "c"})
		NewSuggest("abc", "name", []string{"a", "b", "c"})
		NewSuggest("abc", "name", []string{"a", "b", "c"})
		NewSuggest("abc", "name", []string{"a", "b", "c"})
		NewSuggest("abc", "name", []string{"a", "b", "c"})
		NewSuggest("abc", "name", []string{"a", "b", "c"})
		NewSuggest("abc", "name", []string{"a", "b", "c"})
		NewSuggest("abc", "name", []string{"a", "b", "c"})
		NewSuggest("abc", "name", []string{"a", "b", "c"})
		NewSuggest("abc", "name", []string{"a", "b", "c"})
		NewSuggest("abc", "name", []string{"a", "b", "c"})
	})

	Context("Public functions", func() {
		When("find", func() {
			It("Success", func() {

			})
		})
	})
})
