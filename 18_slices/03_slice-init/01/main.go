package main

import "fmt"

/**
 *
 *Array, Map y Channel son tipos-referencia... tienen una referencia interna a una estructura de datos
 *subyacente. Por eso según los inicialicemos obtendremos una cosa u otra.
 *
 */
func main() {
	var conVar []string //inicializando con var NO estamos inicializando nada por debajo, tenemos nil
	fmt.Println(conVar)
	fmt.Println(conVar == nil) // esto es true Para poder añadir elementos habría que hacer append
}
