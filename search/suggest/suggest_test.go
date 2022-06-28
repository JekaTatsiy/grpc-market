package suggest_test

import (
	repo "github.com/JekaTatsiy/grpc-market/search/suggest"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/format"
	"testing"
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
		When("add one", func() {
			It("Success", func() {

			})
		})
	})
})
