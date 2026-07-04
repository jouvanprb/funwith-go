package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var input string

func main() {
	// To mark the letters found in the loop!
	state := 0

	// For user input in terminal
	fmt.Print("Input: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	// Make all str to lowercase 
	text := strings.ToLower(input)

	// make the output to lowercase and Remov all space (We aren't use this one because it can't return the space)
	// output := strings.ReplaceAll(strings.ToLower(input)," ", "")
	// fmt.Println(output)

	// loop every char bytes in text
	for _, ch := range text {
		if ch == ' ' {
			continue
		}

		// condition for every char found
		if ch == 'i' && state == 0 {
			state = 1
		} else if ch == 'a' && state == 1 {
			state = 2
		} else if ch == 'n' && state == 2 {
			state = 3
			break
		}
	}

	// Condition that all statement Condition that all statements have been found
	if state == 3 {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}