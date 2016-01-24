package main

import "fmt"

func main() {
	a := 1 // Shorthand notation -> Declara la variable y la inicializa. La crea automáticamente del tipo que toca
	b := "lol"
	c := 90.2
	d := true

	fmt.Printf("%T \n", a)
	fmt.Printf("%T \n", b)
	fmt.Printf("%T \n", c)
	fmt.Printf("%T \n", d)
}
