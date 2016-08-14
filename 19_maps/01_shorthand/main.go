package main

import "fmt"

func main() {
	myGreeting := make(map[string]string)
	myGreeting["Paco"] = "Hola"
	myGreeting["John"] = "Hi!"
	fmt.Println(myGreeting)
}
