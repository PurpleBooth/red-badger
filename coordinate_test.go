package red_badger_test

import (
	. "github.com/purplebooth/red-badger"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("coordinate", func() {
	Context("Initialisation from string of co-ordinate alone", func() {

		It("should return co-ordinate", func() {
			actual, err := NewCoordinateFromString("1 2")

			Expect(err).To(BeNil())
			Expect(actual.X).To(Equal(1), "X axis")
			Expect(actual.Y).To(Equal(2), "Y axis")

		})

		It("should return an error on invalid string", func() {
			actual, err := NewCoordinateFromString("invalid")

			Expect(err).ToNot(BeNil())
			Expect(actual).To(BeNil())

		})

		It("X axis should be limited to 0 min", func() {
			actual, err := NewCoordinateFromString("-1 0")

			Expect(err).ToNot(BeNil())
			Expect(actual).To(BeNil())

		})
		It("Y axis should be limited to 0 min", func() {
			actual, err := NewCoordinateFromString("0 -1")

			Expect(err).ToNot(BeNil())
			Expect(actual).To(BeNil())

		})
	})
	Context("Initialisation from string of co-ordinate and direction", func() {

		It("should return co-ordinate", func() {
			actual, err := NewCoordinateFromCombinedString("1 2 N")

			Expect(err).To(BeNil())
			Expect(actual.X).To(Equal(1), "X axis")
			Expect(actual.Y).To(Equal(2), "Y axis")

		})

		It("should return an error on invalid string", func() {
			actual, err := NewCoordinateFromCombinedString("invalid")

			Expect(err).ToNot(BeNil())
			Expect(actual).To(BeNil())

		})

		It("X axis should be limited to 0 min", func() {
			actual, err := NewCoordinateFromCombinedString("-1 0 E")

			Expect(err).ToNot(BeNil())
			Expect(actual).To(BeNil())

		})
		It("Y axis should be limited to 0 min", func() {
			actual, err := NewCoordinateFromCombinedString("0 -1 W")

			Expect(err).ToNot(BeNil())
			Expect(actual).To(BeNil())

		})
	})

	Context("Comparison functions", func() {
		It("Can tell me if it's more northerly than a coordinate", func() {
			a, _ := NewCoordinate(10, 11)
			b, _ := NewCoordinate(50, 10)

			Expect(a.IsFurtherNorthThan(b)).To(BeTrue())
		})
		It("Can tell me if it's more southerly than a coordinate", func() {
			a, _ := NewCoordinate(10, 11)
			b, _ := NewCoordinate(50, 10)

			Expect(a.IsFurtherSouthThan(b)).To(BeFalse())
		})
		It("Can tell me if it's more westerly than a coordinate", func() {
			a, _ := NewCoordinate(10, 11)
			b, _ := NewCoordinate(50, 10)

			Expect(a.IsFurtherWestThan(b)).To(BeTrue())
		})
		It("Can tell me if it's more easterly than a coordinate", func() {
			a, _ := NewCoordinate(10, 11)
			b, _ := NewCoordinate(50, 10)

			Expect(a.IsFurtherEastThan(b)).To(BeFalse())
		})
	})
	Context("Is valid coordinate", func() {
		It("should tell me if the coordinate is not valid", func() {
			a, err := NewCoordinate(0, -1)

			Expect(err).ToNot(BeNil())
			Expect(a).To(BeNil())
		})
		It("should tell me if the coordinate is valid", func() {
			a, _ := NewCoordinate(0, 0)

			Expect(a.IsValid()).To(BeTrue())
		})
	})
})
