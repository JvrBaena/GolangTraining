package main

import "fmt"

func greatest(nums ...int) int {
	max := 0
	for _, n := range nums { // con range al iterar en un for devuelve índice, elemento
		if n > max {
			max = n
		}
	}
	return max
}

func main() {
	fmt.Println(greatest(1, 5, 2, 4, 99, 12))
	fmt.Println(greatest([]int{5, 5, 5, 88, 5}...))
}
