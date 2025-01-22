package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//aboutTransformJson()
	aboutUnmarshalJson()

}

// In go is necessary for exported fields to start with a CAPITAL letter
type Person struct {
	Name       string
	LastName   string
	Age        int
	Profession string
	Wage       float64
}

func aboutTransformJson() {

	johnSnow := Person{Name: "John", LastName: "Snow", Age: 30, Profession: "King of the North", Wage: 1000.50}
	jessanyKaline := Person{Name: "Jessany", LastName: "Kaline", Age: 25, Profession: "Queen of the South", Wage: 2000.00}

	js, err := json.Marshal(johnSnow)
	if err != nil {
		{
			fmt.Println(err)
		}
	}
	jk, err := json.Marshal(jessanyKaline)
	if err != nil {
		{
			fmt.Println(err)
		}
	}

	fmt.Println(string(js), "\n", string(jk))
}

func aboutUnmarshalJson() {
    //In go is necessary use `json:"Name"` to map the json key to the struct field
	//In go is necessary to use []byte to convert a string to a byte slice
	//In go is necessary to use json.Unmarshal to convert a json to a struct
	//In go is necessary to use a struct to convert a json to a struct
	type Animal struct {
		Name    string `json:"Name"`
		Species string `json:"Species"`
		Color   string `json:"Color"`
		Age     int    `json:"Age"`
	}

	jsonAnimals := []byte(`[{"Name":"Dog","Species":"Canis","Color":"Black","Age":5}, {"Name":"Cat","Species":"Felis","Color":"White","Age":3}]`)

	var animals []Animal
	err := json.Unmarshal(jsonAnimals, &animals)
	if err != nil {
		fmt.Println(err)
	}

	for _, animals := range animals {
		fmt.Println(animals.Name, animals.Species)
	}
}
