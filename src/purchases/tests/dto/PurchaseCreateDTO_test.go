package dto

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
	"github.com/fajardm/inventories/src/commons/helpers"
	"github.com/fajardm/inventories/src/purchases/dto"
)

func TestPurchaseCreateDTO(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "PurchaseCreateDTO Test Suite")
}

var _ = BeforeSuite(func() {
	helpers.SetValidator()
})

var _ = Describe("Test ProductCreateDTO", func() {
	It("should return error validation with empty payload", func() {
		dto := &dto.PurchaseCreateDTO{}
		err := helpers.GetValidator().Struct(dto)

		Expect(err).To(Not(BeNil()))
	})

	It("should not return error validation", func() {
		dto := &dto.PurchaseCreateDTO{
			OrderQuantity:  1,
			NumberReceived: 1,
			Price:          1,
			TotalPrice:     1,
			Note:           "note",
			ProductId:      "108291289",
		}
		err := helpers.GetValidator().Struct(dto)

		Expect(err).To(BeNil())
	})
})
