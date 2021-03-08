package main

import ( 
    "fmt"
    "reflect"
) 

func main() {
	arr := []interface{}{0, 2, []interface{}{[]interface{}{2, 3}, 8, 100, 4, []interface{}{[]interface{}{[]interface{}{50}}}}, -2}

	fmt.Println(reflect.TypeOf(arr).String())

	for i, elem := range arr {
		fmt.Println("Elem: ", i, elem, "|" , reflect.TypeOf(elem).Name(), "|")
	}
}
