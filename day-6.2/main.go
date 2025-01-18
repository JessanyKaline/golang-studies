package main

import "fmt"

func main() {
	aboutAnonymousFunction()
	aboutExpressionFunction()
	x := onlyOdd(add, []int{12, 13, 45, 67, 89, 55}...)
	fmt.Println("Result", x)
	aboutCallback()
	// fmt.Println("Factorial is:", (aboutRecursionFactorial(4)))
	// fmt.Println("Factorial is:", (aboutRecursionFactorial2(4)))
	// defer fmt.Println("Done")
	// y := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// fmt.Println(add2(y...))

	// c := circle{radius: 5.25}
	// s := square{side: 15.0}

	// fmt.Println(c.calArea())
	// fmt.Println(s.calArea())

	// info(c)

	// a := closure()
	// b := closure()

	// fmt.Println(a())
	// fmt.Println(a())
	// fmt.Println(a())

	// fmt.Println(b())
	// fmt.Println(b())




}

func aboutAnonymousFunction() {
	// Anonymous function
	func() {
		fmt.Println("Hello, playground")
	}()

	// Anonymous function with parameters
	func(x, y int) {
		fmt.Println(x + y)
	}(2, 3)

	// Anonymous function with return
	result := func(x, y int) int {
		return x + y
	}(2, 3)

	fmt.Println(result)
}

func aboutExpressionFunction() {
	// Expression function
	f := func() {
		fmt.Println("Hello, playground")
	}

	f()

	// Expression function with parameters
	f2 := func(x, y int) {
		fmt.Println(x + y)
	}

	f2(2, 3)

	// Expression function with return
	f3 := func(x, y int) int {
		return x + y
	}

	result := f3(2, 3)

	fmt.Println(result)
}

func aboutCallback() {
	// Callback is a function that is passed as an argument to another function
	callback := func(x, y int) int {
		return x + y
	}

	result := callback(2, 3)

	fmt.Println(result)
}

func add(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

func onlyOdd(f func(x ...int) int, y ...int) int {
	var slice []int
	for _, v := range y {
		if v%2 != 0 {
			slice = append(slice, v)
		}
	}

	total := f(slice...)
	return total

}

func aboutClosure() {
	// Closure is a function that captures the variables from the context in which it is defined
	x := 0
	increment := func() int {
		x++
		return x
	}

	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
}

func aboutRecursionFactorial(x int) int {
	if x == 0 {
		return 1
	}
	return x * aboutRecursionFactorial(x-1)

	//4! = 4 * 3 * 2 * 1 = 24
}

func aboutRecursionFactorial2(x int) int {
	if x == 1 {
		return x
	}
	return x * aboutRecursionFactorial(x-1)

	//4! = 4 * 3 * 2 * 1 = 24
}

func aboutRecursionFibonacci(x int) int {
	if x <= 1 {
		return x
	}
	return aboutRecursionFibonacci(x-1) + aboutRecursionFibonacci(x-2)
}

func aboutRecursionFactorialLoop(x int) int {
	total := 1
	for ; x > 0; x-- {
		total *= x
	}
	return total
}

func add2 (x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}


type circle struct {
	radius float64
}

type square struct {
	side float64
	
}

type shape interface {
	calArea() float64
}

func (c circle) calArea() float64 {
	return 2 * 3.14 * c.radius
}

func (s square) calArea() float64 {
	return s.side * s.side
}

func info(s shape) {
	fmt.Println("Info:", s.calArea())
}

func closure() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}