package main

import "fmt"

func main() {
	mySlice := make([]int, 0, 5) //Construimos un slice, basado en un array de int de capacidad 5, pero inicializado con 0 elementos
	fmt.Println("-------------------")
	fmt.Println(mySlice)
	fmt.Println(len(mySlice)) //Length será 0
	fmt.Println(cap(mySlice)) //Capacity será 5 (lo que le cabe al array en el que está basado)
	fmt.Println("-------------------")

	for i := 0; i < 80; i++ {
		mySlice = append(mySlice, i)
		fmt.Println("Len: ", len(mySlice), "Cap:", cap(mySlice), "Value:", mySlice[i])
	}
}
