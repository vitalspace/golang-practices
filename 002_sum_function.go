package main

import "fmt"

func sum(a int, b int) int {
	return a + b
}

func main() {
	a := 10
	b := 40
	fmt.Println("The sum between a and b is", sum(a, b))
}
