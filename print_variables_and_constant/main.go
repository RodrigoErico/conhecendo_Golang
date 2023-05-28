package main

import (
	"fmt"
	"math/rand"
)


var world = "World!"
const a = 10
const b = 34

func main() {
	//fmt.Print("Hello World!")
	//fmt.Println("Hello World!")
	fmt.Printf("Hello %s\n", world)
	fmt.Printf("2 + 6 = %v\n", 2 + 6)
	fmt.Printf("3 %% 6 = %v\n", 3 % 6)
	fmt.Printf("10 + 34= %v\n", a + b)

	var guessNumber = rand.Intn(10)
	fmt.Println("O número gerado foi: ", guessNumber)
}