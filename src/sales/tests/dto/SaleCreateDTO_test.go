package dto

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
	"github.com/fajardm/inventories/src/commons/helpers"
	"github.com/fajardm/inventories/src/sales/dto"
)

func TestSaleCreateDTO(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TestSaleCreateDTO Test Suite")
}

var _ = Describe("Test TestSaleCreateDTO", func() {
	var _ = BeforeSuite(func() {
		helpers.SetValidator()
	})

	It("should return error validation with empty payload", func() {
		dto := &dto.SaleCreateDTO{}
		err := helpers.GetValidator().Struct(dto)

		Expect(err).To(Not(BeNil()))
	})

	It("should not return error validation", func() {
		dto := &dto.SaleCreateDTO{
			NumberShipped: 1,
			Price:         1,
			TotalPrice:    1,
			Note:          "note",
			ProductId:     "108291289",
		}
		err := helpers.GetValidator().Struct(dto)

		Expect(err).To(BeNil())
	})
})
