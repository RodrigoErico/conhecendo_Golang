package main

import "fmt"

func main() {
	var age uint = 23

	var countInt8 uint8 = 255
	countInt8++ // quando utrapassa o limite retorna ao começo que é zero 0

	var a uint8 = 255 // uint = 0 a 255 != int -128 a 127
	var b uint16 = 65535 // 2 ^ 16 = 65536 uint 0 a 65535
	var c int32
	var d int64
	var e uint

	fmt.Println(countInt8, age, a, b, c, d, e )
}