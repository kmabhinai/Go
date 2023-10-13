package main

import "fmt"

type User struct {
	Name string
	Phone int
	UserName string
	IsVerified bool
}
func main()  {
	abhinai := User{Name:"abhinai",Phone:  6305454460,UserName: "kmabhinai",IsVerified: true }
	// abhinai := User{"abhinai", 6305454460,"kmabhinai",true }
	fmt.Println(abhinai)
	fmt.Println(abhinai.Name)
	fmt.Printf("%+v\n",abhinai)
}