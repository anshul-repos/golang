
Encapsulation: hiding the implementation details of obj from outside world

type Person struct{
 Name string
 Age int
}

func NewPerson(Name string,Age int)*Person{
	return &Person{name,age}
}

func (p *Person)getName()string{
	return p.name
}
func (p *Person)setName()string{
	
}
func (p *Person)getAge()int{
	return p.age
}
func (p *Person)setAge()int{}



Polymorphism : objects of different types to be treated as if they are of the same type


type Shape interface{
	Area() float
}

type Square struct{
	length float
}

type Circle struct{
	radius float
}


func (s Square)Area()float{
	return square formula //a2
}
func (s Circle)Area()float{
	return circle formula //3.14*r2
}


func main(){

	var s Shape
	s = Square(5.0)
	
	fmt.Print("Square",s.Area())
	
	s = Circle(3.0)
	fmt.Print("Circle",s.Area())
}


Inheritence : inheriting Parent struct properties and methods.

type Animal struct{}

type Dog struct{}


func (d Dog)Speak(){
"Dog barks"
}
func (a Animal)Speak(){
"Animal speaks"
}



func main(){

	a := Animal("Tom")
	a.Speak()
	
	
	d := Dog{
		Animal{"Charlie"},
		"Bulldog"
	}
	
	d.Speak()

}
