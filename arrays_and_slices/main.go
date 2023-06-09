package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e Slices")

	var arrayOne[5] string
	arrayOne[0] = "Posição 1"
	fmt.Println(arrayOne)

	arrayTwo := [3]string{"Posição 1", "Posição 2", "Posição 3"}
	arrayTwo[1] = "Nova Posição 2"
	fmt.Println(arrayTwo)

	arrayThree := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(arrayThree)
	
	fmt.Println("Slices:")
	
	slice := []int{12, 23, 33, 44}
	slice = append(slice, 100) // 100
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(arrayThree))
	fmt.Println(reflect.TypeOf(slice))
	//fmt.Printf("%T", slice)

	fmt.Println("---Arrays internos----")
	// Arrays internos
	sliceTwo := make([]float32, 10, 11)
	fmt.Println(sliceTwo)
	fmt.Println(len(sliceTwo)) // tamanho 10
	fmt.Println(cap(sliceTwo)) // capacidade 11

	sliceTwo = append(sliceTwo, 33)
	sliceTwo = append(sliceTwo, 1) 
	
	// capacidade maxima é ultrapassada o go multiplica o tamanho.

	fmt.Println(sliceTwo)
	fmt.Println(len(sliceTwo)) // tamanho = 12
	fmt.Println(cap(sliceTwo)) // capacidade = 24

}