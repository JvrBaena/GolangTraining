package main

import "fmt"

const metersToYards float64 = 1.09361

func main() {
	var meters float64
	fmt.Print("Introduce metros")
	fmt.Scan(&meters)
	yards := meters * metersToYards
	fmt.Println(meters, " metros son ", yards, " yardas")
}
