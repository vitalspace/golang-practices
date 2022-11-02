// Write a program that asks for 3 scores and validates if those scores are between 1 and 10. If they are
// between these parameters a variable of logical type must be set to true and if not
// set it to false. At the end the program must say if the 3 notes are correct using the
// variable of logical type.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calculate(score1 int, score2 int, score3 int) {
	var state bool
	if score1 > 0 && score1 < 10 && score2 > 0 && score2 < 10 && score3 > 0 && score3 < 10 {
		state = true
		fmt.Println(state)
	} else {
		state = false
		fmt.Println(state)
	}
}

func main() {
	r := bufio.NewReader(os.Stdin)
	var sc1, sc2, sc3 string
	fmt.Print("Enter score 1: \n")
	sc1, _ = r.ReadString('\n')
	score1, _ := strconv.Atoi(strings.TrimSpace(sc1))
	fmt.Print("Enter score 2: \n")
	sc2, _ = r.ReadString('\n')
	score2, _ := strconv.Atoi(strings.TrimSpace(sc2))
	fmt.Print("Enter score 3: \n")
	sc3, _ = r.ReadString('\n')
	score3, _ := strconv.Atoi(strings.TrimSpace(sc3))
	calculate(score1, score2, score3)
}
