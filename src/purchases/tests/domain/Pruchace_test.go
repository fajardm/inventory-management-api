package domain

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
	"github.com/fajardm/inventories/src/purchases/domain"
	domain2 "github.com/fajardm/inventories/src/products/domain"
)

func TestPurchases(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Purchases Test Suite")
}

var _ = Describe("Test Purchases", func() {
	Describe("Test created purchase model", func() {
		It("should created purchase properly", func() {
			model := domain.Purchase{
				OrderQuantity:  1,
				NumberReceived: 1,
				Price:          1,
				TotalPrice:     1,
				Note:           "note",
				Product:        domain2.Product{SKU: "SKU", Stock: 0},
			}

			Expect(model.OrderQuantity).To(Equal(1))
			Expect(model.NumberReceived).To(Equal(1))
		})
	})
})
