package integration

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
)

var _ = Describe("Feature-1 integration test", func() {

	var _ = BeforeEach(func() {
		fmt.Println("This is run before every Spec (It)")
	})

	var _ = AfterEach(func() {
		fmt.Println("This is run after every Spec (It)")
	})

	When("running feature-1 integration test", func() {
		It("should check frist functionalities", func() {
			fmt.Println("Write logic that checks the first functionalities")
		})

		It("should check frist functionalities", func() {
			fmt.Println("Write logic that checks the first functionalities")
		})

	})
})
