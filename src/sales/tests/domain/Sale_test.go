package domain

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
	"github.com/fajardm/inventories/src/sales/domain"
	domain2 "github.com/fajardm/inventories/src/products/domain"
)

func TestSales(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Sales Test Suite")
}

var _ = Describe("Test Sales", func() {
	Describe("Test created sale model", func() {
		It("should created sale properly", func() {
			model := domain.Sale{
				NumberShipped: 1,
				Price:         1,
				TotalPrice:    1,
				Note:          "note",
				Product:       domain2.Product{SKU: "SKU", Stock: 0},
			}

			Expect(model.NumberShipped).To(Equal(1))
			Expect(model.Price).To(Equal(float32(1)))
		})
	})
})
