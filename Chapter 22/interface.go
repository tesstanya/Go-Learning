// interface
package main

import (
	"fmt"
	"math"
)

// define an interface
type Shape interface {
	area() float64
}

// define a circle struct
type Circle struct {
	x, y, radius float64
}

// define a rectangle struct
type Rectangle struct {
	width, height float64
}

func (cir Circle) area() float64 {
	return math.Pi * cir.radius * cir.radius
}

func (rect Rectangle) area() float64 {
	return rect.width * rect.height
}

// define a method for shape
func getArea(shape Shape) float64 {
	return shape.area()
}

func main() {
	cir := Circle{x: 0, y: 0, radius: 5}
	rect := Rectangle{width: 10, height: 5}

	fmt.Printf("Circle area: %f\n", getArea(cir))
	fmt.Printf("Rectangle area: %f\n", getArea(rect))
}
