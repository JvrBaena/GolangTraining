package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := Person{"James", "Bond", 20}
	marshaled, _ := json.Marshal(p1)
	fmt.Println(marshaled)
	fmt.Printf("%T \n", marshaled)
	fmt.Println(string(marshaled))
}
