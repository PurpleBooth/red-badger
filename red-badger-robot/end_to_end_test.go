package main_test

import (
	"bufio"
	"bytes"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/purplebooth/red-badger/red-badger-robot"
)

var _ = Describe("Red Badger App", func() {

	Describe("CLI Interface", func() {
		Context("Single robot", func() {
			It("should error if the upper bound co-ordinate is invalid", func() {
				inputLines := []string{
					"-1 -1",
				}

				errorOutput := []string{
					"Please enter valid co-ordinates for the upper bound of the map.",
					"",
				}

				runCommand(inputLines, []string{}, errorOutput, 1, nil)
			})

			It("should error if the first robot starting direction is invalid", func() {
				inputLines := []string{
					"50 50",
					"10 10 G",
				}

				errorOutput := []string{
					"Please enter valid robot starting direction",
					"",
				}

				runCommand(inputLines, []string{}, errorOutput, 2, nil)
			})
			It("should error if the first robot starting position is an invalid co-ordinate", func() {
				inputLines := []string{
					"50 50",
					"1 -1 N",
				}

				errorOutput := []string{
					"Please enter valid robot starting position",
					"",
				}

				runCommand(inputLines, []string{}, errorOutput, 3, nil)
			})
			It("should error if the first robot starting position is out of bounds on the map", func() {
				inputLines := []string{
					"40 40",
					"40 41 N",
				}

				errorOutput := []string{
					"Please enter valid robot starting position on the map",
					"",
				}

				runCommand(inputLines, []string{}, errorOutput, 4, nil)
			})
			It("should error if the robot's instruction set is invalid", func() {
				inputLines := []string{
					"50 50",
					"50 40 N",
					"RZZ",
				}

				errorOutput := []string{
					"The robots commands are invalid.",
					"",
				}

				runCommand(inputLines, []string{}, errorOutput, 5, nil)
			})
		})

		Context("Customer example", func() {

			It("should match the example given for a single robot", func() {
				inputLines := []string{
					"5 3",
					"1 1 E",
					"RFRFRFRF",
					"\n",
				}

				expectedOutput := []string{
					"1 1 E",
					"",
				}

				runCommand(inputLines, expectedOutput, []string{}, 0, nil)
			})

			It("should match the example given for multiple robots", func() {
				inputLines := []string{
					"5 3",
					"1 1 E",
					"RFRFRFRF",
					"3 2 N",
					"FRRFLLFFRRFLL",
					"0 3 W",
					"LLFFFLFLFL",
					"\n",
				}

				expectedOutput := []string{
					"1 1 E",
					"3 3 N LOST",
					"2 3 S",
					"",
				}

				runCommand(inputLines, expectedOutput, []string{}, 0, nil)
			})
		})

	})
})

func runCommand(inputLines []string, expectedOutput []string, expectedErr []string, expectedCode int, expectedError error) {
	stdin := bytes.NewBufferString(strings.Join(inputLines, "\n"))
	stdoutBuffer := new(bytes.Buffer)
	stderrBuffer := new(bytes.Buffer)

	code, err := main.ReadUserInput(
		bufio.NewScanner(stdin),
		bufio.NewWriter(stdoutBuffer),
		bufio.NewWriter(stderrBuffer),
	)
	Expect(stderrBuffer.String()).To(Equal(strings.Join(expectedErr, "\n")), "Standard Error")
	Expect(stdoutBuffer.String()).To(Equal(strings.Join(expectedOutput, "\n")), "Standard Out")
	Expect(code).To(Equal(expectedCode), "Return Code")

	if err != nil {

		Expect(err).To(Equal(expectedError), "Error")
	} else {
		Expect(err).To(BeNil(), "Error")
	}
}
