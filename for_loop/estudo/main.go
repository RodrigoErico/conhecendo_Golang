package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i <= 5; i++ {
		fmt.Println("Carregou...", i * 20,"%")

		time.Sleep(700 * time.Millisecond)
	}
	fmt.Println("Terminou o carregamento.")
}