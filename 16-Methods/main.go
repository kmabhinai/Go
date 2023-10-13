package main

import "fmt"

type User struct {
	Name string
	Phone int
	UserName string
	IsActive bool
}

func (u User) getStatus()  {
	fmt.Println("Is User active ",u.IsActive)
}

func main()  {
	abhinai := User{"abhinai", 6305454460,"kmabhinai",true }
	fmt.Println(abhinai)
	fmt.Println(abhinai.Name)
	fmt.Printf("%+v\n",abhinai)
	abhinai.getStatus()
}