package main

import "fmt"

func main()  {
	var userName string = "Abhinai"
	fmt.Println(userName)
	fmt.Printf("Type :- %T\n",userName)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Type :- %T\n",isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Type :- %T\n",smallVal)

	var smallFloat float32 = 255.3254545545
	fmt.Println(smallFloat)
	fmt.Printf("Type :- %T\n",smallFloat)
	
	//default values
	var tempVar int
	fmt.Println(tempVar)
	fmt.Printf("Type :- %T\n",tempVar)

	//implicit type
	var website = "google.com"
	fmt.Println(website)

	//no var style
	numberOfUsers :=300000 //this method can't be used globally
	fmt.Println(numberOfUsers)
}