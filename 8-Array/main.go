package main

import "fmt"

func main() {
	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Orange"
	fruitList[3] = "Mango"

	fmt.Println(fruitList)

	var vegList = [3]string{"potato", "potato2"}
	fmt.Println(len(vegList))
}
