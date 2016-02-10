package main

import "fmt"

func main() {
	var name string
	fmt.Printf("Please enter your name: ")
	fmt.Scan(&name)
	fmt.Printf("Hello your name is %s", name)
}
