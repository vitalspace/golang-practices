package main

import "fmt"

func division(a int, b int) int {
	return a / b
}

func main() {
	a := 10000
	b := 2
	fmt.Println("The division between a and b is:", division(a, b))
}
