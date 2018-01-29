package services

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
	"github.com/fajardm/inventories/src/commons/helpers"
	"github.com/fajardm/inventories/src/boot"
	"github.com/fajardm/inventories/src/products/services"
	"os"
	"github.com/fajardm/inventories/src/products/domain"
)

func TestProductDeleteService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ProductDeleteService Test Suite")
}

var _ = Describe("Test ProductDeleteService", func() {
	var _ = BeforeEach(func() {
		helpers.SetDatabase()
		boot.Migration()
		helpers.SetValidator()
	})

	var _ = AfterEach(func() {
		os.Remove(helpers.DB_PATH)
	})

	Describe("Test func DeleteProduct", func() {
		It("should return not found", func() {
			service := services.ProductDeleteService{}
			res := service.DeleteProduct("not-found")

			Expect(res.Error).To(Equal(""))
			Expect(res.IsFound).To(BeFalse())
		})

		It("should return product", func() {
			product := domain.Product{SKU: "SKU", Stock: 1}
			helpers.GetDatabase().Create(&product)

			service := services.ProductDeleteService{}
			res := service.DeleteProduct(product.ID)

			Expect(res.Error).To(Equal(""))
			Expect(res.IsFound).To(BeTrue())
		})
	})
})
