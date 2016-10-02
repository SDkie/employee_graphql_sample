package gq_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGq(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gq Suite")
}
