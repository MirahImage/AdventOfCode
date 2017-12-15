package row_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestRow(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Row Suite")
}
