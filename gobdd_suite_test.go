package gobdd_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGobdd(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gobdd Suite")
}
