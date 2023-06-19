package main

import "fmt"

func funcOne() {
	fmt.Println("Execuntando a função 1")
}

func funcTwo() {
	fmt.Println("Execuntando a função 2")
}

func studentIsApproved(numOne, numTwo float32) bool {
	fmt.Println("Aluno foi aprovado?")

	studentAverage := (numOne + numTwo) / 2

	if studentAverage >= 6 {
		return true
	}
	return false
}

func main() {
	// defer = Adiar
	defer funcOne()
	studentIsApproved(7, 8)
	funcTwo()
}
