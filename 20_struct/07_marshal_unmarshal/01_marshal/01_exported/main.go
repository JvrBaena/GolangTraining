package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First       string
	Last        string
	Age         int
	notExported int
}

func main() {
	p1 := Person{"James", "Bond", 20, 007}
	marshaled, _ := json.Marshal(p1)
	fmt.Println(marshaled)
	fmt.Printf("%T \n", marshaled)
	fmt.Println(string(marshaled))
}
