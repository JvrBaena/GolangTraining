package main

import "fmt"

func main() {
	a := 43

	fmt.Println(a)
	fmt.Println(&a)

	var b *int = &a //b es un puntero a entero *int, contiene la dirección de memoria donde hay un int

	fmt.Println(b)
	fmt.Println(*b) //con *b obtenemos el valor almacenado en la dirección de memoria que hay en b

	*b = 666       // Modifiquemos el contenido de la dirección de memoria almacenada en b (a)
	fmt.Println(a) //666
}
