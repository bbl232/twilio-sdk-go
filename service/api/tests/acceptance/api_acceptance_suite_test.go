// +build acceptance autopilot_acceptance

package acceptance

import (
	"fmt"
	"os"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// Prerequisites

// 1) Twilio Account SID set as an environment variable - TWILIO_ACCOUNT_SID
// 2) Twilio Auth Token set as an environment variable - TWILIO_AUTH_TOKEN
// 3) Twilio Phone Number set as an environment variable - TWILIO_PHONE_NUMBER
// 4) Phone Number to send messages to set as an environment variable - DESTINATION_PHONE_NUMBER

func TestAcceptance(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Autopilot Acceptance Test Suite")
}

var _ = BeforeSuite(func() {
	variables := []string{
		"TWILIO_ACCOUNT_SID",
		"TWILIO_AUTH_TOKEN",
		"TWILIO_PHONE_NUMBER",
		"DESTINATION_PHONE_NUMBER",
	}

	for _, variable := range variables {
		if value := os.Getenv(variable); value == "" {
			Fail(fmt.Sprintf("`%s` are required for running acceptance tests", variable))
		}
	}
})
