package main

import "fmt"

var MyVar string = "blablabla" // Declaración e inicialización

func main() {
	a := 1 // Shorthand notation -> Declara la variable y la inicializa. La crea automáticamente del tipo que toca
	b := "lol"
	c := 90.2
	d := true

	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)
}
