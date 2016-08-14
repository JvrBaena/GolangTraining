package main

import "fmt"

func main() {
	var myGreeting = make(map[string]string) // map[tipokey]tipovalue también podríamos inicializarlo con map[string]string{}
	myGreeting["Paco"] = "Hola"
	myGreeting["John"] = "Hi!"
	fmt.Println(myGreeting)
}
