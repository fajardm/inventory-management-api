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

func TestCreateSale(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CreateSale E2E Test Suite")
}

var _ = Describe("Test Create Sale", func() {
	var router *gin.Engine

	var _ = BeforeEach(func() {
		router = boot.SetupMain()
	})

	var _ = AfterEach(func() {
		os.Remove(helpers.DB_PATH)
	})

	It("should return error validation with empty payload", func() {
		payload := map[string]interface{}{}
		body, _ := json.Marshal(payload)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/sales", bytes.NewReader(body))
		req.Header.Add("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		Expect(w.Code).To(Equal(http.StatusBadRequest))
	})

	It("should return entity error", func() {
		payload := map[string]interface{}{
			"NumberShipped": 1,
			"Price":         1,
			"TotalPrice":    1,
			"Note":          "note",
			"ProductId":     "1",
		}
		body, _ := json.Marshal(payload)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/sales", bytes.NewReader(body))
		req.Header.Add("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		Expect(w.Code).To(Equal(http.StatusUnprocessableEntity))
	})

	It("should not return error", func() {
		product := domain.Product{SKU: "SKU", Name: "pc", Stock: 1}
		helpers.GetDatabase().Create(&product)

		payload := map[string]interface{}{
			"NumberShipped": 1,
			"Price":         1,
			"TotalPrice":    1,
			"Note":          "note",
			"ProductId":     product.ID,
		}
		body, _ := json.Marshal(payload)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/sales", bytes.NewReader(body))
		req.Header.Add("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		Expect(w.Code).To(Equal(http.StatusCreated))
	})
})
