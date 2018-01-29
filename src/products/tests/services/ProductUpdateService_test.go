package services

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
	"github.com/fajardm/inventories/src/commons/helpers"
	"github.com/fajardm/inventories/src/boot"
	"github.com/fajardm/inventories/src/products/services"
	"os"
	"github.com/fajardm/inventories/src/products/dto"
	"github.com/fajardm/inventories/src/products/domain"
)

func TestProductUpdateService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ProductUpdateService Test Suite")
}

var _ = Describe("Test ProductUpdateService", func() {
	var _ = BeforeEach(func() {
		helpers.SetDatabase()
		boot.Migration()
		helpers.SetValidator()
	})

	var _ = AfterEach(func() {
		os.Remove(helpers.DB_PATH)
	})

	Describe("Test func UpdateProduct", func() {
		It("should return validation error", func() {
			dto := &dto.ProductUpdateDTO{}

			service := services.ProductUpdateService{}
			res := service.UpdateProduct("id", dto)

			Expect(res.ErrorValidation).To(Not(BeNil()))
			Expect(res.Error).To(Equal(""))
			Expect(res.IsFound).To(BeFalse())
		})

		It("should return not found", func() {
			dto := &dto.ProductUpdateDTO{SKU: "SKU", Name: "PC Update", Stock: 1}

			service := services.ProductUpdateService{}
			res := service.UpdateProduct("id", dto)

			Expect(res.ErrorValidation).To(BeNil())
			Expect(res.Error).To(Equal(""))
			Expect(res.IsFound).To(BeFalse())
		})

		It("should return product", func() {
			product := domain.Product{SKU: "SKU", Name: "PC", Stock: 1}
			helpers.GetDatabase().Create(&product)

			dto := &dto.ProductUpdateDTO{SKU: "SKU", Name: "PC Update", Stock: 1}

			service := services.ProductUpdateService{}
			res := service.UpdateProduct(product.ID, dto)

			Expect(res.ErrorValidation).To(BeNil())
			Expect(res.Error).To(Equal(""))
			Expect(res.IsFound).To(BeTrue())
		})
	})
})
