package main

import "fmt"

type Square struct {
	side float64
}

func (s Square) area() float64 {
	return s.side * s.side
}

type Shape interface {
	area() float64
}

// printArea recibe una shape
func printArea(s Shape) {
	fmt.Println(s.area())
}

func main() {

	s := Square{5}
	fmt.Println(s)
	//Podemos pasar el Square porque la implementación de interfaces
	//es implícita: Si tiene el método area, ya es una shape
	printArea(s)
}
