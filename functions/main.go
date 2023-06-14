package main

import "fmt"

// func sum(a int, b int) (int, string) {
// 	// return a + b
// 	simpleReturn := a + b
// 	formattedReturn := fmt.Sprintf("%d + %d = %d", a, b, simpleReturn)
// 	return simpleReturn, formattedReturn
// }

func subtraction(x int, z int) (simpleReturn int, formattedReturn string) {
	defer fmt.Println("Retorna no fim da execução da função!")
	fmt.Println("Inicio da Função!")
	simpleReturn = x - z
	formattedReturn = fmt.Sprintf("%d + %d = %d", x, z, simpleReturn)
	return
}


func main() {
	simpleReturnResult, formattedReturnResult := subtraction(12, 10)
	fmt.Println(simpleReturnResult)
	fmt.Println(formattedReturnResult)
	fmt.Printf("O tipo de dado do formattedReturnResult é uma %T.\n", formattedReturnResult)
	fmt.Println(45_000_000)
	// _, formattedReturnResult:= sum(12, 10)
	// fmt.Println(simpleReturnResult)
	// fmt.Println(formattedReturnResult)
	// fmt.Printf("O tipo de dado do formattedReturnResult é uma %T.", formattedReturnResult)
}