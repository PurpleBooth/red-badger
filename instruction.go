package red_badger

import (
	"strings"
)

type Instruction interface {
	execute(worldMap *worldMap, robot *Robot) bool
}

type instructionLeft struct {
}

func (i instructionLeft) execute(worldMap *worldMap, robot *Robot) bool {
	robot.Direction = DirectionAntiClockwiseAfter(robot.Direction)

	return true
}

type instructionRight struct {
}

func (i instructionRight) execute(worldMap *worldMap, robot *Robot) bool {
	robot.Direction = DirectionClockwiseAfter(robot.Direction)

	return true
}

type UnknownInstructionError struct {
}

func (e UnknownInstructionError) Error() string {
	return "Unknown instruction"
}

func CreateInstructionsFromString(input string) ([]Instruction, error) {
	instructionTokens := strings.Split(input, "")
	instructions := []Instruction{}

	for i := range instructionTokens {
		instruction, err := NewInstructionFromString(instructionTokens[i])

		if err != nil {
			return nil, err
		}

		instructions = append(instructions, instruction)
	}

	return instructions, nil
}

func NewInstructionFromString(input string) (Instruction, error) {
	instruction, exists := availableInstructions()[input]

	if !exists {
		return nil, UnknownInstructionError{}
	}

	return instruction, nil
}

type instructionForward struct{}

func (i instructionForward) execute(worldMap *worldMap, robot *Robot) bool {
	proposedCoordinate := robot.Direction.NextCoordinates(robot.CurrentCoordinate)

	if !worldMap.HasCoordinate(proposedCoordinate) {
		if worldMap.HasSmell(proposedCoordinate) {

			return true
		}

		worldMap.AddSmell(proposedCoordinate)
		robot.Lost = true

		return false
	}

	robot.CurrentCoordinate = proposedCoordinate

	return true
}

func availableInstructions() map[string]Instruction {
	instructions := make(map[string]Instruction)

	instructions["L"] = &instructionLeft{}
	instructions["R"] = &instructionRight{}
	instructions["F"] = &instructionForward{}

	return instructions
}
