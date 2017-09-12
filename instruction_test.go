package red_badger_test

import (
	. "github.com/purplebooth/red-badger"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Instruction", func() {
	Context("New Instance", func() {
		It("returns an error for an unknown operation", func() {
			actual, err := NewInstructionFromString("Z")
			Expect(err).ToNot(BeNil(), "Error")
			Expect(actual).To(BeNil(), "Instruction")
		})
	})
	Context("Multiple Instances", func() {
		It("returns an error for an unknown operation", func() {
			actual, err := CreateInstructionsFromString("ZRSSE")
			Expect(err).ToNot(BeNil(), "Error")
			Expect(actual).To(BeNil(), "Instruction")
		})
		It("returns an error for an unknown operation", func() {
			actual, err := CreateInstructionsFromString("RFRFRFRF")
			Expect(err).To(BeNil(), "Error")
			Expect(actual).ToNot(BeNil(), "Instructions")
			Expect(actual).To(HaveLen(8), "Instruction Count")
		})
	})
	Context("Instructions", func() {
		It("Left", func() {
			actual, err := NewInstructionFromString("L")
			Expect(err).To(BeNil(), "error")
			Expect(actual).NotTo(BeNil(), "Instruction")
		})
		It("Right", func() {
			actual, err := NewInstructionFromString("R")
			Expect(err).To(BeNil(), "error")
			Expect(actual).NotTo(BeNil(), "Instruction")
		})
		It("Forwards", func() {
			actual, err := NewInstructionFromString("F")
			Expect(err).To(BeNil(), "error")
			Expect(actual).NotTo(BeNil(), "Instruction")
		})
	})

})
