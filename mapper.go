package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

func main() {

	m := make(map[string]int)

	path := "infile.txt"

	inFile, err := os.Open(path)
	defer inFile.Close()

	if err != nil {
		fmt.Println(err.Error() + `: ` + path)
		return
	}
	defer inFile.Close()

	scanner := bufio.NewScanner(inFile)
	for scanner.Scan() {
		line := scanner.Text()
		s := strings.Split(line, "|")
		m[s[3]] = 1
	}

	fmt.Println("num elements:", len(m))
}
