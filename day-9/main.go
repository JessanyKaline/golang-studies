package main

import (
	"fmt"
	"sort"
)

type Car struct {
	name  string
	price int
	year  int
}


// Implementing the sort interface
type OrderByPrice []Car

func (a OrderByPrice) Len() int           { return len(a) }
func (a OrderByPrice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a OrderByPrice) Less(i, j int) bool { return a[i].price < a[j].price }

type OrderByYear []Car

func (a OrderByYear) Len() int           { return len(a) }
func (a OrderByYear) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a OrderByYear) Less(i, j int) bool { return a[i].year < a[j].year }

type OrderByName []Car

func (a OrderByName) Len() int           { return len(a) }
func (a OrderByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a OrderByName) Less(i, j int) bool { return a[i].name < a[j].name }

func main() {

	car := []Car{{"Corsa", 15000, 2012}, {"Gol", 20000, 2000}, {"Uno", 10000, 1999}}

	fmt.Println("Original", car)

	sort.Sort(OrderByPrice(car))

	fmt.Println("Sort by price", car)

	sort.Sort(OrderByYear(car))

	fmt.Println("Sort by year", car)

	sort.Sort(OrderByName(car))

	fmt.Println("Sort by name", car)
	
	// Using sort.Slice, is more easy to use
	sort.Slice(car, func(i, j int) bool {
		return car[i].price > car[j].price
	})

	fmt.Println("Sort by big price with sort.Slice", car)

	// sortStringAndInt()

}

func sortStringAndInt() {
	ss := []string{"Golang", "Typescript", "Javascript", "Java", "Python"}
	sort.Strings(ss)
	fmt.Println(ss)

	ii := []int{5, 3, 8, 1, 2}
	sort.Ints(ii)
	fmt.Println(ii)

}
