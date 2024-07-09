package main

import "fmt"

// Shape interface for Abstraction and Polymorphism
// Any type that implements the Area() method can be used as a Shape
type Shape interface {
	Area() float32
}

// Circle struct with encapsulated fields
type Circle struct {
	radius float32
}

// Rectangle struct with encapsulated fields
type Rectangle struct {
	length  float32
	breadth float32
}

// Square struct using composition to embed Rectangle
type Square struct {
	Rectangle
}

// Area method for Circle to implement Shape interface
func (c *Circle) Area() float32 {
	return 3.14 * c.radius * c.radius // Area of circle = Pi * r^2
}

// Area method for Rectangle to implement Shape interface
func (r *Rectangle) Area() float32 {
	return r.length * r.breadth // Area of rectangle = length * breadth
}

// Main function to create instances of shapes and calculate their areas
func main() {
	// Create instances of Circle, Rectangle, and Square
	cir := Circle{radius: 3}
	rect := Rectangle{length: 2, breadth: 4}
	sq := Square{Rectangle{length: 2, breadth: 2}}

	// Calculate and print the area of each shape directly in main
	fmt.Printf("Circle Area = %.2f \nRectangle Area = %.2f \nSquare Area = %.2f \n",
		cir.Area(), rect.Area(), sq.Area())
}
