package main

import "fmt"


func recoverExecution()  {
	if r := recover(); r != nil {
		fmt.Println("Programa recuperado!")
	}
}

func calculateAverage(noteOne, notetwo float32) bool {

	defer recoverExecution()


	averageResult := (noteOne + notetwo) / 2

	if averageResult > 6 {
		return true
	} else if averageResult < 6 {
		return false
	}

	panic("A media é extamente 6!")
}

func main() {
	fmt.Println(calculateAverage(6, 6))
	fmt.Println("Terminou o programa!")
}