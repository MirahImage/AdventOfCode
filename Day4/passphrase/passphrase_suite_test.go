package passphrase_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestPassphrase(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Passphrase Suite")
}
