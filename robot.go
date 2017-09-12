package red_badger

type Robot struct {
	CurrentCoordinate *Coordinate
	Instructions      []Instruction
	Direction         *Direction
	Lost              bool
}

func NewRobot(currentCoordinate *Coordinate, direction *Direction, instructions []Instruction) *Robot {
	return &Robot{
		CurrentCoordinate: currentCoordinate,
		Instructions:      instructions,
		Direction:         direction,
		Lost:              false,
	}
}

func (r *Robot) Run(world *worldMap) {
	for i := range r.Instructions {
		if keepGoing := r.Instructions[i].execute(world, r); !keepGoing {
			break
		}
	}
}
