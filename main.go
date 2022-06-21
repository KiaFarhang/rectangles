package main

import (
	"fmt"
	"log"

	"github.com/KiaFarhang/rectangles/pkg/shapes"
)

func main() {
	demonstrateContainment()
}

func demonstrateContainment() {
	rectangleA := buildRectangle(shapes.Point{X: 2, Y: 4}, shapes.Point{X: 5, Y: 2})
	rectangleB := buildRectangle(shapes.Point{X: 3, Y: 3}, shapes.Point{X: 4, Y: 2})

	fmt.Println("Containment demonstration:")
	fmt.Printf("Rectangle A: %v\nRectangle B: %v\n", rectangleA, rectangleB)
	fmt.Printf("A contains B: %t\n", rectangleA.Contains(rectangleB))
	fmt.Printf("B contains A: %t\n", rectangleB.Contains(rectangleA))
}

func buildRectangle(topLeft, bottomRight shapes.Point) *shapes.Rectangle {
	rectangle, err := shapes.NewRectangle(topLeft, bottomRight)
	if err != nil {
		log.Fatalf("error constructing rectangle: %s", err)
	}

	return rectangle
}
