// Write a program that ask for a number and determines whether
// that number is even or odd and show it


package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calculate(n int) {
	if n%2 == 0 {
		fmt.Print("This number is even")
	} else {
		fmt.Print("This number is odd")
	}
}

func main() {
	fmt.Print("Enter a number: \n")
	var s string
	r := bufio.NewReader(os.Stdin)
	s, _ = r.ReadString('\n')
	number, _ := strconv.Atoi(strings.TrimSpace(s))
	calculate(number)
}
