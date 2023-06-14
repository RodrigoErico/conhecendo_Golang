package main

import "fmt"

type person struct {
	name string
	lastName string
	gender string
	age uint8

}

type student struct {
	person //person  // embedding
	course string
	courseLocation string
}

func main() {

	personOne := person{"Maria", "Alice", "feminino", 22}
	fmt.Println(personOne)

	studentOne := student{personOne, "BackEnd em Go", "Udemy"}
	fmt.Println(studentOne)
	
	//fmt.Println(studentOne.person.name)
	
}