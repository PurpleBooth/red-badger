package main

import (
	"bufio"
	"log"
	"os"

	"fmt"

	"strings"

	"github.com/purplebooth/red-badger"
)

func main() {
	stdin := bufio.NewScanner(os.Stdin)
	stdout := bufio.NewWriter(os.Stdout)
	stderr := bufio.NewWriter(os.Stderr)

	code, err := ReadUserInput(stdin, stdout, stderr)

	if err != nil {
		log.Fatal(err)
	}

	os.Exit(code)
}

func ReadUserInput(stdin *bufio.Scanner, stdout *bufio.Writer, stderr *bufio.Writer) (int, error) {

	stdin.Scan()
	upperBoundLine := stdin.Text()
	upperCoordinate, err := red_badger.NewCoordinateFromString(strings.Trim(upperBoundLine, " \n"))

	if err != nil {
		stderr.WriteString("Please enter valid co-ordinates for the upper bound of the map.\n")
		stderr.Flush()

		return 1, nil
	}

	worldMap := red_badger.NewWorldMap(upperCoordinate)

	var currentDirection *red_badger.Direction
	var currentCoordinate *red_badger.Coordinate
	lineCount := 0
	robots := []*red_badger.Robot{}

	for stdin.Scan() {
		line := stdin.Text()
		if line == "" {
			break
		}

		switch lineCount {
		case 0:
			currentDirection, err = red_badger.NewDirectionFromCombinedString(line)

			if err != nil {
				stderr.WriteString("Please enter valid robot starting direction\n")
				stderr.Flush()

				return 2, nil
			}

			currentCoordinate, err = red_badger.NewCoordinateFromCombinedString(line)

			if err != nil {
				stderr.WriteString("Please enter valid robot starting position\n")
				stderr.Flush()

				return 3, nil
			}

			if !worldMap.HasCoordinate(currentCoordinate) {
				stderr.WriteString("Please enter valid robot starting position on the map\n")
				stderr.Flush()

				return 4, nil
			}
		case 1:
			currentInstructions, err := red_badger.CreateInstructionsFromString(line)

			if err != nil {
				stderr.WriteString("The robots commands are invalid.\n")
				stderr.Flush()

				return 5, nil
			}

			robots = append(robots, red_badger.NewRobot(currentCoordinate, currentDirection, currentInstructions))
		}

		lineCount = (lineCount + 1) % 2
	}

	for i := range robots {
		robots[i].Run(worldMap)
		stdout.WriteString(
			fmt.Sprintf(
				"%d %d %s",
				robots[i].CurrentCoordinate.X,
				robots[i].CurrentCoordinate.Y,
				robots[i].Direction.Direction,
			),
		)

		if robots[i].Lost {
			stdout.WriteString(" LOST")
		}

		stdout.WriteString("\n")

		stdout.Flush()
	}

	return 0, nil
}
