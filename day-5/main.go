package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
	isHappy   bool
}

type professional struct {
	person
	wage    float64
	company string
}

func main() {
	//aboutStructs()
	//exerciseStruct()
	exerciseStruct2()

}

func aboutStructs() {
	// Structs are a way to create complex data types
	// They are similar to classes in other languages
	// They are used to create custom data types
	// They are used to represent real-world entities
	// They are used to represent objects
	person1 := person{
		firstName: "John",
		lastName:  "Doe",
		age:       25,
		isHappy:   true,
	}

	fmt.Println(person1)
	// They are used to group together different types of data
	person2 := professional{
		person: person{
			firstName: "Jane",
			lastName:  "Doe",
			age:       30,
		},
		wage:    1000.00,
		company: "Google",
	}

	person3 := professional{person{"Jane", "Doe", 30, true}, 1000.00, "Google"}

	fmt.Println(person2)
	fmt.Println(person3)

	// Struct anonymous
	// Anonymous structs are structs without a name
	x := struct {
		age      int
		nickname string
	}{
		age:      25,
		nickname: "JK",
	}

	fmt.Println(x)

}

func exerciseStruct() {
	type young struct {
		name string
		age  int
		food []string
	}

	young1 := young{
		name: "John",
		age:  25,
		food: []string{"Pizza", "Burger", "Pasta"},
	}

	fmt.Println("Name:", young1.name, "Age:", young1.age)
	for _, v := range young1.food {
		fmt.Println("\t", v)
	}

	myMap := make(map[string]young)

	myMap["Test"] = young1

	for key, value := range myMap {
		fmt.Println("Key:", key)
		fmt.Println("\tName:", value.name, "Age:", value.age)
		for _, v := range value.food {
			fmt.Println("\t\t", v)
		}
	}
}

func exerciseStruct2() {
	type car struct {
		brand string
		color string
	}

	type sedan struct {
		car
		luxury bool
	}

	type bike struct {
		car
		sport bool
	}

	sedan1 := sedan{car{"BMW", "Black"}, true}
	bike1 := bike{car{"Yamaha", "Blue"}, true}

	fmt.Println(sedan1)
	fmt.Println(bike1)

	y := struct {
		name    string
		myMap   map[string]int
		mySlice []string
	}{
		name:    "John",
		myMap:   map[string]int{"test": 1, "test2": 2, "test3": 3, "test4": 4},
		mySlice: []string{"a", "b", "c"},
	}

	fmt.Println(y)

}
