package shapes

import "errors"

type Rectangle struct {
	topLeft     *Point
	bottomRight *Point
}

var (
	errInvalidRectangleCoordinates = errors.New("a rectangle's bottom-right Point must be below (lesser Y) and to the right of (greater X) than its top-left Point")
)

func NewRectangle(topLeft, bottomRight *Point) (*Rectangle, error) {
	if topLeft.X >= bottomRight.X || topLeft.Y <= bottomRight.Y {
		return nil, errInvalidRectangleCoordinates
	}
	return &Rectangle{}, nil
}

func (r *Rectangle) contains(other *Rectangle) bool {
	return true
}
