package main

import "fmt"

func wrapper() func() int {
	var x int           //es más idiomático esto para declarar zero value, que hacer := 0
	return func() int { //Al devolver la función, ejecutando wrapper, seteamos x a 0 y la función
		//anónima que devolvemos queda ligada a x por el scope
		x++
		return x
	}
}

func main() {
	increment := wrapper()
	fmt.Println(increment())  //1
	fmt.Println(increment())  //2
	increment2 := wrapper()   // al ejecutar again wrapper, x es una nueva variable que queda seteada 0 y ligada
	fmt.Println(increment2()) //1
}
