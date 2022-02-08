
package main

import "fmt"

//Shape interface implementing Area and Perimeter functions/methods
type Shape interface {
	Area() float64
	Perimeter() float64
}

//Rectangle struct
type Rectangle struct {
	width float64
	height float64
}

//St
func(r Rectangle) Area() float64 {
	return r.width * r.height
}
func(r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}


func main() {
	var s Shape
	s = Rectangle{5.0,6.0}
	r := Rectangle{5.0,6.0}
	fmt.Println(s)
	fmt.Println(r)
}



//In the above program, we’ve created the Shape interface and the struct type Rect.
//Then we defined methods like Area and Perimeter which belongs to Rect type, therefore Rect implemented those methods.
//Since these methods are defined by the Shape interface, the Rect struct type implements the Shape interface.
//Since we haven’t forced Rect to implement the Shape interface, it is all happening automatically.
//Hence, we can say that interfaces in Go are implicitly implemented.
//When a type implements an interface, a variable of that type can also be represented as the type of an interface.
//We can confirm that by creating a nil interface s of type Shape and assign a struct of type Rect.
//
//We have just achieved polymorphism.

