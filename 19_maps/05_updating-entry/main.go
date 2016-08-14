package main

import "fmt"

func main() {
	myGreeting := map[string]string{ //También podemos inicializar directamente aquí
		"Paco": "Hola",
		"John": "Hi!",
	}
	myGreeting["Pierre"] = "Bonjour!"
	fmt.Println(myGreeting)
	myGreeting["Pierre"] = "ARG!"
	fmt.Println(myGreeting)
}
