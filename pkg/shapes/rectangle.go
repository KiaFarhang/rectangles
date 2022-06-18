/*
Package shapes implements code for working with 2D shapes
*/
package shapes

import "errors"

/*
Rectangle represents a rectangle in a 2D plane. Clients should not construct a Rectangle directly;
instead, use the NewRectangle function.
*/
type Rectangle struct {
	/*
		Note for reviewers: I didn't use pointers to these fields because they're small enough
		(two int values) that copying them shouldn't cost much memory. And by avoiding pointers,
		we get immutability because we're working with copies. To me, worth the tradeoff in this case.
	*/
	topLeft     Point
	bottomRight Point
}

var (
	errInvalidRectangleCoordinates = errors.New("a rectangle's bottom-right Point must be below (lesser Y) and to the right of (greater X) than its top-left Point")
)

/*
NewRectangle constructs a new Rectangle from the provided top-left and bottom-right points.

An error is returned if any of the following conditions are true:

- Bottom-right x <= top-left x
- Bottom-right y >= top-left y
*/
func NewRectangle(topLeft, bottomRight Point) (*Rectangle, error) {
	if bottomRight.X <= topLeft.X || bottomRight.Y >= topLeft.Y {
		return nil, errInvalidRectangleCoordinates
	}
	return &Rectangle{topLeft, bottomRight}, nil
}

/*
Contains returns whether the provided rectangle can be contained by this rectangle.
Rectangles with the exact same coordinates can contain one another.
*/
func (r *Rectangle) Contains(other *Rectangle) bool {
	thisTopLeft := r.topLeft
	otherTopLeft := other.topLeft

	if otherTopLeft.X < thisTopLeft.X || otherTopLeft.Y > thisTopLeft.Y {
		return false
	}

	thisBottomRight := r.bottomRight
	otherBottomRight := other.bottomRight

	if otherBottomRight.X > thisBottomRight.X || otherBottomRight.Y < thisBottomRight.Y {
		return false
	}

	return true
}
