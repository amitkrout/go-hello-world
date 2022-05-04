package e2e

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
)

var _ = Describe("End to end scenario - 1", func() {

	var _ = BeforeEach(func() {
		fmt.Println("This is run before every Spec (It)")
	})

	var _ = AfterEach(func() {
		fmt.Println("This is run after every Spec (It)")
	})

	It("Execute step - 1", func() {
		fmt.Println("Write logic to execute step - 1")
	})

	It("Execute step - 2", func() {
		fmt.Println("Write logic to execute step - 2")
	})
})
