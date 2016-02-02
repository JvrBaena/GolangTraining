package main

import "fmt"

//En el switch no necesitamos poner break porque no hay fallthrough por defecto, hay que explicitarlo!
func main() {
	switch "Hello" {
	case "ONE":
		fmt.Println("ONE YAY")
	case "TWO":
		fmt.Println("TWO YAY")
	case "Hello":
		fmt.Println("AQUÍ POZI")
		fallthrough //Si comentas esto, aunque matchee "Hello" no va a ejecutar el default
	default:
		fmt.Println("ESTE TAMBIÉN POR EL FALLTHROUGH")
	}

}
