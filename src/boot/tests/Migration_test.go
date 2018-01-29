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

var _ = Describe("Test Migration", func() {
	var _ = BeforeEach(func() {
		helpers.SetDatabase()
	})

	var _ = AfterEach(func() {
		os.Remove(helpers.DB_PATH)
	})

	Describe("Test func Migration", func() {
		It("should created table", func() {
			db := helpers.GetDatabase()
			boot.Migration()

			Expect(db.HasTable("products")).To(Equal(true))
		})
	})
})
