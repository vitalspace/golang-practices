package main

import "fmt"

func rest(a int, b int) int {
	return a % b
}

func main() {
	a := 20
	b := 5
	fmt.Println("The Rest between a and b is:", rest(a, b))
}
