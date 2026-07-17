package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct{}

type Bird struct{}

type Snake struct{}

// Cow object
func (c Cow) Eat() {
	fmt.Println("grass")
}

func (c Cow) Move() {
	fmt.Println("walk")
}

func (c Cow) Speak() {
	fmt.Println("walk")
}

// Bird Object
func (b Bird) Eat() {
	fmt.Println("worms")
}

func (b Bird) Move() {
	fmt.Println("fly")
}

func (b Bird) Speak() {
	fmt.Println("peep")
}

// Snake object
func (s Snake) Eat() {
	fmt.Println("mice")
}

func (s Snake) Move() {
	fmt.Println("slither")
}

func (s Snake) Speak() {
	fmt.Println("hsss")
}

func main() {

	zoo := make(map[string]Animal)

	for {
		fmt.Print(">")
		reader := bufio.NewReader(os.Stdin)

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "x" || input == "X" {
			fmt.Println("Stop Program!")
			break
		}

		formatSplit := strings.Split(input, " ")

		if len(formatSplit) != 3 {
			fmt.Println("Wrong input!")
			continue
		}
		command := formatSplit[0]
		name := formatSplit[1]

		switch command {
		case "newanimal":
			animalType := formatSplit[2]
			switch animalType {
			case "cow":
				zoo[name] = Cow{}
			case "bird":
				zoo[name] = Bird{}
			case "snake":
				zoo[name] = Snake{}
			default:
				fmt.Println("None of animals type")
			}
			fmt.Println("Created It!")
		case "query":
			action := formatSplit[2]
			pet, found := zoo[name]
			if !found {
				fmt.Println("Animal not found!")
				continue
			}
			switch action {
			case "eat":
				pet.Eat()
			case "move":
				pet.Move()
			case "speak":
				pet.Speak()
			}
		}
	}
}
