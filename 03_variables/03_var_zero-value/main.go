package main

import "fmt"

func main() {
	var a int //Así se asigna automáticamente el zero value del type. El valor por defecto vamos
	var b string
	var c float64
	var d bool

	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)
}
