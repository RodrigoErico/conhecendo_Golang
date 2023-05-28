package main

import "fmt"

func main() {
	height1 := 1.45
	var height2 = 1.89 
	var height3 float32
	var height4 float64
	var height5 float32 = 2.34
	
	fmt.Printf("%.3f\n",height1)
	fmt.Printf("%07.3f\n",height2) // 7 posiçoes.3 precisoes  001.890
	fmt.Printf("%.1f\n",height2) // redondou para 1.9
	fmt.Printf("%f\n",height3)
	fmt.Printf("%f\n",height4)
	fmt.Printf("%4.3f\n",height5) // tamanho.precisao
	fmt.Printf("%T\n\n", height5)
	
	calculator()
}

func calculator() {
	cents := 0.1
	cents += 0.2 // cents == 0.3
	fmt.Println(cents == 0.3) // false
	fmt.Println(cents)

	a := "10"
	b := "10"

	fmt.Println(a >= b)
}