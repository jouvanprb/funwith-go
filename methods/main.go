package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animals struct {
	food       string
	locomotion string
	noise      string
}

func (a Animals) Eat() {
	fmt.Println("> Food eaten :", a.food)
}

func (a Animals) Move() {
	fmt.Println("> Locomotion method :", a.locomotion)
}

func (a Animals) Speak() {
	fmt.Println("> Spoken sound :", a.noise)
}

func main() {
	cow := Animals{
		food:       "grass",
		locomotion: "walk",
		noise:      "moo",
	}

	bird := Animals{
		food:       "worms",
		locomotion: "fly",
		noise:      "peep",
	}

	snake := Animals{
		food:       "mice",
		locomotion: "slither",
		noise:      "hsss",
	}

	fmt.Println("-- Program: Check Animals Behavior -- ")
	fmt.Println("-- Input Rules: (Animals Action)  -- ")

	for {
		fmt.Print("> ")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "x" || input == "X" {
			fmt.Println("> Program Stopped!")
			break
		}

		splitStr := strings.Split(input, " ")

		if len(splitStr) != 2 {
			fmt.Println("> Invalid Input")
			continue
		}

		animals := splitStr[0]
		action := splitStr[1]

		switch animals {
		case "cow":
			switch action {
			case "food":
				cow.Eat()

			case "locomotion":
				cow.Move()

			case "noises":
				cow.Speak()
			}

		case "bird":
			switch action {
			case "food":
				bird.Eat()

			case "locomotion":
				bird.Move()

			case "noises":
				bird.Speak()
			}

		case "snake":
			switch action {
			case "food":
				snake.Eat()

			case "locomotion":
				snake.Move()

			case "noises":
				snake.Speak()
			}
		}
	}
}
