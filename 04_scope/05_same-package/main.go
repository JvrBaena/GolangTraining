package main

import "fmt"

func main() {
	var x = 0
	increment := func() int {
		x++
		return x
	}

	fmt.Println(increment())
	fmt.Println(increment())
}

/*
Para declarar funciones dentro de otras funciones, lo hacemos de forma anónima, con closures.
Así las funciones pueden compartir variables sin que éstas tengan que estar en el nivel de package.
*/
