package main

import "fmt"

func main() {
	m := make([]string, 1, 25)
	fmt.Println(m) //[]
	changeMe(m)    // --> En este caso pasamos un array, que es un tipo referenciado
	//... por eso podemos modificar en changeMe el valor sin haber tenido que pasar la dirección de memoria
	//... es una dirección de memoria implícita
	fmt.Println(m) // ["Todd"]
}

func changeMe(z []string) {
	z[0] = "Todd"
	fmt.Println(z) // ["Todd"]
}
