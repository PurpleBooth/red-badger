package red_badger_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestRedBadger(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "RedBadger Suite")
}
