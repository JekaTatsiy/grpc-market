package suggest_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"encoding/json"

	server "github.com/JekaTatsiy/grpc-market/http/server"
	repo "github.com/JekaTatsiy/grpc-market/http/suggest"
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
	
	g := server.NewGrpcClient("1000")
	
	getall := repo.GetAll(g)
	get := repo.Get(g)
	post := repo.Post(g)
	deleteone := repo.DeleteOne(g)
	delete := repo.Delete(g)

	BeforeEach(func() {
		r := httptest.NewRequest(http.MethodGet, "/suggest", nil)
		w := httptest.NewRecorder()
		delete(w, r)
	})

	Context("Public functions", func() {
		When("add one", func() {
			It("Success", func() {
				payload := &bytes.Buffer{}
				payload.WriteString("id=1&link=abc&title=NAME&query=a&query=b&query=c")

				r := httptest.NewRequest(http.MethodGet, "/suggest", payload)
				r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
				w := httptest.NewRecorder()

				post(w, r)

				res := &repo.Status{}
				e := json.Unmarshal(w.Body.Bytes(), &res)

				Expect(e).ShouldNot(HaveOccurred())
				Expect(res.Status).Should(Equal("ok"))

			})
		})

		When("get one", func() {
			It("Success", func() {
				payload := &bytes.Buffer{}
				payload.WriteString("id=1&link=abc&title=NAME&query=a&query=b&query=c")

				r := httptest.NewRequest(http.MethodGet, "/suggest", payload)
				r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
				w := httptest.NewRecorder()

				post(w, r)

				res := &repo.Status{}
				e := json.Unmarshal(w.Body.Bytes(), &res)
				Expect(e).ShouldNot(HaveOccurred())
				Expect(res.Status).Should(Equal("ok"))

				r = httptest.NewRequest(http.MethodGet, "/suggest/1", nil)
				r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
				w = httptest.NewRecorder()

				get(w, r)

				suggest := &repo.Suggest{}
				e = json.Unmarshal(w.Body.Bytes(), &suggest)

				Expect(e).ShouldNot(HaveOccurred())

				Expect(suggest.LinkURL).Should(Equal("abc"))
				Expect(suggest.Title).Should(Equal("NAME"))
				Expect(len(suggest.Queries)).Should(Equal(3))
				Expect(suggest.Queries).Should(ContainElements("a", "b", "c"))
			})
		})

		When("get all", func() {
			It("Success", func() {
				payload := &bytes.Buffer{}

				payload.WriteString("id=1&link=abc&title=NAME&query=a&query=b&query=c")
				r := httptest.NewRequest(http.MethodGet, "/suggest", payload)
				r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
				w := httptest.NewRecorder()
				post(w, r)
				res := &repo.Status{}
				e := json.Unmarshal(w.Body.Bytes(), &res)
				Expect(e).ShouldNot(HaveOccurred())
				Expect(res.Status).Should(Equal("ok"))

				payload = &bytes.Buffer{}

				payload.WriteString("id=2&link=xyz&title=ITEM&query=x&query=y&query=z")
				r = httptest.NewRequest(http.MethodGet, "/suggest", payload)
				r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
				w = httptest.NewRecorder()
				post(w, r)
				res = &repo.Status{}
				e = json.Unmarshal(w.Body.Bytes(), &res)
				Expect(e).ShouldNot(HaveOccurred())
				Expect(res.Status).Should(Equal("ok"))

				r = httptest.NewRequest(http.MethodGet, "/suggest", nil)
				w = httptest.NewRecorder()
				getall(w, r)

				suggests := make([]repo.Suggest, 0)
				e = json.Unmarshal(w.Body.Bytes(), &suggests)
				Expect(e).ShouldNot(HaveOccurred())
				Expect(len(suggests)).Should(Equal(2))

				Expect(suggests[1].LinkURL).Should(Equal("xyz"))
				Expect(suggests[1].Title).Should(Equal("ITEM"))
				Expect(len(suggests[1].Queries)).Should(Equal(3))
				Expect(suggests[1].Queries).Should(ContainElements("x", "y", "z"))

			})
		})

		When("delete", func() {
			It("Success", func() {
				payload := &bytes.Buffer{}

				payload.WriteString("id=1&link=abc&title=NAME&query=a&query=b&query=c")
				r := httptest.NewRequest(http.MethodGet, "/suggest", payload)
				r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
				w := httptest.NewRecorder()

				post(w, r)

				payload = &bytes.Buffer{}

				payload.WriteString("id=2&link=xyz&title=ITEM&query=x&query=y&query=z")
				r = httptest.NewRequest(http.MethodGet, "/suggest", payload)
				r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
				w = httptest.NewRecorder()

				post(w, r)

				res := &repo.Status{}
				e := json.Unmarshal(w.Body.Bytes(), &res)
				Expect(e).ShouldNot(HaveOccurred())
				Expect(res.Status).Should(Equal("ok"))

				//delete one
				r = httptest.NewRequest(http.MethodGet, "/suggest", nil)
				r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
				w = httptest.NewRecorder()

				deleteone(w, r)

				res = &repo.Status{}
				e = json.Unmarshal(w.Body.Bytes(), &res)
				Expect(e).ShouldNot(HaveOccurred())
				Expect(res.Status).Should(Equal("ok"))

				r = httptest.NewRequest(http.MethodGet, "/suggest", nil)
				r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
				w = httptest.NewRecorder()

				delete(w, r)

				res = &repo.Status{}
				e = json.Unmarshal(w.Body.Bytes(), &res)
				Expect(e).ShouldNot(HaveOccurred())
				Expect(res.Status).Should(Equal("ok"))
			})
		})

		When("import", func() {
			It("Success", func() {
				payload := &bytes.Buffer{}
				payload.WriteString(`
				link1,title1,q1|q2
				link2,title2,q1|q2|q3
				link3,title3,q1|q2
				link4,title4,q1
				link5,title5,q1|q2
				`)

				r := httptest.NewRequest(http.MethodGet, "/suggest/import", payload)
				r.Header.Set("Content-Type", "text/csv")
				w := httptest.NewRecorder()

				post(w, r)

				res := &repo.Status{}
				e := json.Unmarshal(w.Body.Bytes(), &res)

				Expect(e).ShouldNot(HaveOccurred())
				Expect(res.Status).Should(Equal("ok"))

				r = httptest.NewRequest(http.MethodGet, "/suggest", nil)
				w = httptest.NewRecorder()
				getall(w, r)

				suggests := make([]repo.Suggest, 0)
				e = json.Unmarshal(w.Body.Bytes(), &suggests)
				Expect(e).ShouldNot(HaveOccurred())
				Expect(len(suggests)).Should(Equal(5))
			})
		})

		When("add one without args", func() {
			It("Success", func() {
				payload := &bytes.Buffer{}
				payload.WriteString("")

				r := httptest.NewRequest(http.MethodGet, "/suggest", payload)
				r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
				w := httptest.NewRecorder()

				post(w, r)

				res := &repo.Status{}
				e := json.Unmarshal(w.Body.Bytes(), &res)

				Expect(e).ShouldNot(HaveOccurred())
				Expect(res.Status).Should(Equal("no args"))

			})
		})
	})

})
