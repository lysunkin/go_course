package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const maxLength = 20

type Name struct {
	fname string
	lname string
}

func main() {
	var fn string
	fmt.Print("Input file name: ")
	fmt.Scanln(&fn)

	file, err := os.Open(fn)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	data := make([]Name, 1)

	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
	        if len(line) == 0 {
			continue
		}
		s := strings.Split(line, " ")

	        if len(s) < 2 {
			continue
		}

                firstName := s[0]
		if len(firstName) > maxLength {
			firstName = firstName[:maxLength]
		}

                lastName := s[1]
		if len(lastName) > maxLength {
			lastName = lastName[:maxLength]
		}

		n := Name{ fname: firstName, lname: lastName }
		
		if len(data) > i {
			data[i] = n
		} else {
			data = append(data, n)
		}

		i++
	}

	for i, v := range data {
		fmt.Printf("Item# %d, First name: %s, Last Name: %s\n", i, v.fname, v.lname)
	}
}
