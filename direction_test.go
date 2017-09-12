package red_badger_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/purplebooth/red-badger"
)

var _ = Describe("direction", func() {
	Context("Create from string combined with co-ordinate", func() {
		It("Ignores the co-ordinate part of the string", func() {
			actual, err := NewDirectionFromCombinedString("1 1 N")

			Expect(actual.Direction).To(Equal("N"), "direction")
			Expect(err).To(BeNil(), "Error")
		})
		It("Allows N", func() {
			actual, err := NewDirectionFromCombinedString("1 1 N")

			Expect(actual.Direction).To(Equal("N"), "direction")
			Expect(err).To(BeNil(), "Error")
		})
		It("Allows S", func() {
			actual, err := NewDirectionFromCombinedString("1 1 S")

			Expect(actual.Direction).To(Equal("S"), "direction")
			Expect(err).To(BeNil(), "Error")
		})
		It("Allows E", func() {
			actual, err := NewDirectionFromCombinedString("1 1 E")

			Expect(actual.Direction).To(Equal("E"), "direction")
			Expect(err).To(BeNil(), "Error")
		})
		It("Allows W", func() {
			actual, err := NewDirectionFromCombinedString("1 1 W")

			Expect(actual.Direction).To(Equal("W"), "direction")
			Expect(err).To(BeNil(), "Error")
		})
		It("Does not allow other directions", func() {
			actual, err := NewDirectionFromCombinedString("1 1 NW")

			Expect(actual).To(BeNil(), "direction")
			Expect(err).ToNot(BeNil(), "Error")
		})
		It("Totally invalid strings fail", func() {
			actual, err := NewDirectionFromCombinedString("unparsable")

			Expect(actual).To(BeNil(), "direction")
			Expect(err).ToNot(BeNil(), "Error")
		})
	})
	Context("Next direction after turn", func() {
		It("Anti-clockwise north follows east", func() {
			Expect(DirectionAntiClockwiseAfter(&Direction{Direction: "E"}).Direction).To(Equal("N"))
		})
		It("Anti-clockwise east follows south", func() {
			Expect(DirectionAntiClockwiseAfter(&Direction{Direction: "S"}).Direction).To(Equal("E"))
		})
		It("Anti-clockwise south follows west", func() {
			Expect(DirectionAntiClockwiseAfter(&Direction{Direction: "W"}).Direction).To(Equal("S"))
		})
		It("Anti-clockwise west follows north", func() {
			Expect(DirectionAntiClockwiseAfter(&Direction{Direction: "N"}).Direction).To(Equal("W"))
		})
		It("Clockwise north follows ", func() {
			Expect(DirectionClockwiseAfter(&Direction{Direction: "N"}).Direction).To(Equal("E"))
		})
		It("Clockwise east follows ", func() {
			Expect(DirectionClockwiseAfter(&Direction{Direction: "E"}).Direction).To(Equal("S"))
		})
		It("Clockwise south follows", func() {
			Expect(DirectionClockwiseAfter(&Direction{Direction: "S"}).Direction).To(Equal("W"))
		})
		It("Clockwise west follows", func() {
			Expect(DirectionClockwiseAfter(&Direction{Direction: "W"}).Direction).To(Equal("N"))
		})
	})
	Context("direction of travel on co-ordinates", func() {
		It("If the direction is north that's y increment", func() {
			direction, _ := NewDirection("N")
			coordinates := &Coordinate{1, 1}
			actual := direction.NextCoordinates(coordinates)
			Expect(actual.X).To(Equal(1), "X")
			Expect(actual.Y).To(Equal(2), "Y")
		})
		It("If the direction is south that's y decrement", func() {
			direction, _ := NewDirection("S")
			coordinates := &Coordinate{1, 1}
			actual := direction.NextCoordinates(coordinates)
			Expect(actual.X).To(Equal(1), "X")
			Expect(actual.Y).To(Equal(0), "Y")
		})
		It("If the direction is east that's x increment", func() {
			direction, _ := NewDirection("E")
			coordinates := &Coordinate{1, 1}
			actual := direction.NextCoordinates(coordinates)
			Expect(actual.X).To(Equal(2), "X")
			Expect(actual.Y).To(Equal(1), "Y")
		})
		It("If the direction is west that's x decrement", func() {
			direction, _ := NewDirection("W")
			coordinates := &Coordinate{1, 1}
			actual := direction.NextCoordinates(coordinates)
			Expect(actual.X).To(Equal(0), "X")
			Expect(actual.Y).To(Equal(1), "Y")
		})
	})
})
