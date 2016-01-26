package main

import "fmt"

func zero(x *int) { // zero recibe un puntero a entero
	*x = 0 // modificamos el contenido de la dirección de memoria almacenada en x
}

func main() {
	x := 5
	zero(&x)
	fmt.Println(x) // x es 0, se ha modificado desde la función zero porque hemos pasado la dirección de memoria
}
