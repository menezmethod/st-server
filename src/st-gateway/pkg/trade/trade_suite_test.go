package trade_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestTrade(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Trade Suite")
}
