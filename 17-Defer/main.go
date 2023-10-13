package main

import "fmt"

//Defer places the line to be executed at the last if more than one defer exists then it maintains stack ie. it follows LIFO

func main()  {
	defer fmt.Println("Three")
	defer fmt.Println("Two")
	defer fmt.Println("One")
	defer fmt.Println("World")
	fmt.Println("Hello")
	myDefer()
}

func myDefer()  {
	for i:=0; i<5; i++{
		defer fmt.Println(i)
	}
}