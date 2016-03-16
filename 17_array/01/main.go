package main

import "fmt"

func main() {
	var x [58]string //La longitud del Array es parte del tipado. Inicializa a 58 strings (cada uno al valor zero)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[42])
	for i := 65; i <= 122; i++ {
		x[i-65] = string(i)
	}
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[42])
}
