package main

import "fmt"

func main() {
	// var ptr *int
	// fmt.Println("Val of pointer is", ptr)

	myNumber := 23
	var ptr = &myNumber
	fmt.Println(ptr)
	fmt.Println(*ptr)
}
