package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://lco.dev"
func main()  {
	res,err:=http.Get(url)
	chkNilError(err)
	// fmt.Println(res)
	fmt.Printf("Response type %T\n",res)

	defer res.Body.Close() // Caller's responsibility

	dataBytes,err:=io.ReadAll(res.Body)
	chkNilError(err)
	content:=string(dataBytes)
	fmt.Println(content)
}

func chkNilError(err error)  {
	if err != nil{
		panic(err)
	}
}