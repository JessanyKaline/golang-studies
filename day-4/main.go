package main

import "fmt"

// Is declaration, after the value is assigned
var y [3]int

func main() {
	// y[0] =1
	// y[1]=2
	// fmt.Println(y[0], y[1])
	// // Index [3] in the case is 0 because the array is empty
	// fmt.Println(y)
	// aboutArray()
	// aboutSlice()
	// aboutSliceTwo()
	// aboutSliceMake()
	//
	// aboutMaps()
	aboutMapsRange()

}

func aboutArray() {
	//In go is necessary to define the length of the array
	x := [6]int{1, 2, 3, 4, 5, 6}
	// For range loop is used to iterate over the array (slices, string, etc)
	// i is the index and v is the value, and x is collection
	for i, v := range x {
		fmt.Println(i, v)
	}

}

func aboutSlice() {
	// Slice is more flexible than array in go and more used
	// Slice can have "n" quantity elements, capacity and length
	s := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(s)

	sTwo := append(s, 7, 8, 9)

	fmt.Println(sTwo)

	sliceString := []string{"a", "b", "c", "d", "e", "f"}

	for i, value := range sliceString {
		fmt.Println("The index is:", i, "value is:", value)
	}

	sliceInt := []int{1, 2, 3, 4, 5, 6}

	total := 0

	for _, v := range sliceInt {
		total += v
	}

	fmt.Println("The total is:", total)

}

func aboutSliceTwo() {

	pizzaFlavours := []string{"pepperoni", "margherita", "cheese", "hawaiian", "meatFeast"}
	// From index 1 to 3
	cutSlice := pizzaFlavours[1:3]
	// From index 2 to the end
	cutSliceTwo := pizzaFlavours[2:len(pizzaFlavours)]
	// From index 0 to 3
	cutSliceThree := pizzaFlavours[:3]
	// Acess all elements
	allSlices := pizzaFlavours[:]

	fmt.Println("Original", pizzaFlavours)
	fmt.Println("CutSlice", cutSlice)
	fmt.Println("CutSliceTwo", cutSliceTwo)
	fmt.Println("CutSliceThree", cutSliceThree)
	fmt.Println("AllSlices", allSlices)

	oneSlice := []int{1, 2, 3, 4, 5}
	otherslice := []int{9, 10}

	fmt.Println("Original oneSlice", oneSlice)

	oneSlice = append(oneSlice, 6, 7, 8)

	fmt.Println("oneSlice", oneSlice)

	oneSlice = append(oneSlice, otherslice...)

	fmt.Println("oneSlice + otherSlice", oneSlice)

	// With delete index in slice, here we are deleting the index 3, because between 2 and 4
	oneSlice = append(oneSlice[:2], oneSlice[4:]...)

	fmt.Println("oneSlice Delete", oneSlice)

}

func aboutSliceMake() {
	// Make is used to create a slice with a length and capacity
	slice := make([]int, 5, 8)

	slice[0], slice[1], slice[2], slice[3], slice[4] = 1, 2, 3, 4, 5

	fmt.Println(slice, "lenght:", len(slice), "cap:", cap(slice))

	// My lenght is 5, if I need more space, I can use append
	slice = append(slice, 6)
	slice = append(slice, 7)

	fmt.Println(slice, "lenght:", len(slice), "cap:", cap(slice))

	slice = append(slice, 8, 9, 10)

	fmt.Println(slice, "lenght:", len(slice), "cap:", cap(slice))

}

func aboutSliceMultiDimensional() {
	// Slice multi-dimensional

	ss := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9}}

	// Access the index 1 and 1
	fmt.Println(ss[1][1])

	slice := []string{"a", "b", "c"}
	sliceTwo := []string{"d", "e", "f"}

	sliceMulti := [][]string{slice, sliceTwo}

	fmt.Println(sliceMulti)
	fmt.Println(sliceMulti[1][2])

}

func aboutMaps() {
	// Maps is a collection of key value pair
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	fmt.Println(m)

	// Access the value of key "b"
	fmt.Println(m["b"])

	// Add a new key value pair
	m["d"] = 4

	fmt.Println(m)

	// Delete a key value pair
	delete(m, "d")

	fmt.Println(m)

	// Check if the key exists
	if v, ok := m["d"]; ok {
		fmt.Println("This is the value:", v)
	} else {
		fmt.Println("The key doesn't exist")
	}

	namesPhones := map[string]string{
		"John": "122345",
		"Jane": "123456",
		"Joe":  "123",
	}

	fmt.Println(namesPhones["Joe"])
	namesPhones["Jessany"] = "1234"
	// Is important the map is not ordered
	fmt.Println(namesPhones)
	delete(namesPhones, "Joe")

	if v, ok := namesPhones["Joe"]; ok {
		fmt.Println("The value is:", v)
	} else {
		fmt.Println("The key doesn't exist")
	}

	if v, ok := namesPhones["Joe"]; !ok {
		fmt.Println("The key doesn't exist")
	} else {
		fmt.Println("The key exists and the value is:", v)
	}

	if v, ok := namesPhones["Jessany"]; !ok {
		fmt.Println("The key doesn't exist")
	} else {
		fmt.Println("The key exists and the value is:", v)
	}
}

func aboutMapsRange() {
	// Maps is a collection of key value pair
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	// Range over the map
	for k, v := range m {
		fmt.Println(k, v)
	}

	// Range over the map and print only the key
	for k := range m {
		fmt.Println(k)
	}

	// Range over the map and print only the value
	for _, v := range m {
		fmt.Println(v)
	}

	numbers := map[string]int{
		"first":  3,
		"second": 2,
		"third":  1,
	}

	for k, v := range numbers {
		fmt.Println(k, v)
	}

	total := 0

	for _, v := range numbers {
		total += v
	}

	fmt.Println("The total is:", total)

}
