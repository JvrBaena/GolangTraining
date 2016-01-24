package main

import "fmt"

func main() {
	fmt.Printf("%d - %b - %X", 42, 42, 42)
	fmt.Printf("%d - %b - %#X", 42, 42, 42) //# añade el prefijo 0x a hex, 0 a octo...etc
}
