package red_badger_test

import (
	. "github.com/purplebooth/red-badger"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("robot", func() {
	Context("Initialisation", func() {

		It("should be creatable from co-ordinate, and starting direction and instructions", func() {
			actual := NewRobot(
				&Coordinate{X: 0, Y: 0},
				&Direction{Direction: "N"},
				[]Instruction{},
			)
			Expect(actual).ToNot(BeNil())
		})
	})
})
