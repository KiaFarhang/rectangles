package main

import (
	"fmt"
	"log"

	"github.com/KiaFarhang/rectangles/pkg/shapes"
)

func main() {
	demonstrateContainment()
	demonstrateAdjacency()
}

func demonstrateContainment() {
	rectangleA := buildRectangle(shapes.Point{X: 2, Y: 4}, shapes.Point{X: 5, Y: 2})
	rectangleB := buildRectangle(shapes.Point{X: 3, Y: 3}, shapes.Point{X: 4, Y: 2})

	fmt.Println("Containment demonstration:")
	fmt.Printf("Rectangle A: %v\nRectangle B: %v\n", rectangleA, rectangleB)
	fmt.Printf("A contains B: %t\n", rectangleA.Contains(rectangleB))
	fmt.Printf("B contains A: %t\n", rectangleB.Contains(rectangleA))
}

func demonstrateAdjacency() {
	rectangleA := buildRectangle(shapes.Point{X: 2, Y: 4}, shapes.Point{X: 5, Y: 2})
	adjacentRectangle := buildRectangle(shapes.Point{X: 2, Y: 2}, shapes.Point{X: 3, Y: 0})

	fmt.Println("Adjacency demonstration:")
	fmt.Printf("Rectangle A: %v\n Rectangle B: %v\n", rectangleA, adjacentRectangle)
	fmt.Printf("Are rectangles adjacent: %t\n", rectangleA.Adjacent(adjacentRectangle))

	nonAdjacentRectangle := buildRectangle(shapes.Point{X: 10, Y: 10}, shapes.Point{X: 14, Y: 7})
	fmt.Printf("Rectangle A: %v\n Rectangle C: %v\n", rectangleA, nonAdjacentRectangle)
	fmt.Printf("Are rectangles adjacent: %t\n", rectangleA.Adjacent(nonAdjacentRectangle))
}

func buildRectangle(topLeft, bottomRight shapes.Point) *shapes.Rectangle {
	rectangle, err := shapes.NewRectangle(topLeft, bottomRight)
	if err != nil {
		log.Fatalf("error constructing rectangle: %s", err)
	}

	return rectangle
}
