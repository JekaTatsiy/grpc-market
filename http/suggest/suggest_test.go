package suggest_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"encoding/json"

	"flag"

	server "github.com/JekaTatsiy/grpc-market/http/server"
	repo "github.com/JekaTatsiy/grpc-market/http/suggest"
	"github.com/gorilla/mux"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/format"
)

var searchaddr = flag.String("s", "grpc-search:1000", "adres grpc-search service")

func TestHTTPSuggest(t *testing.T) {
	format.MaxLength = 0

	RegisterFailHandler(Fail)
	RunSpecs(t, "HTTPSuggest")
}

var _ = Describe("HTTPSuggest", func() {

	suggest1 := url.Values{
		"id":    []string{"1"},
		"link":  []string{"abc"},
		"title": []string{"NAME"},
		"query": []string{"a", "b", "c"}}
	suggest2 := url.Values{
		"id":    []string{"2"},
		"link":  []string{"xyz"},
		"title": []string{"ITEM"},
		"query": []string{"x", "y", "z"}}

	flag.Parse()
	//g := server.NewGrpcClient(*searchaddr)
	g := server.NewGrpcClient(":1000")

	getall := repo.GetAll(g)
	get := repo.Get(g)
	post := repo.Post(g)
	imp := repo.Import(g)
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
				r := httptest.NewRequest(http.MethodGet, "/suggest", nil)
				r.Header.Set("Content-Type", "maultipart/form-data")
				r.Form = suggest1
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
				r := httptest.NewRequest(http.MethodGet, "/suggest", nil)
				r.Header.Set("Content-Type", "maultipart/form-data")
				r.Form = suggest1
				w := httptest.NewRecorder()

				post(w, r)

				res := &repo.Status{}
				e := json.Unmarshal(w.Body.Bytes(), &res)
				Expect(e).ShouldNot(HaveOccurred())
				Expect(res.Status).Should(Equal("ok"))

				r = httptest.NewRequest(http.MethodGet, "/suggest/1", nil)
				r = mux.SetURLVars(r, map[string]string{"id": "1"})

				r.Header.Set("Content-Type", "maultipart/form-data")
				w = httptest.NewRecorder()

				get(w, r)

				suggest := &repo.Suggest{}
				e = json.Unmarshal(w.Body.Bytes(), &suggest)

				Expect(e).ShouldNot(HaveOccurred())

				Expect(suggest.LinkUrl).Should(Equal("abc"))
				Expect(suggest.Title).Should(Equal("NAME"))
				Expect(len(suggest.Queries)).Should(Equal(3))
				Expect(suggest.Queries).Should(ContainElements("a", "b", "c"))
			})
		})

		When("get all", func() {
			It("Success", func() {
				r := httptest.NewRequest(http.MethodGet, "/suggest", nil)
				r.Header.Set("Content-Type", "maultipart/form-data")
				r.Form = suggest1
				w := httptest.NewRecorder()
				post(w, r)
				res := &repo.Status{}
				e := json.Unmarshal(w.Body.Bytes(), &res)
				Expect(e).ShouldNot(HaveOccurred())
				Expect(res.Status).Should(Equal("ok"))

				r = httptest.NewRequest(http.MethodGet, "/suggest", nil)
				r.Header.Set("Content-Type", "maultipart/form-data")
				r.Form = suggest2

				w = httptest.NewRecorder()
				post(w, r)
				res = &repo.Status{}
				e = json.Unmarshal(w.Body.Bytes(), &res)
				Expect(e).ShouldNot(HaveOccurred())
				Expect(res.Status).Should(Equal("ok"))

				time.Sleep(1000 * time.Millisecond)

				r = httptest.NewRequest(http.MethodGet, "/suggest", nil)
				w = httptest.NewRecorder()
				getall(w, r)

				suggests := make([]repo.Suggest, 0)
				e = json.Unmarshal(w.Body.Bytes(), &suggests)
				Expect(e).ShouldNot(HaveOccurred())
				Expect(len(suggests)).Should(Equal(2))

				Expect(suggests[1].LinkUrl).Should(Equal("xyz"))
				Expect(suggests[1].Title).Should(Equal("ITEM"))
				Expect(len(suggests[1].Queries)).Should(Equal(3))
				Expect(suggests[1].Queries).Should(ContainElements("x", "y", "z"))

			})
		})

		When("delete", func() {
			It("Success", func() {
				payload := &bytes.Buffer{}

				r := httptest.NewRequest(http.MethodGet, "/suggest", payload)
				r.Header.Set("Content-Type", "maultipart/form-data")
				r.Form = suggest1

				w := httptest.NewRecorder()

				post(w, r)

				payload = &bytes.Buffer{}

				r = httptest.NewRequest(http.MethodGet, "/suggest", payload)
				r.Header.Set("Content-Type", "maultipart/form-data")
				r.Form = suggest2

				w = httptest.NewRecorder()

				post(w, r)

				res := &repo.Status{}
				e := json.Unmarshal(w.Body.Bytes(), &res)
				Expect(e).ShouldNot(HaveOccurred())
				Expect(res.Status).Should(Equal("ok"))

				//delete one
				r = httptest.NewRequest(http.MethodGet, "/suggest/1", nil)
				r = mux.SetURLVars(r, map[string]string{"id": "1"})
				r.Header.Set("Content-Type", "maultipart/form-data")
				w = httptest.NewRecorder()

				deleteone(w, r)

				res = &repo.Status{}
				e = json.Unmarshal(w.Body.Bytes(), &res)
				Expect(e).ShouldNot(HaveOccurred())
				Expect(res.Status).Should(Equal("ok"))

				r = httptest.NewRequest(http.MethodGet, "/suggest", nil)
				r.Header.Set("Content-Type", "maultipart/form-data")
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
				link5,title5,q1|q2`)

				r := httptest.NewRequest(http.MethodGet, "/suggest/import", payload)
				r.Header.Set("Content-Type", "text/csv")
				w := httptest.NewRecorder()

				imp(w, r)

				res := &repo.Status{}
				e := json.Unmarshal(w.Body.Bytes(), &res)

				Expect(e).ShouldNot(HaveOccurred())
				Expect(res.Status).Should(Equal("ok"))

				time.Sleep(1000 * time.Millisecond)

				r = httptest.NewRequest(http.MethodGet, "/suggest", nil)
				w = httptest.NewRecorder()
				getall(w, r)

				suggests := make([]*repo.Suggest, 0)
				e = json.Unmarshal(w.Body.Bytes(), &suggests)
				Expect(e).ShouldNot(HaveOccurred())
				Expect(len(suggests)).Should(Equal(5))
			})
		})

		When("add one without args", func() {
			It("Success", func() {

				r := httptest.NewRequest(http.MethodGet, "/suggest", nil)
				r.Header.Set("Content-Type", "maultipart/form-data")
				w := httptest.NewRecorder()

				post(w, r)

				res := &repo.Status{}
				e := json.Unmarshal(w.Body.Bytes(), &res)

				Expect(e).ShouldNot(HaveOccurred())
				Expect(res.Status).Should(Equal("expect id"))

			})
		})
	})

})
