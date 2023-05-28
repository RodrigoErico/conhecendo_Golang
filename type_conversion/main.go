package main

import (
	"fmt"
	"strconv"
)

func main() {

	// converte numero para string
	amount := 2

	message := fmt.Sprintf("Voce adicionou %d itens!\n", amount)
	// or
	message2 := "Voce adicionou " + strconv.Itoa(amount) + " itens!"
	fmt.Println(message, message2)
	fmt.Println("/|/|/|/|/|/|/|/|/|/|/|")

	// converte string em numero
	weight := "115" // "115d" retorna como 0 apos a conversao
	weightInInt, _ := strconv.Atoi(weight)
	fmt.Printf("weight agora é do tipo: %T\n", weightInInt)
	fmt.Println(weightInInt)
}