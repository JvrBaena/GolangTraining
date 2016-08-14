package main

import "fmt"

func main() {
	myGreeting := map[int]string{
		0: "Blau",
		1: "Blau, Blau",
		2: "Blau, Blau, Blau",
		3: "Blau, Blau, Blau, Blau",
	}

	fmt.Println(myGreeting)

	if val, exists := myGreeting[2]; exists { //El acceso por índice en un map devuelve val,exists...es un patrón típico en go
		delete(myGreeting, 2)
		fmt.Println(val)
		fmt.Println(exists)
	} else {
		fmt.Println("That Value does not exist")
		fmt.Println(val)
		fmt.Println(exists)
	}

	fmt.Println(myGreeting)
}
