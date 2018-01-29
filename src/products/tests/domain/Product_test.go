package domain

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
	"github.com/fajardm/inventories/src/products/domain"
)

func TestInventories(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ProductModel Test Suite")
}

var _ = Describe("Test ProductModel", func() {
	Describe("Test created product model", func() {
		It("should created product properly", func() {
			model := domain.Product{SKU: "SKU", Name: "PC", Stock: 1}

			Expect(model.SKU).To(Equal("SKU"))
			Expect(model.Name).To(Equal("PC"))
			Expect(model.Stock).To(Equal(1))
		})
	})
})
