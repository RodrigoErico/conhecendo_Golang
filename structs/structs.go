package main

import "fmt"

type user struct {
	name string
	age uint8
	address address
}

type address struct {
	publicPlace string
	number uint8
}

func main() {

	var userOne user
	userOne.name = "Pedro"
	userOne.age = 33
	fmt.Println(userOne)

	addressMaria := address{"Rua Capivara Alegre", 12}
	userTwo := user{"Maria", 21, addressMaria}
	fmt.Println(userTwo)

	userThree := user{age: 56}
	fmt.Println(userThree)
}