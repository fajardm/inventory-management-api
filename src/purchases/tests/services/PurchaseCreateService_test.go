package services

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
	"github.com/fajardm/inventories/src/commons/helpers"
	"os"
	"github.com/fajardm/inventories/src/boot"
	"github.com/fajardm/inventories/src/purchases/dto"
	"github.com/fajardm/inventories/src/purchases/services"
	"github.com/fajardm/inventories/src/products/domain"
)

func TestPurchaseCreateService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "PurchaseCreateService Test Suite")
}

var _ = Describe("Test PurchaseCreateService", func() {
	var _ = BeforeEach(func() {
		helpers.SetDatabase()
		boot.Migration()
		helpers.SetValidator()
	})

	var _ = AfterEach(func() {
		dbPath := "../../../../inventory.db"
		os.Remove(dbPath)
	})

	Describe("Test func CreatePurchase", func() {
		It("should return error validation", func() {
			dto := &dto.PurchaseCreateDTO{}

			service := services.PurchaseCreateService{}
			res := service.CreatePurchase(dto)

			Expect(res.ErrorValidation).To(Not(BeNil()))
			Expect(res.Error).To(Equal(""))
			Expect(res.Purchase.ID).To(Equal(""))
		})

		It("should return entity error", func() {
			dto := &dto.PurchaseCreateDTO{
				OrderQuantity:  1,
				NumberReceived: 1,
				Price:          1,
				TotalPrice:     1,
				Note:           "note",
				Product:        "1",
			}

			service := services.PurchaseCreateService{}
			res := service.CreatePurchase(dto)

			Expect(res.ErrorValidation).To(BeNil())
			Expect(res.Error).To(Equal(""))
			Expect(res.EntityErrorMessage).To(Not(BeEmpty()))
		})

		It("should return purchase", func() {
			product := domain.Product{SKU: "SKU", Name: "pc", Stock: 1}
			helpers.GetDatabase().Create(&product)

			dto := &dto.PurchaseCreateDTO{
				OrderQuantity:  1,
				NumberReceived: 1,
				Price:          1,
				TotalPrice:     1,
				Note:           "note",
				Product:        product.ID,
			}

			service := services.PurchaseCreateService{}
			res := service.CreatePurchase(dto)

			Expect(res.ErrorValidation).To(BeNil())
			Expect(res.Error).To(Equal(""))
			Expect(res.Purchase.ID).To(Not(BeNil()))
		})
	})
})
