package domain

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
	"time"
	"github.com/fajardm/inventories/src/commons/domain"
)

func TestInventories(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Model Test Suite")
}

var _ = Describe("Test Model", func() {
	Describe("Test created model", func() {
		It("should created model properly", func() {
			currentTime := time.Now()
			model := domain.Model{ID: "uuid", CreatedAt: currentTime, UpdatedAt: currentTime}

			Expect(model.ID).To(Equal("uuid"))
			Expect(model.CreatedAt).To(Equal(currentTime))
			Expect(model.UpdatedAt).To(Equal(currentTime))
			Expect(model.DeletedAt).To(BeNil())
		})
	})
})
