package main

import "fmt"

// Go में अगर कोई struct उन सभी methods को implement करता है जो किसी interface में declared हैं, तो वह अपने आप उस interface को implement कर देता है।
// Rectangle aur Circle ने Area() और Perimeter() method implement कर दिए हैं, इसलिए यह अपने आप Shape interface को implement कर रहा है — हमें कुछ declare करने की ज़रूरत नहीं।

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Length  float64
	Breadth float64
}

func (c Circle) Area() float64 {
	return 2 * 3.24 * c.Radius * c.Radius
}

func (r *Rectangle) Area() float64 {
	return r.Length * r.Breadth
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.24 * c.Radius
}

func (r *Rectangle) Perimeter() float64 {
	return 2 * (r.Length + r.Breadth)
}

func main() {
	var s Shape
	s = Circle{Radius: 2}
	fmt.Println("\n Area of Circle: ", s.Area())

	s = &Rectangle{Length: 2, Breadth: 2}
	fmt.Println("\n Area of Rectangle: ", s.Area())
}

// एक ही interface variable s के ज़रिए, आप अलग-अलग type के objects (Circle और Rectangle) को handle कर रहे हैं, और उनके methods को call कर रहे हैं — बिना ये जाने कि s के अंदर कौन सा exact type रखा गया है।
// 🎯 यही होता है Polymorphism:
// "एक interface, कई रूप (Many Forms)"
