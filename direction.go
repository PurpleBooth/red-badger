package red_badger

import (
	"fmt"
	"regexp"
	"strings"
)

type BadPatternDirectionError struct {
}

func (e BadPatternDirectionError) Error() string {
	return "Invalid Direction, incorrect pattern"
}

type Direction struct {
	Direction string
}

func NewDirection(token string) (*Direction, error) {
	if _, exists := directions()[token]; !exists {
		return nil, BadPatternDirectionError{}
	}

	return &Direction{Direction: token}, nil
}

func NewDirectionFromCombinedString(input string) (*Direction, error) {
	directions := strings.Join(directionsOrderedClockwise(), "")

	matches, err := regexp.Match(fmt.Sprintf(" [%s]$", directions), []byte(input))

	if err != nil {
		return nil, err
	}

	if !matches {
		return nil, BadPatternDirectionError{}
	}

	return &Direction{Direction: input[len(input)-1:]}, nil
}

func DirectionAntiClockwiseAfter(start *Direction) *Direction {
	directionsOrderedClockwise := directionsOrderedClockwise()

	for i := len(directionsOrderedClockwise) - 1; i >= 0; i-- {
		if directionsOrderedClockwise[i] == start.Direction {
			if i == 0 {
				return &Direction{Direction: directionsOrderedClockwise[len(directionsOrderedClockwise)-1]}
			}

			return &Direction{Direction: directionsOrderedClockwise[(i - 1)]}
		}
	}

	return nil // Unreachable statement
}

func (d *Direction) NextCoordinates(start *Coordinate) *Coordinate {
	switch d.Direction {
	case "N":
		return &Coordinate{X: start.X, Y: start.Y + 1}
	case "E":
		return &Coordinate{X: start.X + 1, Y: start.Y}
	case "W":
		return &Coordinate{X: start.X - 1, Y: start.Y}
	default:
		return &Coordinate{X: start.X, Y: start.Y - 1}
	}
}

func DirectionClockwiseAfter(start *Direction) *Direction {
	directionsOrderedClockwise := directionsOrderedClockwise()

	for i := range directionsOrderedClockwise {
		if directionsOrderedClockwise[i] == start.Direction {
			if i == (len(directionsOrderedClockwise) - 1) {
				return &Direction{Direction: directionsOrderedClockwise[0]}
			}

			return &Direction{Direction: directionsOrderedClockwise[(i + 1)]}
		}
	}

	return nil // Unreachable statement
}

func directionsOrderedClockwise() []string {
	return []string{"N", "E", "S", "W"}
}

func directions() map[string]bool {
	set := make(map[string]bool)
	set["N"] = true
	set["E"] = true
	set["S"] = true
	set["W"] = true

	return set
}
