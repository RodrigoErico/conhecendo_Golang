package main

import "fmt"

type user struct {
	name string
	age  uint8
}

func (userDate user) describesTheUser() {
	fmt.Printf("Usuario %s tem %d anos.\n", userDate.name, userDate.age)
}

func (userDate *user) changeUserData() {
	userDate.name = "Rodrigo"
	userDate.age = 31
}

func main() {
	userOne := user{"Pedro", 20}
	userTwo := user{"Maria", 40}

	userOne.describesTheUser()
	userTwo.describesTheUser()
	fmt.Println("Endereço da memória do primeiro usuario: ", &userOne.name)
	fmt.Printf("Usuario %s foi excluido!\n", userOne.name)

	fmt.Println(" ")
	userOne.changeUserData()
	fmt.Println("Endereço da memória do primeiro usuario: ", &userOne.name)
	fmt.Printf("Usuario %s foi adicionado!\n", userOne.name)
	userOne.describesTheUser()
	userTwo.describesTheUser()
}
