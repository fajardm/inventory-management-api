package helpers

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"os"
	"testing"
	"github.com/fajardm/inventories/src/commons/helpers"
)

func TestInventories(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Database Test Suite")
}

var _ = Describe("Test Database", func() {
	var _ = AfterEach(func() {
		os.Remove(helpers.DB_PATH)
	})

	Describe("Test func SetDatabase()", func() {
		It("should error nil", func() {
			_, err := helpers.SetDatabase()
			Expect(err).To(BeNil())
		})
	})

	Describe("Test func GetDatabase()", func() {
		It("should get database not nil", func() {
			Expect(helpers.GetDatabase()).To(Not(BeNil()))
		})

		It("should ping database properly", func() {
			Expect(helpers.GetDatabase().DB().Ping()).To(BeNil())
		})
	})

	Describe("Test close database connection", func() {
		It("should closed properly", func() {
			helpers.GetDatabase().Close();
			Expect(helpers.GetDatabase().DB().Ping()).To(Not(BeNil()))
		})
	})
})
