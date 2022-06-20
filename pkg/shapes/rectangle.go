/*
Package shapes implements code for working with 2D shapes
*/
package shapes

import (
	"errors"
)

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

func (r *Rectangle) getBottomLeft() Point {
	return Point{r.topLeft.X, r.bottomRight.Y}
}

func (r *Rectangle) getTopRight() Point {
	return Point{r.bottomRight.X, r.topLeft.Y}
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

/*
Adjacent returns whether this rectangle and the other rectangle provided
share a side. Rectangles that share a corner point are NOT considered adjacent.
(Problem statement didn't specify, so I chose to implement it this way)

This function is reversible; a.Adjacent(b) and b.Adjacent(a) will always return the
same value.
*/
func (r *Rectangle) Adjacent(other *Rectangle) bool {
	/*
		Find the wider of the two rectangles. If the narrower one has an x value that could allow
		for adjacency, check for adjacency. We need to know which rectangle is narrower to avoid
		situations where calling a.Adjacent(b) returns true but b.Adjacent(a) returns false.
	*/
	widerRectangle := r
	narrowerRectangle := other

	if (other.bottomRight.X - other.topLeft.X) > (r.bottomRight.X - r.topLeft.X) {
		widerRectangle = other
		narrowerRectangle = r
	}

	xIsInRange := (narrowerRectangle.topLeft.X >= widerRectangle.topLeft.X && narrowerRectangle.topLeft.X < widerRectangle.bottomRight.X) || (narrowerRectangle.bottomRight.X > widerRectangle.topLeft.X && narrowerRectangle.bottomRight.X <= widerRectangle.bottomRight.X)

	if xIsInRange {
		// Adjacent on top
		if other.bottomRight.Y == r.topLeft.Y {
			return true
		}

		// Adjacent on bottom
		if other.topLeft.Y == r.bottomRight.Y {
			return true
		}
	}

	tallerRectangle := r
	shorterRectangle := other

	if (other.topLeft.Y - other.bottomRight.Y) > (r.topLeft.Y - other.bottomRight.Y) {
		tallerRectangle = other
		shorterRectangle = r
	}

	yIsInRange := (shorterRectangle.topLeft.Y <= tallerRectangle.topLeft.Y && shorterRectangle.topLeft.Y > tallerRectangle.bottomRight.Y) || (shorterRectangle.bottomRight.Y < tallerRectangle.topLeft.Y && shorterRectangle.bottomRight.Y >= tallerRectangle.bottomRight.Y)

	if yIsInRange {
		// Adjacent on the right
		if other.topLeft.X == r.bottomRight.X {
			return true
		}

		// Adjacent on the left
		if other.bottomRight.X == r.topLeft.X {
			return true
		}
	}

	return false
}

/*
PointsOfIntersection returns a slice of all Points where this rectangle
intersects with the other rectangle provided, if any. The returned slice
will be empty if the rectangles do not intersect.

If two rectangles share a side (e.g. adjacency), any points on that side are NOT
considered points of intersection. (Problem statement didn't specify, so I chose to implement it this way)
*/
func (r *Rectangle) PointsOfIntersection(other *Rectangle) []Point {

	/**
	First, find the rectangle representing the area of intersection between the two rectangles.
	If this results in an error, there is no intersection, and we can just return an empty slice.
	*/

	topLeftX := max(r.topLeft.X, other.topLeft.X)
	topLeftY := min(r.topLeft.Y, other.topLeft.Y)
	bottomRightX := min(r.bottomRight.X, other.bottomRight.X)
	bottomRightY := max(r.bottomRight.Y, other.bottomRight.Y)

	rectangleOfIntersection, err := NewRectangle(Point{topLeftX, topLeftY}, Point{bottomRightX, bottomRightY})

	if err != nil {
		return []Point{}
	}

	/*
		Now, look at each corner of the rectangle of intersection. Determine if it intersects both original rectangles.
	*/

	allPointsOnRectangle := []Point{rectangleOfIntersection.topLeft, rectangleOfIntersection.bottomRight, rectangleOfIntersection.getBottomLeft(), rectangleOfIntersection.getTopRight()}

	pointsOfIntersection := []Point{}

	for _, point := range allPointsOnRectangle {
		if doesPointIntersect(point, r, other) {
			pointsOfIntersection = append(pointsOfIntersection, point)
		}
	}

	/*
		Finally, filter out any points of intersection that are also corners of the original rectangles - because
		these aren't considered real points of intersection. (Because they don't intersect THROUGH both rectangles)
	*/

	withoutCorners := []Point{}

	for _, point := range pointsOfIntersection {
		if point == r.topLeft || point == r.bottomRight || point == r.getBottomLeft() || point == r.getTopRight() {
			continue
		}
		if point == other.topLeft || point == other.bottomRight || point == other.getBottomLeft() || point == other.getTopRight() {
			continue
		}

		withoutCorners = append(withoutCorners, point)
	}

	return withoutCorners

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

/*
A point intersects with two rectangles if it's on the x-axis of one rectangle and the y-axis of the other, or vice versa
*/
func doesPointIntersect(point Point, rectangleA *Rectangle, rectangleB *Rectangle) bool {
	if rectangleA.topLeft.X == point.X || rectangleA.bottomRight.X == point.X {
		if rectangleB.topLeft.Y == point.Y || rectangleB.bottomRight.Y == point.Y {
			return true
		}
	}

	if rectangleA.topLeft.Y == point.Y || rectangleA.bottomRight.Y == point.Y {
		if rectangleB.topLeft.X == point.X || rectangleB.bottomRight.X == point.X {
			return true
		}
	}

	return false
}
