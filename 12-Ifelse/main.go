package main

import (
	"fmt"
)

func main()  {
	var age int
	fmt.Scanf("%d",&age)
	fmt.Println(age)
	if age>=18 {
		fmt.Println("Eligible for Vote")
	} else {
		fmt.Println("Not Eligible For Vote")
	}
}