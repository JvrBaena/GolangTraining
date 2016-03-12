package main

import "fmt"

func main() {
	exp := func(n int) (float64, bool) {
		half := float64(n) / 2.0
		rem := n % 2
		return half, (rem == 0)
	}

	a, b := exp(5)
	fmt.Println(a, b)
	a, b = exp(6)
	fmt.Println(a, b)
}
