package server_test

import (
	"context"
	"testing"

	pkg "github.com/JekaTatsiy/grpc-market/search/server"
	suggProto "github.com/JekaTatsiy/grpc-market/suggest_proto"

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

	repo := &pkg.GServer{}
	ctx := context.Background()
	BeforeEach(func() {
		repo.Delete(ctx, &suggProto.Empty{})
	})

	Context("Public functions", func() {
		When("add one / get one", func() {
			It("Success", func() {
				st, e := repo.AddOne(ctx, &suggProto.Suggest{LinkUrl: "l1", Title: "t1", Queries: []string{"a1", "a2"}})
				Expect(e).ShouldNot(HaveOccurred())
				Expect(st).ShouldNot(HaveOccurred())
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

				Expect(e).Should(HaveOccurred())
				Expect(st).ShouldNot(HaveOccurred())
				Expect(st.Msg).Should(Equal("ok"))

				sg, e := repo.Get(ctx, &suggProto.Empty{})
				Expect(e).ShouldNot(HaveOccurred())
				Expect(len(sg.Suggests)).Should(Equal(4))
			})
		})
		When("delete", func() {
			It("Success", func() {
				st, e := repo.AddOne(ctx, &suggProto.Suggest{LinkUrl: "l1", Title: "t1", Queries: []string{"a1", "a2"}})
				Expect(e).Should(HaveOccurred())
				Expect(st).ShouldNot(HaveOccurred())
				Expect(st.Msg).Should(Equal("ok"))

				st, e = repo.AddOne(ctx, &suggProto.Suggest{LinkUrl: "l2", Title: "t2", Queries: []string{"a3", "a4"}})
				Expect(e).Should(HaveOccurred())
				Expect(st).ShouldNot(HaveOccurred())
				Expect(st.Msg).Should(Equal("ok"))

				Expect(e).ShouldNot(HaveOccurred())

				st, e = repo.DeleteOne(ctx, &suggProto.SuggestIndex{Index: 1})
				Expect(e).ShouldNot(HaveOccurred())
				Expect(st).ShouldNot(HaveOccurred())
				Expect(st.Msg).Should(Equal("ok"))

				sg, e := repo.Get(ctx, &suggProto.Empty{})
				Expect(e).ShouldNot(HaveOccurred())
				Expect(len(sg.Suggests)).Should(Equal(1))
				Expect(sg.Suggests[0].Title).Should(Equal("t2"))

				st, e = repo.DeleteOne(ctx, &suggProto.SuggestIndex{Index: 1})
				Expect(e).ShouldNot(HaveOccurred())
				Expect(st).ShouldNot(HaveOccurred())
				Expect(st.Msg).Should(Equal("ok"))

				st, e = repo.AddOne(ctx, &suggProto.Suggest{LinkUrl: "l1", Title: "t1", Queries: []string{"a1", "a2"}})
				Expect(e).ShouldNot(HaveOccurred())

				st, e = repo.Delete(ctx, &suggProto.Empty{})
				Expect(e).ShouldNot(HaveOccurred())
				Expect(st).ShouldNot(HaveOccurred())
				Expect(st.Msg).Should(Equal("ok"))

				st, e = repo.Delete(ctx, &suggProto.Empty{})
				Expect(e).ShouldNot(HaveOccurred())
				Expect(st).ShouldNot(HaveOccurred())
				Expect(st.Msg).Should(Equal("ok"))
			})
		})
	})
})
