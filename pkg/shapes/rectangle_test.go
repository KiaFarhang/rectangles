package shapes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRectangle_Constructor(t *testing.T) {
	t.Run("returns an error if the top left and bottom right parameters are improperly positioned", func(t *testing.T) {
		type test struct {
			topLeft     Point
			bottomRight Point
			description string
		}

		tests := []test{
			{topLeft: Point{0, 0}, bottomRight: Point{-1, -3}, description: "Bottom right behind top left"},
			{topLeft: Point{0, 0}, bottomRight: Point{2, 2}, description: "Bottom right above top left"},
		}

		for _, testCase := range tests {

			_, err := NewRectangle(testCase.topLeft, testCase.bottomRight)
			assert.ErrorIsf(t, err, errInvalidRectangleCoordinates, "test case for invalid rectangle construction failed: %s", testCase.description)

		}
	})
	t.Run("Returns a Rectangle with the given points if their positioning is valid", func(t *testing.T) {
		topLeft := Point{2, 5}
		bottomRight := Point{4, 2}

		rectangle, err := NewRectangle(topLeft, bottomRight)

		assert.Equal(t, rectangle.topLeft.X, topLeft.X)
		assert.Equal(t, rectangle.topLeft.Y, topLeft.Y)
		assert.Equal(t, rectangle.bottomRight.X, bottomRight.X)
		assert.Equal(t, rectangle.bottomRight.Y, bottomRight.Y)
		assert.NoError(t, err)

	})
}

func TestContainment(t *testing.T) {
	t.Run("returns true if the other rectangle is the same size and coordinates as the first", func(t *testing.T) {
		topLeft := Point{4, 4}
		bottomRight := Point{6, 2}

		rectangleA, err := NewRectangle(topLeft, bottomRight)
		assert.NoError(t, err)

		rectangleB, err := NewRectangle(topLeft, bottomRight)
		assert.NoError(t, err)

		assert.True(t, rectangleA.Contains(rectangleB))
	})
	t.Run("returns true if the other rectangle is fully contained within the first", func(t *testing.T) {
		topLeftA := Point{2, 4}
		bottomRightA := Point{5, 2}

		topLeftB := Point{3, 3}
		bottomRightB := Point{4, 2}

		rectangleA, err := NewRectangle(topLeftA, bottomRightA)
		assert.NoError(t, err)

		rectangleB, err := NewRectangle(topLeftB, bottomRightB)
		assert.NoError(t, err)

		assert.True(t, rectangleA.Contains(rectangleB))
	})
	t.Run("returns false if the other rectangle is completely outside the first", func(t *testing.T) {
		topLeftA := Point{2, 4}
		bottomRightA := Point{5, 2}

		topLeftB := Point{2, 7}
		bottomRightB := Point{3, 6}

		rectangleA, err := NewRectangle(topLeftA, bottomRightA)
		assert.NoError(t, err)

		rectangleB, err := NewRectangle(topLeftB, bottomRightB)
		assert.NoError(t, err)

		assert.False(t, rectangleA.Contains(rectangleB))
	})
	t.Run("returns false if the other rectangle is partially contained within the first", func(t *testing.T) {
		topLeftA := Point{2, 4}
		bottomRightA := Point{5, 2}

		topLeftB := Point{3, 3}
		bottomRightB := Point{6, 2}

		rectangleA, err := NewRectangle(topLeftA, bottomRightA)
		assert.NoError(t, err)

		rectangleB, err := NewRectangle(topLeftB, bottomRightB)
		assert.NoError(t, err)

		assert.False(t, rectangleA.Contains(rectangleB))
	})
}

func TestAdjacency(t *testing.T) {
	rectangle, _ := NewRectangle(Point{2, 4}, Point{5, 2})
	type test struct {
		topLeft     Point
		bottomRight Point
		expected    bool
		description string
	}
	tests := []test{
		{topLeft: Point{6, 4}, bottomRight: Point{7, 3}, expected: false, description: "no adjacency on right side"},
		{topLeft: Point{5, 2}, bottomRight: Point{7, 0}, expected: false, description: "no adjacency on right side - corner case"},
		{topLeft: Point{5, 4}, bottomRight: Point{7, 3}, expected: true, description: "adjacency on right side"},
		{topLeft: Point{3, 6}, bottomRight: Point{4, 5}, expected: false, description: "no adjacency on top side"},
		{topLeft: Point{3, 6}, bottomRight: Point{4, 4}, expected: true, description: "adjacency on top side"},
		{topLeft: Point{-1, 6}, bottomRight: Point{1, 5}, expected: false, description: "no adjacency on left side"},
		{topLeft: Point{-1, 6}, bottomRight: Point{2, 4}, expected: false, description: "no adjacency on left side - corner case"},
		{topLeft: Point{-1, 6}, bottomRight: Point{2, 3}, expected: true, description: "adjacency on left side"},
		{topLeft: Point{2, 1}, bottomRight: Point{3, 0}, expected: false, description: "no adjacency on bottom side"},
		{topLeft: Point{2, 2}, bottomRight: Point{3, 0}, expected: true, description: "adjacency on bottom side"},
		{topLeft: Point{10, 10}, bottomRight: Point{14, 7}, expected: false, description: "no adjacency"},
	}

	for _, testCase := range tests {
		t.Run(testCase.description, func(t *testing.T) {
			otherRectangle, err := NewRectangle(testCase.topLeft, testCase.bottomRight)
			assert.NoError(t, err)
			assert.Equal(t, testCase.expected, rectangle.Adjacent(otherRectangle))
			// Adjacency should work the same regardless of which rectangle you call it on
			assert.Equal(t, testCase.expected, otherRectangle.Adjacent(rectangle))
		})
	}
}

func TestPointsOfIntersection(t *testing.T) {
	type test struct {
		topLeftA             Point
		bottomRightA         Point
		topLeftB             Point
		bottomRightB         Point
		pointsOfIntersection []Point
	}

	tests := []test{
		{topLeftA: Point{4, 7}, bottomRightA: Point{6, 4}, topLeftB: Point{2, 5}, bottomRightB: Point{5, 3}, pointsOfIntersection: []Point{{4, 5}, {5, 4}}},
		{topLeftA: Point{3, 6}, bottomRightA: Point{6, 3}, topLeftB: Point{4, 7}, bottomRightB: Point{7, 4}, pointsOfIntersection: []Point{{4, 6}, {6, 4}}},
		{topLeftA: Point{-4, 4}, bottomRightA: Point{2, 2}, topLeftB: Point{-2, 3}, bottomRightB: Point{2, 0}, pointsOfIntersection: []Point{{-2, 2}}},
		{topLeftA: Point{1, 3}, bottomRightA: Point{2, 2}, topLeftB: Point{1, 1}, bottomRightB: Point{4, 0}, pointsOfIntersection: []Point{}},
		{topLeftA: Point{-2, 10}, bottomRightA: Point{2, 5}, topLeftB: Point{-3, 8}, bottomRightB: Point{3, 6}, pointsOfIntersection: []Point{{-2, 8}, {-2, 6}, Point{2, 8}, Point{2, 6}}},
	}

	for _, testCase := range tests {
		rectangleA, err := NewRectangle(testCase.topLeftA, testCase.bottomRightA)
		assert.NoError(t, err)

		rectangleB, err := NewRectangle(testCase.topLeftB, testCase.bottomRightB)
		assert.NoError(t, err)

		pointsOfIntersection := rectangleA.PointsOfIntersection(rectangleB)
		assert.Len(t, pointsOfIntersection, len(testCase.pointsOfIntersection))

		// Like adjacency, intersection should give the same results both ways
		reversedPointsOfIntersection := rectangleB.PointsOfIntersection(rectangleA)
		assert.Len(t, reversedPointsOfIntersection, len(testCase.pointsOfIntersection))
		for _, point := range testCase.pointsOfIntersection {
			assert.Contains(t, pointsOfIntersection, point)
			assert.Contains(t, reversedPointsOfIntersection, point)
		}
	}
}
