package server_test

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/JekaTatsiy/grpc-market/http/server"
	sugg "github.com/JekaTatsiy/grpc-market/http/suggest"
	pkg "github.com/JekaTatsiy/grpc-market/search/server"
	suggProto "github.com/JekaTatsiy/grpc-market/suggest_proto"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/format"
)

func NewSuggest(link, title string, queries []string) {
	post := sugg.Post(server.NewGrpcClient("1000"))

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

func ExtractLinks(response *suggProto.SearchResponse) []string {
	links := []string{}
	for _, x := range response.Suggests {
		links = append(links, x.LinkUrl)
	}
	return links
}

var _ = Describe("Search", func() {
	repo := &pkg.GServer{}
	ctx := context.Background()

	BeforeSuite(func() {
		NewSuggest("link1", "title1", []string{"Автомобиль", "Машина", "Ехать"})
		NewSuggest("link2", "title2", []string{"Поезд", "Железная дорога"})
		NewSuggest("link3", "title3", []string{"Самолет", "Вертолет", "Воздушный шар"})
		NewSuggest("link4", "title4", []string{"Корабль", "Круиз", "Море"})
	})

	Context("Public functions", func() {
		When("find", func() {
			It("Success", func() {
				res, e := repo.Search(ctx, &suggProto.SearchRequest{Query: "как доехать в воронеж"})
				Expect(e).ShouldNot(HaveOccurred())
				Expect(ExtractLinks(res)).Should(Equal([]string{"link1"}))
			})
		})
		When("find", func() {
			It("Success", func() {
				res, e := repo.Search(ctx, &suggProto.SearchRequest{Query: "скоо тоит саолет"})
				Expect(e).ShouldNot(HaveOccurred())
				Expect(ExtractLinks(res)).Should(Equal([]string{"link3"}))
			})
		})
		When("find", func() {
			It("Success", func() {
				res, e := repo.Search(ctx, &suggProto.SearchRequest{Query: "dtc dthnjktnf"})
				Expect(e).ShouldNot(HaveOccurred())
				Expect(ExtractLinks(res)).Should(Equal([]string{"link2"}))
			})
		})
		When("find", func() {
			It("Success", func() {
				res, e := repo.Search(ctx, &suggProto.SearchRequest{Query: "PUTICHESTVIE NA KORABLE"})
				Expect(e).ShouldNot(HaveOccurred())
				Expect(ExtractLinks(res)).Should(Equal([]string{"link4"}))
			})
		})
	})
})
