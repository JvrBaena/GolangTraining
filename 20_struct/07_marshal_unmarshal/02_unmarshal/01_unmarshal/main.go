package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First string
	Last  string
	Age   int
}

func main() {
	var p1 Person

	bs := []byte(`{"First": "James", "Last": "Bond", "Age": 20}`)
	json.Unmarshal(bs, &p1)
	fmt.Println(p1)
}
