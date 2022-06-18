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
		topLeft := Point{X: 4, Y: 4}
		bottomRight := Point{X: 6, Y: 2}

		rectangleA, err := NewRectangle(topLeft, bottomRight)
		assert.NoError(t, err)

		rectangleB, err := NewRectangle(topLeft, bottomRight)
		assert.NoError(t, err)

		assert.Equal(t, true, rectangleA.Contains(rectangleB))
	})
	t.Run("returns true if the other rectangle is fully contained within the first", func(t *testing.T) {
		topLeftA := Point{X: 2, Y: 4}
		bottomRightA := Point{X: 5, Y: 2}

		topLeftB := Point{X: 3, Y: 3}
		bottomRightB := Point{X: 4, Y: 2}

		rectangleA, err := NewRectangle(topLeftA, bottomRightA)
		assert.NoError(t, err)

		rectangleB, err := NewRectangle(topLeftB, bottomRightB)
		assert.NoError(t, err)

		assert.Equal(t, true, rectangleA.Contains(rectangleB))
	})
	t.Run("Returns false if the other rectangle is completely outside the first", func(t *testing.T) {
		topLeftA := Point{X: 2, Y: 4}
		bottomRightA := Point{X: 5, Y: 2}

		topLeftB := Point{X: 2, Y: 7}
		bottomRightB := Point{X: 3, Y: 6}

		rectangleA, err := NewRectangle(topLeftA, bottomRightA)
		assert.NoError(t, err)

		rectangleB, err := NewRectangle(topLeftB, bottomRightB)
		assert.NoError(t, err)

		assert.Equal(t, false, rectangleA.Contains(rectangleB))
	})
	t.Run("Returns false if the other rectangle is partially contained within the first", func(t *testing.T) {
		topLeftA := Point{X: 2, Y: 4}
		bottomRightA := Point{X: 5, Y: 2}

		topLeftB := Point{X: 3, Y: 3}
		bottomRightB := Point{X: 6, Y: 2}

		rectangleA, err := NewRectangle(topLeftA, bottomRightA)
		assert.NoError(t, err)

		rectangleB, err := NewRectangle(topLeftB, bottomRightB)
		assert.NoError(t, err)

		assert.Equal(t, false, rectangleA.Contains(rectangleB))
	})
}
