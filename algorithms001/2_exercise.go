// Write a program that asks to enter only phrases or words of 6 characters. If he
// user enters a 6-character word or phrase, a message should be printed
// by screen that says "CORRECT", otherwise, it should be printed
// "INCORRECT".

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func calculate(n int) {
	if n > 6 {
		fmt.Print("CORRECT")
	} else {
		fmt.Print("INCORRECT")
	}
}

func main() {
	fmt.Print("Enter only phrases or words of 6 characters: \n")
	var s string
	r := bufio.NewReader(os.Stdin)
	s, _ = r.ReadString('\n')
	trim := strings.TrimSpace(s)
	calculate(len(trim))
}
