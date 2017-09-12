package red_badger_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/purplebooth/red-badger"
)

var _ = Describe("Map", func() {
	Context("Initialisation", func() {

		It("should be creatable from co-ordinate", func() {
			actual := NewWorldMap(&Coordinate{50, 40})
			Expect(actual).ToNot(BeNil())
		})
	})
	Context("Has Coordinates", func() {
		It("should tell me if a coordinate is on the map", func() {
			worldMap := NewWorldMap(&Coordinate{50, 40})
			coordinate := &Coordinate{X: 40, Y: 40}

			Expect(worldMap.HasCoordinate(coordinate)).To(BeTrue())
		})
		It("should tell me if a coordinate is not on the map", func() {
			worldMap := NewWorldMap(&Coordinate{50, 40})
			coordinate := &Coordinate{X: 51, Y: 40}

			Expect(worldMap.HasCoordinate(coordinate)).To(BeFalse())
		})
		It("should tell me if a coordinate is totally invalid", func() {
			worldMap := NewWorldMap(&Coordinate{-1, 0})
			coordinate := &Coordinate{X: 51, Y: 40}

			Expect(worldMap.HasCoordinate(coordinate)).To(BeFalse())
		})
	})
	Context("Smells", func() {
		It("it can tell me if a coordinate has a smell", func() {
			worldMap := NewWorldMap(&Coordinate{50, 40})
			worldMap.AddSmell(&Coordinate{X: 51, Y: 30})
			Expect(worldMap.HasSmell(&Coordinate{X: 51, Y: 30})).To(BeTrue())
			Expect(worldMap.HasSmell(&Coordinate{X: 30, Y: 51})).To(BeFalse())
		})
	})
})
