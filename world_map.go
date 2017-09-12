package red_badger

type worldMap struct {
	topMostCoordinate *Coordinate
	robots            []*Robot
	smells            map[int]bool
}

type OutOfBoundsOnWorldMapError struct {
}

func (e OutOfBoundsOnWorldMapError) Error() string {
	return "Invalid world map location, out of bounds"
}

func NewWorldMap(topMostCoordinate *Coordinate) *worldMap {
	return &worldMap{topMostCoordinate: topMostCoordinate, smells: map[int]bool{}}
}

func (w worldMap) HasCoordinate(coordinate *Coordinate) bool {
	if coordinate.IsFurtherNorthThan(w.topMostCoordinate) || coordinate.IsFurtherEastThan(w.topMostCoordinate) {
		return false
	}

	if !coordinate.IsValid() {
		return false
	}

	return true
}

func (i worldMap) HasSmell(coordinate *Coordinate) bool {

	if _, exists := i.smells[(2*coordinate.X)+coordinate.Y]; exists {
		return true
	}
	return false

}
func (i worldMap) AddSmell(coordinate *Coordinate) {
	i.smells[(2*coordinate.X)+coordinate.Y] = true
}
