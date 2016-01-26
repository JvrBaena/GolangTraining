package main

import "fmt"

func zero(x int) { // Pasamos x por "copia" (sólo se pasa el valor, no la variable actual, recibe su propia copia)
	x = 0 // al modificar x, no se está modificando la x de main, sino el parámetro de esta función
}

func main() {
	x := 5
	zero(x)
	fmt.Println(x) // x es 5, no ha sido modificado
}
