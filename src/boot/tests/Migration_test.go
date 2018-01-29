package tests

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
	"github.com/fajardm/inventories/src/boot"
	"github.com/fajardm/inventories/src/commons/helpers"
	"os"
)

func TestInventories(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Migration Test Suite")
}

var _ = BeforeSuite(func() {
	helpers.SetDatabase()
})

var dbPath = "../../../inventory.db"

var _ = AfterSuite(func() {
	os.Remove(dbPath)
})

var _ = Describe("Test Migration", func() {
	Describe("Test func Migration", func() {
		It("should created table", func() {
			db := helpers.GetDatabase()
			boot.Migration()

			Expect(db.HasTable("products")).To(Equal(true))
		})
	})
})
