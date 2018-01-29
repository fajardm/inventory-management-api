package dto

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
	"github.com/fajardm/inventories/src/products/dto"
	"github.com/fajardm/inventories/src/commons/helpers"
)

func TestInventories(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ProductCreateDTO Test Suite")
}

var _ = BeforeSuite(func() {
	helpers.SetValidator()
})

var _ = Describe("Test ProductCreateDTO", func() {
	It("should return error validation with empty payload", func() {
		dto := &dto.ProductCreateDTO{}
		err := helpers.GetValidator().Struct(dto)

		Expect(err).To(Not(BeNil()))
	})

	It("should not return error validation with stock under 0", func() {
		dto := &dto.ProductCreateDTO{SKU: "SKU", Name: "PC", Stock: -1}
		err := helpers.GetValidator().Struct(dto)

		Expect(err).To(Not(BeNil()))
	})

	It("should not return error validation", func() {
		dto := &dto.ProductCreateDTO{SKU: "SKU", Name: "PC", Stock: 1}
		err := helpers.GetValidator().Struct(dto)

		Expect(err).To(BeNil())
	})
})
