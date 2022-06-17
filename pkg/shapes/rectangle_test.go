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
	t.Run("returns true if the other rectangle is the topLeft.Y <= bottomRight.Ysame size + location as the first", func(t *testing.T) {
		topLeft := Point{X: 4, Y: 4}
		bottomRight := Point{X: 6, Y: 2}

		rectangleA, err := NewRectangle(topLeft, bottomRight)
		assert.NoError(t, err)

		rectangleB, err := NewRectangle(topLeft, bottomRight)
		assert.NoError(t, err)

		assert.Equal(t, true, rectangleA.Contains(rectangleB))

	})
}
