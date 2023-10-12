package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome!!"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	fmt.Println(name)
}
