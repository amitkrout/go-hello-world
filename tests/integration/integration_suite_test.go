package integration_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/amitkrout/go-hello-world/tests/helper"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestIntegration(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Integration Suite")
}

var _ = BeforeSuite(func() {
	appURL := "http://localhost:8081/helloworld"
	resp, err := http.Get(appURL)
	Expect(err).ShouldNot(HaveOccurred())
	defer resp.Body.Close()
	Expect(helper.HttpWaitForStatus(appURL, 10, 1, 200)).Should(BeTrue())
	Expect(resp).Should(HaveHTTPStatus(http.StatusOK))
	fmt.Println("Test pre cecks successfully completed for running integration test")
})
