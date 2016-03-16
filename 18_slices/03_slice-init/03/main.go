package main

import "fmt"

func main() {
	conMake := make([]int, 35) // Con make inicializamos el array subyacente y en este caso asignamos 35 de capacdad
	conMake[0] = 16
	fmt.Println(conMake)
	fmt.Println(conMake == nil) //Esto es falso

}
