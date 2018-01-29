package domain

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
	"github.com/fajardm/inventories/src/commons/helpers"
	"os"
	"github.com/fajardm/inventories/src/boot"
	"github.com/fajardm/inventories/src/reports/services"
)

func TestProductSaleService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ProductSaleService Test Suite")
}

var _ = Describe("Test ProductSaleService", func() {
	var _ = AfterEach(func() {
		os.Remove(helpers.DB_PATH)
	})

	It("should return reports", func() {
		helpers.SetDatabase()
		boot.Migration()
		boot.Seeder()

		service := services.ReportProductSaleService{}
		res := service.ReportProductSale()

		Expect(res.Error).To(Equal(""))
	})
})
