package main

import "fmt"

func main() {
	greet("Paco") //Pasas argumentos a la función
	greet("Pepe")
}

func greet(name string) { //Declaras greet con un parámetro de tipo string
	fmt.Println(name)
}
