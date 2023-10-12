package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating of the Pizza!")
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating ", input)

	/* this involves the \n in the end of the string*/
	// numRating, err := strconv.ParseFloat(input, 64)
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(numRating + 1)
	}
}
