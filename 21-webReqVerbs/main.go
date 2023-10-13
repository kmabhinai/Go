package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main()  {
	// PerformGetRequest()
	// PerformPostRequest()
	PerformPostFormRequest()
}

func PerformGetRequest(){
	const myUrl = "http://127.0.0.1:8000/get"
	
	res,err := http.Get(myUrl)
	if err!=nil {
		panic(err)
	}
	defer res.Body.Close()
	
	fmt.Println("StatusCode: ",res.StatusCode)
	fmt.Println("Content Length: ",res.ContentLength)

	var resString strings.Builder
	content, _ := io.ReadAll(res.Body)
	byteCount,_:=resString.Write(content)
	
	fmt.Println("Byte Count is: ",byteCount)
	fmt.Println(resString.String())
	
	// fmt.Println(string(content))
}

func PerformPostRequest(){
	const myUrl = "http://127.0.0.1:8000/post"
	
	reqBody := strings.NewReader(`
	{
		"CourseName":"Go",
		"Price":1000
	}
	`)
	
	res,err:=http.Post(myUrl,"application/json",reqBody)
	if err!=nil{
		panic(err)
	}
	fmt.Println(res)
	defer res.Body.Close()
	content, _ := io.ReadAll(res.Body)
	fmt.Println(string(content))
}

func PerformPostFormRequest()  {
	const myUrl = "http://127.0.0.1:8000/postform"
	
	data:=url.Values{}
	data.Add("firstName","Abhinai")
	data.Add("lastName","Kambhampati")
	data.Add("email","kmabhinai@gmail.com")
	
	res,err := http.PostForm(myUrl,data)
	if err!=nil{
		panic(err)
	}
	defer res.Body.Close()
	content, _ := io.ReadAll(res.Body)
	fmt.Println(string(content))
}