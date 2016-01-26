package main

import "fmt"

func zero(x *int) {
	fmt.Println("Dentro")
	fmt.Println(x)
	fmt.Println(*x)
	*x = 0
}

func main() {
	x := 5
	fmt.Println("Antes")
	fmt.Println(x)
	fmt.Println(&x)
	zero(&x) // Pasamos la dirección de memoria, es la misma en todo el recorrido
	fmt.Println("Después")
	fmt.Println(&x)
	fmt.Println(x) // x es 5, no ha sido modificado
}
