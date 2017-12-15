package captcha_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCaptcha(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Captcha Suite")
}
