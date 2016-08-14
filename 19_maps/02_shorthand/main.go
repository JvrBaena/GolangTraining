package main

import "fmt"

func main() {
	myGreeting := map[string]string{ //También podemos inicializar directamente aquí
		"Paco": "Hola",
		"John": "Hi!",
	}

	fmt.Println(myGreeting)
}
