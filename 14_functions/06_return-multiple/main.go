package main

import "fmt"

func main() {
	fmt.Println(greet("Jane", "Doe"))
}

func greet(fname, lname string) (string, string) { //Devolvemos valores múltiples
	return fmt.Sprint(fname, lname), fmt.Sprint(lname, fname)
}
