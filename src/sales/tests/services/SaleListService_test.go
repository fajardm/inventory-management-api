package services

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
	"github.com/fajardm/inventories/src/commons/helpers"
	"github.com/fajardm/inventories/src/boot"
	"os"
	"github.com/fajardm/inventories/src/sales/domain"
	"github.com/fajardm/inventories/src/sales/services"
)

func TestPurchaseListService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "PurchaseListService Test Suite")
}

var _ = Describe("Test PurchaseListService", func() {
	var _ = BeforeEach(func() {
		helpers.SetDatabase()
		boot.Migration()
		helpers.SetValidator()
	})

	var _ = AfterEach(func() {
		os.Remove(helpers.DB_PATH)
	})

	Describe("Test func ListPurchases", func() {
		It("should return purchase list", func() {
			purchase := domain.Sale{
				NumberShipped: 1,
				Price:         1,
				TotalPrice:    1,
				Note:          "note",
				ProductId:     "108291289",
			}
			helpers.GetDatabase().Create(&purchase)

			service := services.SaleListService{}
			res := service.ListSales()

			Expect(res.Error).To(Equal(""))
		})
	})
})