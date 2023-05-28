package main

import "fmt"

var a = "Sou texto" // escopo de pacote
// a := 1 => Short Variable Declaration Operator(:=), Não permitido no global.

func mostrarValorDeEscopoPacote() {
	fmt.Printf("Escopo global é do tipo %T\n",a)
}

func main() {
	var a = 10 // escopo de funcao. shadowing
	// a = 10 => Alterando o valor da variavel global(variavel pacote)
	fmt.Printf("Escopo de funcao é do tipo %T\n",a)
	mostrarValorDeEscopoPacote()
	variavel()
}

func variavel() {
	for b := 1; b <= 5; b++ {
		fmt.Println(b)
	}
	// fmt.Println(b) variavel não acessivel fora do escopo do loop
}