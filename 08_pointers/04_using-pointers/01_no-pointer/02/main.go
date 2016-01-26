package main

import "fmt"

func zero(x int) {
	fmt.Printf("%p\n", &x) //la dirección de memoria de x dentro de zero...
	fmt.Println(&x)
	x = 0
}

func main() {
	x := 5
	fmt.Printf("%p\n", &x) //es distinta de la de x fuera
	fmt.Println(&x)
	zero(x)
	fmt.Println(x) // x es 5, no ha sido modificado
}
