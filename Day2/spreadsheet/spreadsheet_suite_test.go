package spreadsheet_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSpreadsheet(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Spreadsheet Suite")
}
