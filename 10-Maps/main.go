package main

import "fmt"

func main()  {
	languages:=make(map[string]string)
	languages["Js"]="javascript"
	languages["Py"]="Python"
	languages["rb"]="ruby"

	fmt.Println(languages)
	
	delete(languages,"Js")
	fmt.Println(languages)

	for key,value:= range languages{
		fmt.Printf("%v - %v\n",key,value)
	}
}