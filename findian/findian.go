package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("input:")

	// bufio.NewReader() = create a reader with a buffer
	// os.Stdin = keyboard (standard input)
	reader := bufio.NewReader(os.Stdin)

	// reader.ReadString('\n') = read input until the user presses ENTER (\n)
	input, _ := reader.ReadString('\n')

	// strings.TrimSpace() = removes whitespace at the beginning and end of a string ->  space, tab (\t), enter (\n), carriage return (\r), etc
	input = strings.TrimSpace(input)

	// strings.ToLower() = changes all letters to lowercase
	text := strings.ToLower(input)

	// Check that condition return True
	if strings.HasPrefix(text, "i") &&
		strings.HasSuffix(text, "n") &&
		strings.Contains(text, "a") {
		fmt.Println("Found!")
		// false condition
	} else {
		fmt.Println("Not Found!")
	}

}
