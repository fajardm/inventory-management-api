package services

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
	"github.com/fajardm/inventories/src/products/services"
	"github.com/fajardm/inventories/src/products/dto"
	"github.com/fajardm/inventories/src/commons/helpers"
	"os"
	"github.com/fajardm/inventories/src/boot"
)

func TestProductCreateService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ProductCreateService Test Suite")
}

var _ = Describe("Test ProductCreateService", func() {
	var _ = BeforeEach(func() {
		helpers.SetDatabase()
		boot.Migration()
		helpers.SetValidator()
	})

	var _ = AfterEach(func() {
		dbPath := "../../../../inventory.db"
		os.Remove(dbPath)
	})

	Describe("Test func CreateProduct", func() {
		It("should return error validation", func() {
			dto := &dto.ProductCreateDTO{}

			service := services.ProductCreateService{}
			res := service.CreateProduct(dto)

			Expect(res.ErrorValidation).To(Not(BeNil()))
			Expect(res.Error).To(Equal(""))
			Expect(res.Product.ID).To(Equal(""))
		})

		It("should return product", func() {
			dto := &dto.ProductCreateDTO{SKU: "SKU", Name: "PC", Stock: 1}

			service := services.ProductCreateService{}
			res := service.CreateProduct(dto)

			Expect(res.ErrorValidation).To(BeNil())
			Expect(res.Error).To(Equal(""))
			Expect(res.Product.ID).To(Not(BeNil()))
		})
	})
})
