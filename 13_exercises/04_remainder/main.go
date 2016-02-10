package main

import "fmt"

func main() {
	fmt.Print("Please enter a number: ")
	var num1 int64
	fmt.Scan(&num1)
	fmt.Print("Please enter a second (larger) number: ")
	var num2 int64
	fmt.Scan(&num2)
	remainder := float64(num2 % num1)
	fmt.Printf("The remainder of %d/%d is %f", num1, num2, remainder)
}
