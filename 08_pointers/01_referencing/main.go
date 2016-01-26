package main

import "fmt"

func main() {
	a := 43

	fmt.Println(a)
	fmt.Println(&a)

	var b *int = &a //b es un puntero a entero *int

	fmt.Println(b)
}
