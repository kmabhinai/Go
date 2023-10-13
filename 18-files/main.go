package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main()  {
	content := "Abhinai"
	file,err:=os.Create("./myName.txt")
	chkNilError(err)
	
	length,err:=io.WriteString(file,content)
	chkNilError(err)
	fmt.Println("length is: ",length)
	defer file.Close()
	readFile("./myName.txt")
}

func readFile(fileName string)  {
	dataByte,err:=ioutil.ReadFile(fileName)
	chkNilError(err)
	fmt.Println("The text data in file is : \n",dataByte)
	fmt.Println("The text data in file is : ",string(dataByte))
}

func chkNilError(err error)  {
	if err != nil{
		panic(err)
	}
}