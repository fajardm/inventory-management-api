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

func TestProductShowService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ProductShowService Test Suite")
}

var _ = Describe("Test ProductShowService", func() {
	var _ = BeforeEach(func() {
		helpers.SetDatabase()
		boot.Migration()
		helpers.SetValidator()
	})

	var _ = AfterEach(func() {
		os.Remove(helpers.DB_PATH)
	})

	Describe("Test func ShowProduct", func() {
		It("should return not found", func() {
			service := services.ProductShowService{}
			res := service.ShowProduct("not-found")

			Expect(res.Error).To(Equal(""))
			Expect(res.IsFound).To(BeFalse())
		})

		It("should return product", func() {
			product := domain.Product{SKU: "SKU", Name: "pc", Stock: 1}
			helpers.GetDatabase().Create(&product)

			service := services.ProductShowService{}
			res := service.ShowProduct(product.ID)

			Expect(res.Error).To(Equal(""))
			Expect(res.IsFound).To(BeTrue())
		})
	})
})
