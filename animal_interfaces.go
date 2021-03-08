package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
}

type Bird struct {
}

type Snake struct {
}

func (a Cow) Eat() {
	fmt.Println("grass")
}

func (a Cow) Move() {
	fmt.Println("walk")
}

func (a Cow) Speak() {
	fmt.Println("moo")
}

func (a Bird) Eat() {
	fmt.Println("worms")
}

func (a Bird) Move() {
	fmt.Println("fly")
}

func (a Bird) Speak() {
	fmt.Println("peep")
}

func (a Snake) Eat() {
	fmt.Println("mice")
}

func (a Snake) Move() {
	fmt.Println("slither")
}

func (a Snake) Speak() {
	fmt.Println("hsss")
}

func main() {
	animals := make(map[string]Animal)
	usage := "instruction example: 'newanimal <name> [cow|bird|snake]' or 'query <name> [eat|move|speak]'"

	fmt.Println("enter three word instruction, press q to quit, press h for help")
	for {
		fmt.Print("> ")
		inputReader := bufio.NewReader(os.Stdin)
		instruction, _ := inputReader.ReadString('\n')

		// clean up the input
		instruction = strings.TrimSuffix(instruction, "\n")
		instruction = strings.TrimSuffix(instruction, "\r")
		re := regexp.MustCompile(`^[ \t]+`)
		instruction = re.ReplaceAllString(instruction, "")
		re = regexp.MustCompile(`[ \t]+$`)
		instruction = re.ReplaceAllString(instruction, "")
		re = regexp.MustCompile(`[ \t]{2,}`)
		instruction = re.ReplaceAllString(instruction, " ")
		instruction = strings.ToLower(instruction)

		// quit
		if instruction == "q" {
			break
		}

		if instruction == "h" {
			fmt.Println(usage)
			continue
		}

		words := strings.Split(instruction, " ")

		if len(words) != 3 {
			fmt.Println("error: there must be three word instruction")
			continue
		}

		command := words[0]
		animalName := words[1]

		if command == "newanimal" {
			var animal Animal
			animalType := words[2]
			switch animalType {
			case "cow":
				animal = Cow{}
			case "bird":
				animal = Bird{}
			case "snake":
				animal = Snake{}
			default:
				fmt.Printf("error: unknown animal type '%s'\n", animalType)
				continue
			}

			animals[animalName] = animal
			fmt.Println("Created it!")
			continue
		}

		if command == "query" {
			animal, ok := animals[animalName]

			if !ok {
				fmt.Printf("Animal with name %s not found\n", animalName)
				continue
			}

			action := words[2]
			switch action {
			case "eat":
				animal.Eat()
			case "move":
				animal.Move()
			case "speak":
				animal.Speak()
			default:
				fmt.Printf("error: unknown action '%s'\n", action)
			}
			continue
		}

		fmt.Println(usage)
	}
}
