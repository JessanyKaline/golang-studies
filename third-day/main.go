package main

import "fmt"

func main() {
	aboutFor()
	aboutForGeneralMonth()
	aboutBreak()
	aboutContinue()
	aboutNumberParity()
	aboutSwitch()
	aboutSwitchGeneral()
	aboutSwitchFallthrough()
	threeTimesUnicode()
	moduleFour()

}

// There are not while in Go, but you can use for loop to simulate it
func aboutFor() {
	// For loop
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("_______________")

	// For loop with condition
	j := 0
	for j <= 5 {
		fmt.Println(j)
		j++
	}

	fmt.Println("_______________")

	// For loop with range
	s := "Hello, World!"
	for k, v := range s {
		fmt.Println(k, string(v))
	}

	fmt.Println("_______________")

	// For loop with range and slice
	slice := []string{"a", "b", "c"}
	for k, v := range slice {
		fmt.Println(k, v)
	}

	fmt.Println("_______________")

	// For loop with range and map
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("_______________")
}

func aboutHowWhileWorks() {
	// While loop
	i := 2
	for i < 5 {
		fmt.Println(i)
		i++
	}
}

func aboutForGeneralMonth() {
	for i := 1; i <= 12; i++ {
		fmt.Println("\nMonth\t:", i)
		for i := 1; i <= 30; i++ {
			fmt.Print("\tDay\t", i)
			fmt.Println("_______________")
		}
	}
}

func aboutBreak() {
	x := 1
	for {
		if x < 5 {
			x++
			fmt.Println("Repetição:", x)
		} else {
			fmt.Println("End")
			break
		}
	}
	fmt.Println("Exit the loop")
}

func aboutContinue() {
	for i := 0; i <= 5; i++ {
		if i == 2 {
			// Skip the current iteration - not show the number 2
			continue
		}
		fmt.Println(i)
	}
}

func aboutNumberParity() {
	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, "is even")
		} else {
			fmt.Println(i, "is odd")
		}
	}
}

func aboutSwitch() {
	x := 7
	switch {
	case (x < 5), (x == 7):
		fmt.Println("x is less than 5, or is 7")
	case x > 5:
		fmt.Println("x is greater than 5")
	}
}

func aboutSwitchGeneral() {
	x := 5
	switch x {
	case 1:
		fmt.Println("x is 1")
	case 2:
		fmt.Println("x is 2")
	case 3:
		fmt.Println("x is 3")
	case 4:
		fmt.Println("x is 4")
	case 5:
		fmt.Println("x is 5")
	default:
		fmt.Println("x is not 1, 2, 3, 4 or 5")
	}
}

func aboutSwitchFallthrough() {
	x := 3
	switch x {
	case 1:
		fmt.Println("x is 1")
	case 2:
		fmt.Println("x is 2")
	case 3:
		// Skip the next comparison and always run the case 6
		fmt.Println("x is 3")
		fallthrough
	case 4:
		fmt.Println("x is 5")

	}
}

func threeTimesUnicode () {
	for i := 65 ; i <= 90; i ++{
		fmt.Println((i))
		for j := 0; j < 3; j++ {
			fmt.Printf("\t%#U\n", i)
		}
	}
}

func moduleFour(){
	for i := 10; i <= 100; i++ {
		fmt.Println( i % 4)
	}
}
