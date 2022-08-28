package main

import "fmt"

func sub(a int, b int) int {
	return a - b
}

func main() {
	a := 100
	b := 10
	fmt.Println("The sub between a and b id:", sub(a, b))
}
