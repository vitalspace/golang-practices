package main

import "fmt"

func multiply(a int, b int) int {
	return a * b
}

func main() {
	a := 30
	b := 2000

	fmt.Println("The multiplication between a and b is:", multiply(a, b))
}
