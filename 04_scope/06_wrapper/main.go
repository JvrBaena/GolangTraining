package main

import "fmt"

func wrapper() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func main() {
	increment := wrapper()
	fmt.Println(increment())
	fmt.Println(increment())
}

/*
Aquí, de una forma muy parecida a JS, podemos definir funciones que devuelven otras funciones... y el scope
de esa función que devuelve queda ligado al valor de la variable en el wrapper
*/
