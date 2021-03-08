package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Person struct {
	fname string
	lname string
}
func main() {
	var fileName string
	fmt.Println("Enter a name of the text file (e.g. file.txt):")
	_, _ = fmt.Scan(&fileName)
	file, _ := os.Open(fileName)
	scanner := bufio.NewScanner(file)
	var persons []Person
	for scanner.Scan() {
		line := scanner.Text()
		name := strings.Fields(line)
		person := Person{
			fname: name[0],
			lname: name[1],
		}
		persons = append(persons, person)
	}
	_ = file.Close()
	for _, el := range persons {
		fmt.Printf("first name: %s, last name: %s\n", el.fname, el.lname)
	}
}