package programData_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestProgramData(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ProgramData Suite")
}
