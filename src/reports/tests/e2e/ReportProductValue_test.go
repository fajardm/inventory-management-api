package domain

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
	"github.com/fajardm/inventories/src/commons/helpers"
	"os"
	"github.com/fajardm/inventories/src/boot"
	"encoding/json"
	"net/http/httptest"
	"net/http"
	"bytes"
	"github.com/gin-gonic/gin"
)

func TestProductValueController(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ProductValueController Test Suite")
}

var _ = Describe("Test Product Value", func() {
	var router *gin.Engine

	var _ = BeforeEach(func() {
		helpers.SetDatabase()
		boot.Migration()
		boot.Seeder()
		router = boot.SetupMain()
	})

	var _ = AfterEach(func() {
		dbPath := "../../../../inventory.db"
		os.Remove(dbPath)
	})

	It("should return reports", func() {
		payload := map[string]interface{}{}
		body, _ := json.Marshal(payload)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/reports/product_values", bytes.NewReader(body))
		req.Header.Add("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		Expect(w.Code).To(Equal(http.StatusOK))
	})
})
