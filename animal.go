package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Animal struct {
	food string
	locomotion string
	noise string
}

func NewCow() Animal {
    return Animal{"grass", "walk", "moo"}
}

func NewBird() Animal {
    return Animal{"worms", "fly", "peep"}
}

func NewSnake() Animal {
    return Animal{"mice", "slither", "hsss"}
}

func (a Animal) Eat() string {
	return a.food
}

func (a Animal) Move() string {
	return a.locomotion
}

func (a Animal) Speak() string {
	return a.noise
}

func main() {
	fmt.Println("enter two word instruction, press q to quit, press h for help")
	for {
		fmt.Print("> ")
		inputReader := bufio.NewReader(os.Stdin)
		instruction, _ := inputReader.ReadString('\n')
		instruction = strings.TrimSuffix(instruction, "\n")
		instruction = strings.TrimSuffix(instruction, "\r")
		re := regexp.MustCompile(`^[ \t]+`)
		instruction = re.ReplaceAllString(instruction, "")
		re = regexp.MustCompile(`[ \t]+$`)
		instruction = re.ReplaceAllString(instruction, "")
		re = regexp.MustCompile(`[ \t]{2,}`)
		instruction = re.ReplaceAllString(instruction, " ")
		instruction = strings.ToLower(instruction)

		if instruction == "q" {
			break
		}

		if instruction == "h" {
			fmt.Println("instruction example: [cow|bird|snake] [eat|move|speak]")
			continue
		}
	
		words := strings.Split(instruction, " ")

		if len(words) != 2 {
			fmt.Println("error: there must be two word instruction")
			continue
		}

		var animal Animal
		animalName := words[0]
		switch animalName {
			case "cow":
				animal = NewCow()
			case "bird":
				animal = NewBird()
			case "snake":
				animal = NewSnake()
			default:
				fmt.Printf("error: unknown animal type '%s'\n", animalName)
				continue
		}

		action := words[1]
		switch action {
			case "eat":
				fmt.Println(animal.Eat())
			case "move":
				fmt.Println(animal.Move())
			case "speak":
				fmt.Println(animal.Speak())
			default:
				fmt.Printf("error: unknown action '%s'\n", action)
				continue
		}
	}
}
