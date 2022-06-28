package search_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	sugg "github.com/JekaTatsiy/grpc-market/http/suggest"
	repo "github.com/JekaTatsiy/grpc-market/search/search"
	//"github.com/elastic/go-elasticsearch/v8"
	//"github.com/elastic/go-elasticsearch/v8/esapi"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/format"
)

func NewSuggest(link0, title0 string, queries []string) {
	post := sugg.Post()

	payload := &bytes.Buffer{}
	payload.WriteString("link0=link0&title0=title0&query=a&query=b&query=c")
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

func ExtractLinks(suggests []sugg.Suggest) []string {
	links := []string{}
	for _, x := range suggests {
		links = append(links, x.LinkURL)
	}
	return links
}

var _ = Describe("Search", func() {

	//es, err := elasticsearch.NewDefaultClient()
	//Expect(err).ShouldNot(HaveOccurred())

	BeforeSuite(func() {
		NewSuggest("link1", "title1", []string{"Автомобиль", "Машина", "Ехать"})
		NewSuggest("link2", "title2", []string{"Поезд", "Железная дорога"})
		NewSuggest("link3", "title3", []string{"Самолет", "Вертолет", "Воздушный шар"})
		NewSuggest("link4", "title4", []string{"Корабль", "Круиз", "Море"})
	})

	Context("Public functions", func() {
		When("find", func() {
			It("Success", func() {
				res := repo.Find("как доехать в воронеж")
				Expect(ExtractLinks(res)).Should(Equal([]string{"link1"}))
			})
		})
		When("find", func() {
			It("Success", func() {
				res := repo.Find("скоо тоит саолет")
				Expect(ExtractLinks(res)).Should(Equal([]string{"link3"}))
			})
		})
		When("find", func() {
			It("Success", func() {
				res := repo.Find("dtc dthnjktnf")
				Expect(ExtractLinks(res)).Should(Equal([]string{"link2"}))
			})
		})
		When("find", func() {
			It("Success", func() {
				res := repo.Find("PUTICHESTVIE NA KORABLE")
				Expect(ExtractLinks(res)).Should(Equal([]string{"link4"}))
			})
		})
	})
})
