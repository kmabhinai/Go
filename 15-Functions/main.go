package main

import "fmt"

func main()  {
	fmt.Println(addTwo(3,5))
	fmt.Println(addMany(2,3,5,4,6))
}

func addTwo(a int, b int) int {
	return a+b
}

func addMany(values ...int)(int,string) {
	total :=0
	for _, value := range values{
		total+=value
	} 
	return total, "eww"
}