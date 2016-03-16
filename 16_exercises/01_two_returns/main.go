package main

import "fmt"

func main() {
	a, b := ex(5)
	fmt.Println(a, b)
	a, b = ex(6)
	fmt.Println(a, b)
}

func ex(n int) (float64, bool) {
	half := float64(n) / 2
	rem := n % 2
	return half, (rem == 0)
}
