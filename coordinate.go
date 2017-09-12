package red_badger

import (
	"regexp"
	"strconv"
	"strings"
)

const COORDINATE_MIN_X_VALUE = 0
const COORDINATE_MIN_Y_VALUE = 0

const COORDINATE_MAX_X_VALUE = 50
const COORDINATE_MAX_Y_VALUE = 50

type BadPatternCoordinatesError struct {
}

func (e BadPatternCoordinatesError) Error() string {
	return "Invalid co-ordinates, incorrect pattern"
}

type OutOfBoundsCoordinatesError struct {
}

func (e OutOfBoundsCoordinatesError) Error() string {
	return "Invalid co-ordinates, out of bounds"
}

type Coordinate struct {
	X int
	Y int
}

func NewCoordinateFromCombinedString(input string) (*Coordinate, error) {
	matches, err := regexp.MatchString("^\\d\\d? \\d\\d? ", input)

	if err != nil {
		return nil, err
	}

	if !matches {
		return nil, BadPatternCoordinatesError{}
	}

	return newCoordinateFromSafeString(input)
}

func NewCoordinateFromString(input string) (*Coordinate, error) {
	matches, err := regexp.MatchString("^\\d\\d? \\d\\d?$", input)

	if err != nil {
		return nil, err
	}

	if !matches {
		return nil, BadPatternCoordinatesError{}
	}

	return newCoordinateFromSafeString(input)
}

func NewCoordinate(x int, y int) (*Coordinate, error) {
	coordinate := &Coordinate{X: x, Y: y}
	if !coordinate.IsValid() {
		return nil, OutOfBoundsCoordinatesError{}
	}

	return coordinate, nil
}

func (c Coordinate) IsValid() bool {
	lowestPossibleCoordinate := &Coordinate{X: COORDINATE_MIN_X_VALUE, Y: COORDINATE_MIN_Y_VALUE}
	highestPossibleCoordinate := &Coordinate{X: COORDINATE_MAX_X_VALUE, Y: COORDINATE_MAX_Y_VALUE}

	if c.IsFurtherWestThan(lowestPossibleCoordinate) || c.IsFurtherSouthThan(lowestPossibleCoordinate) {
		return false
	}
	if c.IsFurtherEastThan(highestPossibleCoordinate) || c.IsFurtherNorthThan(highestPossibleCoordinate) {
		return false
	}

	return true
}

func (c Coordinate) IsFurtherNorthThan(b *Coordinate) bool {
	return c.Y > b.Y
}

func (c Coordinate) IsFurtherSouthThan(b *Coordinate) bool {
	return c.Y < b.Y
}

func (c Coordinate) IsFurtherWestThan(b *Coordinate) bool {
	return c.X < b.X
}
func (c Coordinate) IsFurtherEastThan(b *Coordinate) bool {
	return c.X > b.X
}

func newCoordinateFromSafeString(input string) (*Coordinate, error) {
	rawCoordinates := strings.Split(input, " ")
	parsedCoordinates := []int{}

	for i := range rawCoordinates {
		if i > 1 {
			break
		}

		parsedCoordinate, err := strconv.Atoi(rawCoordinates[i])

		if err != nil {
			return nil, err
		}

		parsedCoordinates = append(parsedCoordinates, parsedCoordinate)
	}

	return NewCoordinate(parsedCoordinates[0], parsedCoordinates[1])
}
