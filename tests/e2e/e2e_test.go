package e2e

import (
	"io/ioutil"
	"net/http"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("End to end scenario", func() {

	It("should able to access the URL", func() {
		appURL := "http://localhost:8081/helloworld"
		resp, err := http.Get(appURL)
		Expect(err).ShouldNot(HaveOccurred())
		defer resp.Body.Close()
		Expect(resp).Should(HaveHTTPStatus(http.StatusOK))
	})

	It("won't able to access invalid URL payload", func() {
		appURL := "http://localhost:8081/helloworld-wrong-payload"
		resp, err := http.Get(appURL)
		Expect(err).ShouldNot(HaveOccurred())
		defer resp.Body.Close()
		Expect(resp).Should(HaveHTTPStatus("404 Not Found"))
	})

	It("should display the correct content - Hello, World!", func() {
		appURL := "http://localhost:8081/helloworld"
		resp, err := http.Get(appURL)
		Expect(err).ShouldNot(HaveOccurred())
		defer resp.Body.Close()
		Expect(resp).Should(HaveHTTPStatus(http.StatusOK))
		content, err := ioutil.ReadAll(resp.Body)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(content).To(ContainSubstring("Hello, World!"))
	})
})
