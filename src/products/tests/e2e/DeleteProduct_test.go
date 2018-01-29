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
	"github.com/fajardm/inventories/src/products/domain"
	"os"
)

func TestDeleteProduct(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "DeleteProduct E2E Test Suite")
}

var _ = Describe("Test Delete Product", func() {
	var router *gin.Engine

	var _ = BeforeEach(func() {
		router = boot.SetupMain()
	})

	var _ = AfterEach(func() {
		os.Remove(helpers.DB_PATH)
	})

	It("should return not found", func() {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("DELETE", "/products/1", nil)
		req.Header.Add("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		Expect(w.Code).To(Equal(http.StatusNotFound))
	})

	It("should return success", func() {
		product := domain.Product{SKU: "SKU", Stock: 1}
		helpers.GetDatabase().Create(&product)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("DELETE", "/products/"+product.ID, nil)
		req.Header.Add("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		Expect(w.Code).To(Equal(http.StatusOK))
	})
})
