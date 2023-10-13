package main

import (
	"fmt"
	"math/rand"
)

func main()  {
	diceNumber := rand.Intn(6)+1
	fmt.Println(diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice is 1 You can open")
	case 2:
		fmt.Println("You can mv 2 spots")
	case 3:
		fmt.Println("You can mv 3 spots")
		fallthrough // this runs the next case as well
	case 4:
		fmt.Println("You can mv 4 spots")
	case 5:
		fmt.Println("You can mv 5 spots")
	case 6:
		fmt.Println("You can mv 6 spots and roll dice again")
	default:
		fmt.Println("What was that!!")
	}
}