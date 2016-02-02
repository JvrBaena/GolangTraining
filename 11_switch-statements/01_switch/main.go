package main

import "fmt"

//En el switch no necesitamos poner break porque no hay fallthrough por defecto, hay que explicitarlo!
func main() {
	switch "Hello" {
	case "ONE":
		fmt.Println("ONE YAY")
	case "TWO":
		fmt.Println("TWO YAY")
	default:
		fmt.Println("Con String fijo se va a venir por aquí, cabezón")
	}

}
