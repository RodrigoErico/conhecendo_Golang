package main

import (
	"fmt"
	"sync"
	"time"
)


func main() {

	var waitGroup sync.WaitGroup

	waitGroup.Add(2)

	go func ()  {
		write("Olá mundo!")
		waitGroup.Done() // -1
	}()
	
	go func ()  {
		write("Programando em Go!")
		waitGroup.Done() // -1
	}()

	waitGroup.Wait() // 0
}

	func write(text string) {
	for i := 0; i < 5; i++ {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
