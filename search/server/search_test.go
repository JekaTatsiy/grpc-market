package server_test

import (
	"context"
	"testing"
	"time"

	"flag"

	pkg "github.com/JekaTatsiy/grpc-market/search/server"
	suggProto "github.com/JekaTatsiy/grpc-market/suggest_proto"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/format"
)

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

var esaddr = flag.String("s", "grpc-es:9200", "adres grpc-search service")

var _ = Describe("Search", func() {
	flag.Parse()
	//repo := &pkg.GServer{ESaddr: *esaddr}
	repo := &pkg.GServer{ESaddr: "0.0.0.0:9200"}

	ctx := context.Background()

	BeforeEach(func() {
		repo.Delete(ctx, &suggProto.Empty{})
	})

	Context("Public functions", func() {
		When("find", func() {
			It("Success", func() {
				ctx := context.Background()
				repo.AddOne(ctx, &suggProto.Suggest{ID: 1, LinkUrl: "link1", Title: "title1", Queries: []string{"Автомобиль", "Машина", "Ехать"}})
				repo.AddOne(ctx, &suggProto.Suggest{ID: 2, LinkUrl: "link2", Title: "title2", Queries: []string{"Поезд", "Железная дорога"}})
				repo.AddOne(ctx, &suggProto.Suggest{ID: 3, LinkUrl: "link3", Title: "title3", Queries: []string{"Самолет", "Вертолет", "Воздушный шар"}})
				repo.AddOne(ctx, &suggProto.Suggest{ID: 4, LinkUrl: "link4", Title: "title4", Queries: []string{"Корабль", "Круиз", "Море"}})
				time.Sleep(1000 * time.Millisecond)
				res, e := repo.Search(ctx, &suggProto.SearchRequest{Query: "как доехать в воронеж"})
				Expect(e).ShouldNot(HaveOccurred())
				Expect(ExtractLinks(res)[0]).Should(Equal("link2"))
			})
		})
		When("find", func() {
			It("Success", func() {
				ctx := context.Background()
				repo.AddOne(ctx, &suggProto.Suggest{ID: 1, LinkUrl: "link1", Title: "title1", Queries: []string{"Автомобиль", "Машина", "Ехать"}})
				repo.AddOne(ctx, &suggProto.Suggest{ID: 2, LinkUrl: "link2", Title: "title2", Queries: []string{"Поезд", "Железная дорога"}})
				repo.AddOne(ctx, &suggProto.Suggest{ID: 3, LinkUrl: "link3", Title: "title3", Queries: []string{"Самолет", "Вертолет", "Воздушный шар"}})
				repo.AddOne(ctx, &suggProto.Suggest{ID: 4, LinkUrl: "link4", Title: "title4", Queries: []string{"Корабль", "Круиз", "Море"}})
				time.Sleep(1000 * time.Millisecond)
				res, e := repo.Search(ctx, &suggProto.SearchRequest{Query: "скоо тоит саолет"})
				Expect(e).ShouldNot(HaveOccurred())
				Expect(ExtractLinks(res)[0]).Should(Equal("link3"))
			})
		})
		When("find", func() {
			It("Success", func() {
				ctx := context.Background()
				repo.AddOne(ctx, &suggProto.Suggest{ID: 1, LinkUrl: "link1", Title: "title1", Queries: []string{"Автомобиль", "Машина", "Ехать"}})
				repo.AddOne(ctx, &suggProto.Suggest{ID: 2, LinkUrl: "link2", Title: "title2", Queries: []string{"Поезд", "Железная дорога"}})
				repo.AddOne(ctx, &suggProto.Suggest{ID: 3, LinkUrl: "link3", Title: "title3", Queries: []string{"Самолет", "Вертолет", "Воздушный шар"}})
				repo.AddOne(ctx, &suggProto.Suggest{ID: 4, LinkUrl: "link4", Title: "title4", Queries: []string{"Корабль", "Круиз", "Море"}})
				time.Sleep(1000 * time.Millisecond)
				res, e := repo.Search(ctx, &suggProto.SearchRequest{Query: "dtc dthnjktnf"})
				Expect(e).ShouldNot(HaveOccurred())
				Expect(ExtractLinks(res)[0]).Should(Equal("link3"))
			})
		})
		When("find", func() {
			It("Success", func() {
				ctx := context.Background()
				repo.AddOne(ctx, &suggProto.Suggest{ID: 1, LinkUrl: "link1", Title: "title1", Queries: []string{"Автомобиль", "Машина", "Ехать"}})
				repo.AddOne(ctx, &suggProto.Suggest{ID: 2, LinkUrl: "link2", Title: "title2", Queries: []string{"Поезд", "Железная дорога"}})
				repo.AddOne(ctx, &suggProto.Suggest{ID: 3, LinkUrl: "link3", Title: "title3", Queries: []string{"Самолет", "Вертолет", "Воздушный шар"}})
				repo.AddOne(ctx, &suggProto.Suggest{ID: 4, LinkUrl: "link4", Title: "title4", Queries: []string{"Корабль", "Круиз", "Море"}})
				time.Sleep(1000 * time.Millisecond)
				res, e := repo.Search(ctx, &suggProto.SearchRequest{Query: "PUTICHESTVIE NA KORABLE"})
				Expect(e).ShouldNot(HaveOccurred())
				Expect(ExtractLinks(res)[0]).Should(Equal("link4"))
			})
		})

		Context("Public functions", func() {
			When("add one / get one", func() {
				It("Success", func() {
					st, e := repo.AddOne(ctx, &suggProto.Suggest{ID: 1, LinkUrl: "l1", Title: "t1", Queries: []string{"a1", "a2"}})
					Expect(e).ShouldNot(HaveOccurred())
					Expect(st.Msg).Should(Equal("ok"))

					sg, e := repo.GetOne(ctx, &suggProto.SuggestIndex{Index: 1})
					Expect(e).ShouldNot(HaveOccurred())
					Expect(sg.Title).Should(Equal("t1"))
				})
			})
			When("get if empty", func() {
				It("Success", func() {
					_, e := repo.GetOne(ctx, &suggProto.SuggestIndex{Index: 1})
					Expect(e).Should(HaveOccurred())

					sg, e := repo.Get(ctx, &suggProto.Empty{})
					Expect(e).ShouldNot(HaveOccurred())
					Expect(len(sg.Suggests)).Should(Equal(0))
				})
			})
			When("from file", func() {
				It("Success", func() {
					st, e := repo.AddFile(ctx, &suggProto.CSV{Text: []byte(`
link1,title1,q1|q2
link2,title2,q1|q2|q3
link3,title3,q1|q2
link4,title4,q1
link5,title5,q1|q2`)})

					Expect(e).ShouldNot(HaveOccurred())
					Expect(st.Msg).Should(Equal("ok"))
					time.Sleep(1000 * time.Millisecond)
					sg, e := repo.Get(ctx, &suggProto.Empty{})
					Expect(e).ShouldNot(HaveOccurred())
					Expect(len(sg.Suggests)).Should(Equal(5))
				})
			})
			When("delete", func() {
				It("Success", func() {
					st, e := repo.AddOne(ctx, &suggProto.Suggest{ID: 1, LinkUrl: "l1", Title: "t1", Queries: []string{"a1", "a2"}})
					Expect(e).ShouldNot(HaveOccurred())
					Expect(st.Msg).Should(Equal("ok"))

					st, e = repo.AddOne(ctx, &suggProto.Suggest{ID: 2, LinkUrl: "l2", Title: "t2", Queries: []string{"a3", "a4"}})
					Expect(e).ShouldNot(HaveOccurred())
					Expect(st.Msg).Should(Equal("ok"))

					time.Sleep(1000 * time.Millisecond)
					
					sg, e := repo.Get(ctx, &suggProto.Empty{})
					Expect(e).ShouldNot(HaveOccurred())
					Expect(len(sg.Suggests)).Should(Equal(2))
					
					st, e = repo.DeleteOne(ctx, &suggProto.SuggestIndex{Index: 1})
					Expect(e).ShouldNot(HaveOccurred())
					Expect(st.Msg).Should(Equal("ok"))
					
					time.Sleep(1000 * time.Millisecond)
					
					sg, e = repo.Get(ctx, &suggProto.Empty{})
					Expect(e).ShouldNot(HaveOccurred())
					Expect(len(sg.Suggests)).Should(Equal(1))
					Expect(sg.Suggests[0].Title).Should(Equal("t2"))

					st, e = repo.DeleteOne(ctx, &suggProto.SuggestIndex{Index: 1})
					Expect(e).ShouldNot(HaveOccurred())
					Expect(st.Msg).Should(Equal("ok"))

					st, e = repo.AddOne(ctx, &suggProto.Suggest{LinkUrl: "l1", Title: "t1", Queries: []string{"a1", "a2"}})
					Expect(e).ShouldNot(HaveOccurred())

					st, e = repo.Delete(ctx, &suggProto.Empty{})
					Expect(e).ShouldNot(HaveOccurred())
					Expect(st.Msg).Should(Equal("ok"))

					st, e = repo.Delete(ctx, &suggProto.Empty{})
					Expect(e).ShouldNot(HaveOccurred())
					Expect(st.Msg).Should(Equal("ok"))
				})
			})
		})
	})
})
