package services

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
	"github.com/fajardm/inventories/src/commons/helpers"
	"os"
	"github.com/fajardm/inventories/src/boot"
	"github.com/fajardm/inventories/src/sales/dto"
	"github.com/fajardm/inventories/src/sales/services"
	"github.com/fajardm/inventories/src/products/domain"
)

func TestSaleCreateService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "SaleCreateService Test Suite")
}

var _ = Describe("Test SaleCreateService", func() {
	var _ = BeforeEach(func() {
		helpers.SetDatabase()
		boot.Migration()
		helpers.SetValidator()
	})

	var _ = AfterEach(func() {
		os.Remove(helpers.DB_PATH)
	})

	Describe("Test func CreateSale", func() {
		It("should return error validation", func() {
			dto := &dto.SaleCreateDTO{}

			service := services.SaleCreateService{}
			res := service.CreateSale(dto)

			Expect(res.ErrorValidation).To(Not(BeNil()))
			Expect(res.Error).To(Equal(""))
			Expect(res.Sale.ID).To(Equal(""))
		})

		It("should return entity error", func() {
			dto := &dto.SaleCreateDTO{
				NumberShipped: 1,
				Price:         1,
				TotalPrice:    1,
				Note:          "note",
				Product:       "1",
			}

			service := services.SaleCreateService{}
			res := service.CreateSale(dto)

			Expect(res.ErrorValidation).To(BeNil())
			Expect(res.Error).To(Equal(""))
			Expect(res.EntityErrorMessage).To(Not(BeEmpty()))
		})

		It("should return sale", func() {
			product := domain.Product{SKU: "SKU", Name: "pc", Stock: 1}
			helpers.GetDatabase().Create(&product)

			dto := &dto.SaleCreateDTO{
				NumberShipped: 1,
				Price:         1,
				TotalPrice:    1,
				Note:          "note",
				Product:       product.ID,
			}

			service := services.SaleCreateService{}
			res := service.CreateSale(dto)

			Expect(res.ErrorValidation).To(BeNil())
			Expect(res.Error).To(Equal(""))
			Expect(res.Sale.ID).To(Not(BeNil()))
		})
	})
})
