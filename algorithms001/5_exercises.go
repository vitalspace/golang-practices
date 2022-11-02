// Write a program that asks for a word or phrase and checks if the first letter of that phrase
// is an ?A?. If the first letter is an ?A?, a message must be printed on the screen
// that says "CORRECT", otherwise, "WRONG" should be printed. Note:
// investigate the SubString function of PseInt.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func calculate(s string) {
	if string(s[0]) == "a" || string(s[0]) == "A" {
		fmt.Print("CORRECT")
	} else {
		fmt.Print("WRONG")
	}
}

func main() {
	var s string
	fmt.Print("Enter a character: \n")
	r := bufio.NewReader(os.Stdin)
	s, _ = r.ReadString('\n')
	calculate(strings.TrimSpace(s))
}
