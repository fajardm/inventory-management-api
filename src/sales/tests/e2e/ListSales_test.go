package e2e

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
	"net/http"
	"net/http/httptest"
	"github.com/fajardm/inventories/src/boot"
	"github.com/gin-gonic/gin"
	"github.com/fajardm/inventories/src/commons/helpers"
	"github.com/fajardm/inventories/src/sales/domain"
	"os"
)

func TestListSales(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ListSales E2E Test Suite")
}

var _ = Describe("Test List Sales", func() {
	var router *gin.Engine

	var _ = BeforeEach(func() {
		router = boot.SetupMain()
	})

	var _ = AfterEach(func() {
		os.Remove(helpers.DB_PATH)
	})

	It("should return sales list", func() {
		purchase := domain.Sale{
			NumberShipped: 1,
			Price:         1,
			TotalPrice:    1,
			Note:          "note",
			ProductId:     "108291289",
		}
		helpers.GetDatabase().Create(&purchase)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/sales", nil)
		req.Header.Add("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		Expect(w.Code).To(Equal(http.StatusOK))
	})
})
