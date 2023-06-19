package main

import "fmt"

func functionClosure() func() int {
	value := 0
	return func() int {
		value += 1
		return value
	}
}

func main() {
	newClosureOne := functionClosure()
	fmt.Println(newClosureOne())
	fmt.Println(newClosureOne())
	fmt.Println(newClosureOne())

	newClosureTwo := functionClosure()
	fmt.Println(newClosureTwo())
	fmt.Println(newClosureTwo())
}
