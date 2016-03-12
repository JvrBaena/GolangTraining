package main

import "fmt"

func hello() {
	fmt.Print("Hello ")
}

func world() {
	fmt.Println("World")
}

func main() {
	defer world() //defer retrasa la ejecución de esta función hasta justo antes de que la función en que se encuentra (main) salga
	hello()
}
