package main

import "fmt"

// := is used to declare and initialize a variable
// Examples keywords in go: package, func, var, const, etc
// Go is a statically typed language, so you need to declare the type of the variable

var z = "variable_level_pck" // This is a package level variable - global in file
var b = 42
var c = "test"
var a string //This declaration is a package level variable - global in file, type string
var isFalse bool
var g int

func main() {
	generalFormatting()
	simpleType()
	exercises()

	s := fmt.Sprintf("%v\t%v\t%v", z, b, c) // Sprintf is used to format a string
	fmt.Println(s)

}

func generalFormatting() {
	numberbytes, er := fmt.Println("Hello, World!")

	fmt.Println("Numbers of bytes: ", numberbytes, "Erro: ", er)

	//For example if you don't want to use the variable numberbytes, you can use the _ to ignore it

	x := 10
	y := "test"

	fmt.Printf("x: %v, y: %v\n", x, y)
	fmt.Printf("x: %T, y: %T\n", x, y) // %T is used to print the type of the variable

	fmt.Printf("a: %v, %T\n", a, a)

	x = 20 // Reassigning the value of x

	interpretedString := "oi\ntudo ok?"
	rawString := `"oi\ntudo ok?"` // Raw or literal string, it will print the string as it is

	fmt.Println(interpretedString)
	fmt.Println(rawString)
	fmt.Sprint(interpretedString, rawString) // Sprint is used to concatenate strings

}

func simpleType() {
	type iceCream int

	var test iceCream = 2

	x := 10

	fmt.Printf("%T\n", test)
	fmt.Printf("%T\n", x)

	conversion := iceCream(x) // Converting x to iceCream type

	fmt.Printf("%T\n", test)
	fmt.Printf("%T\n", conversion)

}

func exercises() {
	infoInit := "--------Init exercises------"
	fmt.Println(infoInit)
	type hamburguer int
	var x hamburguer
	fmt.Println(x)
	x = 44
	fmt.Println(x)

	g = int(x)
	fmt.Printf("%T\t%v\n", g, g)

	infoEnd := "--------End exercises------"
	fmt.Println(infoEnd)
}
