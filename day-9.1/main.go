package main

import (
	"encoding/json"
	"fmt"
	"sort"
)

func main (){
	//exerciseMarshal()
	//exerciseUnmarshal()
	//exerciseSort()
	orderBy()

}

func exerciseMarshal(){
	type User struct {
		Name string
		Age int
	}

	u1 := User{"John", 30}
	u2 := User{"Jane", 25}

	users := []User{u1, u2}
	fmt.Println(users)

	jsU1, err := json.Marshal(u1)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsU1))

	jsU2, err := json.Marshal(u2)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsU2))

	sb, err := json.Marshal(users)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(sb))

}

func exerciseUnmarshal(){
	type User struct {
		Name string
		Age int
	}

	jsonUser := []byte(`{"Name":"John","Age":30}`)

	var user User

	err := json.Unmarshal(jsonUser, &user)

	if err != nil {
		panic(err)
	}

	fmt.Println(user)
}

func exerciseSort(){
	x := []int{4, 7, 3, 42, 99, 18, 5, 1, 6, 8}

	sort.Ints(x)
	fmt.Println(x)

	sort.Sort(sort.Reverse(sort.IntSlice(x)))

	fmt.Println(x)
}

func orderBy(){
	type user struct {
		First   string
		Last    string
		Age     int
		Sayings []string
	}

	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	sort.Slice(users, func(i, j int) bool {
		return users[i].Age < users[j].Age
	})

	for _, user := range users {
		fmt.Println("Order by age:", user.First, user.Last, user.Age)
	}

	sort.Slice(users, func(i, j int) bool {
		return users[i].Last < users[j].Last
	})

	for _, user := range users {
		fmt.Println("Order by lastName:", user.First, user.Last, user.Age)
	}
}