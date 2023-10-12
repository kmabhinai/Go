package main

import (
	"fmt"
	"math/big"

	// "math/rand"
	"crypto/rand"
)

func main() {
	//random number
	// fmt.Println(rand.Intn(5) + 1)

	//crypto
	myRandomNumer, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myRandomNumer)
}
