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
)

func TestCreateProduct(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CreateProduct E2E Test Suite")
}

var _ = Describe("Test Create Product", func() {
	var router *gin.Engine

	var _ = BeforeEach(func() {
		router = boot.SetupMain()
	})

	It("should return error validation with empty payload", func() {
		payload := map[string]interface{}{
			"SKU": "sku",
		}
		body, _ := json.Marshal(payload)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/products", bytes.NewReader(body))
		req.Header.Add("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		Expect(w.Code).To(Equal(http.StatusBadRequest))
	})

	It("should not return error validation", func() {
		payload := map[string]interface{}{
			"SKU":   "sku",
			"Name":  "pc",
			"Stock": 1,
		}
		body, _ := json.Marshal(payload)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/products", bytes.NewReader(body))
		req.Header.Add("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		Expect(w.Code).To(Equal(http.StatusCreated))
	})
})
