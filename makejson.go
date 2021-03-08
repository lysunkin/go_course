package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	var name string
	fmt.Print("Input the name: ")
	fmt.Scanln(&name) 
	fmt.Print("Input the address: ")
	inputReader := bufio.NewReader(os.Stdin)
	address, _ := inputReader.ReadString('\n')
	address = strings.TrimSuffix(address, "\r\n")

	var data = make(map[string]string)
	data["name"] = name 
	data["address"] = address 

	jsonData, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	fmt.Println("JSON data: ", string(jsonData))
}
