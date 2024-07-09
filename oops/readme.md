# OOPS in Go

OOPS is based on 4 main principles: 
1. Encapsulation 
2. Abstraction
3. Polymorphism
4. Inheritance

## Encapsulation
- Hiding implementation details of one object from another object. 
- It is a process which encourages decoupling.

## Abstraction
- hiding the complex implementation details of an object and exposing only the relevant information to the user.
- data are visible only to semantically related functions, so as to prevent misuse
- it can be achieved using interfaces

## Polymorphism
- perform a single action in different ways
- objects of different types to be treated as if they were the same type
- Polymorphism makes it easier to write reusable code that can work with different types of objects, without needing to know the specific details of each object.
- in other words polymorphism allows you to define one interface and have multiple implementations

## Inheritance
- creating a new class that is a modified version of an existing class. 
- The new class, called the child class or subclass, inherits the properties and methods of the parent class or superclass.
- Can be implemented using embeded struct

## Code Example Explanation
```
Abstraction:
The Shape interface abstracts the concept of a shape that can calculate its area through the Area method.

Encapsulation:
The Circle and Rectangle structs encapsulate their fields (radius, length, breadth) and provide methods to operate on them.

Polymorphism:
The Area method is implemented by different structs (Circle, Rectangle), and they can all be treated as Shape types due to the interface.

Inheritance (Composition):
The Square struct embeds the Rectangle struct, allowing it to inherit its fields and methods. This demonstrates composition in Go, which is used in place of traditional inheritance.

Main Function:
Instances of Circle, Rectangle, and Square are created.
The areas of these shapes are calculated and printed directly within the main function, demonstrating how each shape correctly implements the Area method.
```