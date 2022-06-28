package suggest_test

import (
	"testing"

	"github.com/JekaTatsiy/grpc-market/http/suggest"
	repo "github.com/JekaTatsiy/grpc-market/search/suggest"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/format"
)

func TestHTTPSuggest(t *testing.T) {
	format.MaxLength = 0

	RegisterFailHandler(Fail)
	RunSpecs(t, "HTTPSuggest")
}

var _ = Describe("HTTPSuggest", func() {

	BeforeEach(func() {
		repo.DeleteAll()
	})

	Context("Public functions", func() {
		When("add one / get one", func() {
			It("Success", func() {
				e := repo.CreateOne(&suggest.Suggest{ID: 1, LinkURL: "l1", Title: "t1", Queries: []string{"a1", "a2"}, Active: true})
				Expect(e).ShouldNot(HaveOccurred())
				sg, e := repo.GetOne(1)
				Expect(e).ShouldNot(HaveOccurred())
				Expect(sg.Title).Should(Equal("t1"))
			})
		})
		When("get if empty", func() {
			It("Success", func() {
				_, e := repo.GetOne(1)
				Expect(e).Should(HaveOccurred())

				sg, e := repo.GetAll()
				Expect(e).ShouldNot(HaveOccurred())
				Expect(len(sg)).Should(Equal(0))
			})
		})
		When("from file", func() {
			It("Success", func() {
				e := repo.CreateFromFile([]byte(`link1,title1,q1|q2
				link2,title2,q1|q2|q3
				link3,title3,q1|q2
				link4,title4,q1
				link5,title5,q1|q2`))

				Expect(e).Should(HaveOccurred())

				sg, e := repo.GetAll()
				Expect(e).ShouldNot(HaveOccurred())
				Expect(len(sg)).Should(Equal(4))
			})
		})
		When("delete", func() {
			It("Success", func() {
				e := repo.CreateOne(&suggest.Suggest{ID: 1, LinkURL: "l1", Title: "t1", Queries: []string{"a1", "a2"}, Active: true})
				e = repo.CreateOne(&suggest.Suggest{ID: 2, LinkURL: "l2", Title: "t2", Queries: []string{"a3", "a4"}, Active: true})
				Expect(e).ShouldNot(HaveOccurred())

				e = repo.DeleteOne(1)
				Expect(e).ShouldNot(HaveOccurred())

				sg, e := repo.GetAll()
				Expect(e).ShouldNot(HaveOccurred())
				Expect(len(sg)).Should(Equal(1))
				Expect(sg[0].Title).Should(Equal("t2"))

				e = repo.DeleteOne(1)
				Expect(e).ShouldNot(HaveOccurred())

				e = repo.CreateOne(&suggest.Suggest{ID: 1, LinkURL: "l1", Title: "t1", Queries: []string{"a1", "a2"}, Active: true})
				Expect(e).ShouldNot(HaveOccurred())

				e = repo.DeleteAll()
				Expect(e).ShouldNot(HaveOccurred())
				e = repo.DeleteAll()
				Expect(e).ShouldNot(HaveOccurred())
			})
		})
	})
})
