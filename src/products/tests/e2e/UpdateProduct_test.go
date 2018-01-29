package e2e

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
	"net/http"
	"net/http/httptest"
	"github.com/fajardm/inventories/src/boot"
	"github.com/gin-gonic/gin"
	"encoding/json"
	"bytes"
	"github.com/fajardm/inventories/src/commons/helpers"
	"github.com/fajardm/inventories/src/products/domain"
	"os"
)

func TestUpdateProduct(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "tUpdateProduct E2E Test Suite")
}

var _ = Describe("Test Update Product", func() {
	var router *gin.Engine

	var _ = BeforeEach(func() {
		router = boot.SetupMain()
	})

	var _ = AfterEach(func() {
		os.Remove(helpers.DB_PATH)
	})

	It("should return error validation with empty stock", func() {
		payload := map[string]interface{}{
			"SKU": "sku",
		}
		body, _ := json.Marshal(payload)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("PUT", "/products/id", bytes.NewReader(body))
		req.Header.Add("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		Expect(w.Code).To(Equal(http.StatusBadRequest))
	})

	It("should return not found", func() {
		payload := map[string]interface{}{
			"SKU":   "sku",
			"Name":  "pc",
			"stock": 1,
		}
		body, _ := json.Marshal(payload)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("PUT", "/products/id", bytes.NewReader(body))
		req.Header.Add("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		Expect(w.Code).To(Equal(http.StatusNotFound))
	})

	It("should return product", func() {
		product := domain.Product{SKU: "SKU", Stock: 1}
		helpers.GetDatabase().Create(&product)

		payload := map[string]interface{}{
			"SKU":   "sku",
			"Name":  "pc",
			"stock": 1,
		}
		body, _ := json.Marshal(payload)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("PUT", "/products/"+product.ID, bytes.NewReader(body))
		req.Header.Add("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		Expect(w.Code).To(Equal(http.StatusOK))
	})
})
