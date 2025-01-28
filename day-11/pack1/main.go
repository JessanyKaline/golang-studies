package main

//Is necessary the "go mod init day-11" for Go import correctly, then create go.mod
// Run is possible with "go run pack1/*.go"
import (
	"fmt"
	"day-11/pack2"
)

func main(){
	fmt.Println("main")
	other()
	pack2.Pack2()
}