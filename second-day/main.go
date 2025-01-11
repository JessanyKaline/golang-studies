package main

import "fmt"

// Constant only define your type when you are going to use it
const test = 10

// The type of the variable is defined when you assign a value to it
var a = 42

// Is possible declare multiple constants in the same line
const (
	h = 10
	p = 11
	j = 12
)

// When you declare a iota, the value will be incremented by 1, or more, depending on the number of constants declared
// And is possible not use the iota value with _
const (
	sequential  = iota + 10
	sequential2 = iota
	_           = iota
	sequential4 = iota
)

func main() {
	//fmt.Println(sequential, sequential2, sequential4)
	// conversionInByte()
	// aboutStrings()
	aboutDecimal()
	aboutBinary()
	aboutHexadecimal()
	// aboutChangeBytes()

}

func conversionInByte() {

	a := "e"
	b := "d"

	d := []byte(a)
	e := []byte(b)

	fmt.Printf("%v, %v\n", d, e)
}

func aboutStrings() {
	// For format with space
	s := `Hello, 
	World!`

	sbytes := []byte(s)

	fmt.Printf("%v, %T\n", s, s)
	fmt.Printf("%v, %T\n", sbytes, sbytes)

}

func aboutDecimal() {
	// Decimal and binary - 10
	x := 100
	fmt.Printf("decimal: %d - binary: %b\n", x, x)
}

func aboutBinary() {
	// Binary and decimal
	y := 200
	fmt.Printf("binary: %b - decimal: %d\n", y, y)
}

func aboutHexadecimal() {
	// Hexadecimal, binary and decimal
	z := 300
	fmt.Printf("hexadecimal: %#x - binary: %b - decimal: %d\n", z, z, z)
}

func aboutChangeBytes() {

	x := 4
	fmt.Printf("decimal: %d - binary: %b\n", x, x)

	y := x << 1
	fmt.Printf("decimal: %d - binary: %b\n", y, y)

	z := x >> 1
	fmt.Printf("decimal: %d - binary: %b\n", z, z)
}
