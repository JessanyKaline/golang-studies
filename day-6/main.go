package main

import "fmt"

func main() {
	// msg, y := add(1, 2, 3, 4, 5, 6, 7, 8, 9)

	// fmt.Println(msg, y)

	// sliceTest := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// //Is necessary use ... to pass the slice to the function
	// z := learningSyntax(sliceTest...)

	// fmt.Println("The total2 is:", z)

	// aboutDefers()

	// jessany := person{"Jessany", 25}

	// jessany.sayHello()

	kaline := developer{person{"Kaline", 25}, 1000, true}

	kaline.hi()
}

// About function in golang with variadic parameter
// Is posible to pass a slice to a function using the ... syntax or zero or more arguments
// The variadic parameter must be the last parameter in the function
func add(x ...int) (string, int) {
	total := 0
	for _, v := range x {
		total += v
	}
	return "The total is:", total
}

func learningSyntax(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}

func aboutDefers() {
	// Defer is used to delay the execution of a statement until the end of the function
	// The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns
	// Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed

	defer fmt.Println("This is the first line")
	fmt.Println("This is the second line")
	fmt.Println("This is the third line")
}

// Methods are functions that are associated with a particular type
// A method is a function with a special receiver argument
// The receiver appears in its own argument list between the func keyword and the method name
// In this example, the Abs method has a receiver of type Vertex named v
// Methods with pointer receivers can modify the value to which the receiver points
// Since methods often need to modify their receiver, pointer receivers are more common than value
type person struct {
	name string
	age int
}

type developer struct {
	person
	wage float64
	isBackEnd bool
}

type architect struct { 
	person
	lastName string
}

func (x architect ) hi(){
	fmt.Println("Hi, my name is", x.name, "My last name is", x.lastName)
}

func (x developer) hi(){
	fmt.
	Println("Hi, my name is", x.name, "My wage is", x.wage)
}


func (p person) sayHello() {
	fmt.Println("Hello, my name is", p.name)
	
}

