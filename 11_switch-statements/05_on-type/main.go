package main

import "fmt"

type Contact struct {
	greeting string
	name     string
}

func switchOnType(x interface{}) {
	switch x.(type) { //Esto es un assert. Se aserta 'x es tipo tal'
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case Contact:
		fmt.Println("Contact")
	default:
		fmt.Println("wtf??")
	}
}

func main() {
	switchOnType(7)
	switchOnType("BLABLABLA")
	var t = Contact{"JAJAJAJA", "PACO"}
	switchOnType(t)
}
