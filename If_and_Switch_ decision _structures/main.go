package main

import "fmt"


func main() {
	var age = 180

	if (age >= 18) {
		fmt.Printf("Usuario maior de idade\n")
	} else {
		fmt.Printf("Usuario maior de idade")
	}

	
	switch {
	case age == 18:
		fmt.Printf("Tem 18")

	case age > 18 && age <= 200:
		fmt.Print("Maior de 18")
	
	default:
		fmt.Print("Idade errada")
	}

}