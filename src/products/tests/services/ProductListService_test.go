package services

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
	"github.com/fajardm/inventories/src/commons/helpers"
	"github.com/fajardm/inventories/src/boot"
	"os"
	"github.com/fajardm/inventories/src/products/domain"
	"github.com/fajardm/inventories/src/products/services"
)

func TestProductListService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ProductListService Test Suite")
}

var _ = Describe("Test ProductListService", func() {
	var _ = BeforeEach(func() {
		helpers.SetDatabase()
		boot.Migration()
		helpers.SetValidator()
	})

	var _ = AfterEach(func() {
		dbPath := "../../../../inventory.db"
		os.Remove(dbPath)
	})

	Describe("Test func ListProduct", func() {
		It("should return product list", func() {
			product := domain.Product{SKU: "SKU", Stock: 1}
			helpers.GetDatabase().Create(&product)

			service := services.ProductListService{}
			res := service.ListProduct()

			Expect(res.Error).To(Equal(""))
		})
	})
})
