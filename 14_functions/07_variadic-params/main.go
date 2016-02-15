package main

import "fmt"

func main() {
	n := average(23, 22, 54, 12, 33, 54)
	fmt.Println(n)
}

func average(sf ...float64) float64 { //nombre ... tipo permite definir parámetros variádicos (num variable)
	fmt.Println(sf)
	fmt.Printf("%T \n", sf)
	var total float64
	for _, v := range sf { // con range al iterar en un for devuelve índice, elemento
		total += v
	}
	return total / float64(len(sf)) // hay que castear a float64 para poder hacer la división
}
