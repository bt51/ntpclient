package ntpclient_test

import (
    . "github.com/bt51/ntpclient"
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
)

var _ = Describe("Ntp", func() {
    Context("Current time", func() {
        It("should get current time", func() {
            now, err := GetNetworkTime("0.pool.ntp.org", 123)
            Expect(err).Should(BeNil())
            Expect(now).ShouldNot(BeNil())
        })
    })
})
