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

var dbPath = "../../../../inventory.db"

var _ = AfterSuite(func() {
	os.Remove(dbPath)
})

var _ = Describe("Test Database", func() {
	Describe("Test func SetDatabase()", func() {
		It("should error nil", func() {
			_, err := helpers.SetDatabase()
			Expect(err).To(BeNil())
		})

		It("should error nil when open inventory.db", func() {
			_, err := os.Open(dbPath)
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
