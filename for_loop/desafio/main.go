package main

import (
	"fmt"
	"math/rand"
)

func advinhaNumero() {
	var meuNumero = 39

	for {
		var tentativa = rand.Intn(99) + 1
		if tentativa < meuNumero {
			fmt.Println("Numero é maior", tentativa)
		} else if tentativa > meuNumero {
			fmt.Println("Numero esta acima", tentativa)
		} else {
			fmt.Println("Voce acertou!", tentativa)
			break
		}
	}
}

func main() {
	advinhaNumero()
}
