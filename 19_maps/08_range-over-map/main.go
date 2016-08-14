package main

import "fmt"

func main() {
	myGreeting := map[int]string{
		0: "Blau",
		1: "Blau, Blau",
		2: "Blau, Blau, Blau",
		3: "Blau, Blau, Blau, Blau",
	}

	for key, val := range myGreeting {
		fmt.Println("key: ", key, " val: ", val)
	}
}
