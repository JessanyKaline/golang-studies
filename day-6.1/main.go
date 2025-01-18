package main

import "fmt"

type Animal interface {
	speak() string
}

type dog struct{}

type cat struct{}

type bird struct{}

func (d dog) speak() string {
	return "Woof!"
}

func (c cat) speak() string {
	return "Meow!"
}

func (b bird) speak() string {
	return "Tweet!"
}

func makeSound(a Animal) {
	fmt.Println(a.speak())
}

func main() {
	var animal Animal

	animal = dog{}
	fmt.Println(animal.speak())

	animal = cat{}
	fmt.Println(animal.speak())

	animal = bird{}
	fmt.Println(animal.speak())

	dog := dog{}
	cat := cat{}
	bird := bird{}
	makeSound(dog)
	makeSound(cat)
	makeSound(bird)

	multiplesAnimals := []Animal{dog, cat, bird}

	for i, v := range multiplesAnimals {
		fmt.Println("Here", i, v.speak())
	}

}
