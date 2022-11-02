// Write a program that asks for a phrase or word and if the phrase or word is 4
// characters long, the program will concatenate an exclamation mark at the end, and if not
// is 4 characters long, the program will concatenate a question mark at the end. The
// program will then display the final sentence.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func calculate(s string) {
	phrase := strings.TrimSpace(s)
	if len(phrase) == 4 {
		fmt.Println(phrase + "!")
	} else {
		fmt.Println(phrase + "?")
	}
}

func main() {
	fmt.Print("Enter a phrase or world: \n")
	var s string
	r := bufio.NewReader(os.Stdin)
	s, _ = r.ReadString('\n')
	calculate(s)
}
