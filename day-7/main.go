package main

import "fmt"

type person struct {
	name string
	lastName string
	age int
}

func main() {
	aboutPointers()
	personOne := person{name: "Alice", lastName: "Ferreira", age: 20}
	fmt.Println("1", personOne)
	//Here we are passing the memory address of the personOne struct
	replacePerson(&personOne)
	//The modification is reflected in the original struct
	fmt.Println("2", personOne)

}

func aboutPointers() {
	// Pointers
	x := 5
	y := &x
	fmt.Println("init:", x, "in memory", y, "pointer", *y)
}

// The function replacePerson receives a pointer to a person struct
func replacePerson(p *person) {
	p.name = "John"
	p.age = 30
}
