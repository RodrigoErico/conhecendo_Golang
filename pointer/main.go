package main

import "fmt"

func main() {

	var valueOne int = 10
	var valueTwo int = valueOne
	
	fmt.Println(valueOne, valueTwo)

	valueOne++
	fmt.Println(valueOne, valueTwo)
	fmt.Println("-----------")

	var valueThree int
	var valueFourPointer *int

	valueThree = 100
	valueFourPointer = &valueThree

	valueThree++
	fmt.Println(valueThree, valueFourPointer)
	fmt.Println(valueThree, *valueFourPointer)

}