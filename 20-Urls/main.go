package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://127.0.0.1:3000/auth?userType=admin&isActive=true" 
func main()  {
	fmt.Println(myUrl)

	result,_:= url.Parse(myUrl)
	
	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.Port())
	// fmt.Println(result.RawQuery)

	qParams := result.Query()
	fmt.Printf("The type of the query params are: %T\n",qParams)
	fmt.Println(qParams)

	for _,value := range qParams{
		fmt.Println("Params is: ",value)
	}

	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "127.0.0.1",
		Path: "/auth",
		RawPath: "user=abhinai",
	}

	fmt.Println(partsOfUrl.String())
}