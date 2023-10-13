package main

import (
	"encoding/json"
	"fmt"
)

type course struct{
	Name string `json:"courseName"`   // aliases can be used to rename the keys in json
	Price int
	Platform string `json:"website"`
	Password string `json:"-"` // - excludes the felid
	Tags []string `json:"tags,omitempty"`  //omitempty excludes the nil values
}

func main()  {
	// EncodeJson()
	DecodeJson()
}

func EncodeJson()  {
	knlsCourses := []course{
		{"React",100,"Knls.com","abc123",[]string{"Web-dev","Frontend"}},
		{"Node",500,"google.com","def456",[]string{"abc","def"}},
		{"Angular",1000,"angular.com","wqe7890",nil},
	}

	// finalJson,err := json.Marshal(knlsCourses)
	finalJson,err := json.MarshalIndent(knlsCourses,"","\t")  // "" gives the starting word of the line and \t is for space separation
	if err!=nil{
		panic(err)
	}
	fmt.Printf("%s\n",finalJson)
}

func DecodeJson()  {
	jsonDataFromWeb := []byte(`
		{
			"courseName": "React",
			"Price": 100,
			"website": "Knls.com",
			"tags": ["Web-dev","Frontend"]
		}
	`)

	var knlsCourse course
	checkValid:=json.Valid(jsonDataFromWeb)

	if checkValid{
		fmt.Println("Json is Valid")
		json.Unmarshal(jsonDataFromWeb,&knlsCourse)
		fmt.Printf("%#v\n",knlsCourse)
		} else{
			fmt.Println("Json is not valid!!")
		}
		
		var myOnlineData map[string]interface{}
		json.Unmarshal(jsonDataFromWeb,&myOnlineData) 
		fmt.Printf("%#v\n",knlsCourse)
		
		for key,val := range myOnlineData{
			fmt.Printf("Key is %v and value is %v Type is %T\n",key,val,val)
		}
}