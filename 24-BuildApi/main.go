package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//Model for course
type Course struct{
	CourseID string  `json:"CourseID"`
	CourseName string `json:"CourseName"`
	CoursePrice int `json:"CoursePrice"`
	Author *Author `json:"Author"`
}

type Author struct{
	FullName string `json:"fullName"`
	Website string `json:"website"`
}

//fake Db
var courses []Course

//middleware, helper - file
func (c *Course) IsEmpty()bool  {
	// return c.CourseID=="" && c.CourseName==""
	return c.CourseName==""
}

func main()  {
	
}

// controllers - file

// serve home route
func serveHome(res http.ResponseWriter, req *http.Request)  {
	res.Write([]byte("<h1>Welcome to API</h1>"))
}

func getAllCourses(res http.ResponseWriter, req *http.Request){
	fmt.Println("Get all courses")
	res.Header().Set("Content-Type","application/json")
	json.NewEncoder(res).Encode(courses)
}

func getOneCourse(res http.ResponseWriter, req *http.Request)  {
	res.Header().Set("Content-Type","application/json")
	
	params := mux.Vars(req)
	//TODO: print params and chk
	
	for _,course := range courses{
		if course.CourseID == params["id"] {
			json.NewEncoder(res).Encode(course)
			return
		}
	}
	json.NewEncoder(res).Encode("No course found with given Id")
}

func createOneCourse(res http.ResponseWriter, req *http.Request)  {
	res.Header().Set("Content-Type","application/json")
	
	if req.Body == nil{
		json.NewEncoder(res).Encode("Please send some data")
	}

	var course Course
	json.NewDecoder(req.Body).Decode(&course)

	if course.IsEmpty() {
		json.NewEncoder(res).Encode("No Data inside the JSON")
		return
	}

	course.CourseID= strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(res).Encode(course)
	return
}

func updateOneCourse(res http.ResponseWriter, req *http.Request)  {
	res.Header().Set("Content-Type","application/json")
	
	params := mux.Vars(req)

	for index,course := range courses{
		if course.CourseID == params["id"]{
			courses = append(courses[:index],courses[index+1:]... )	
			var course Course
			json.NewDecoder(req.Body).Decode(&course)
			course.CourseID = params["id"]
			courses = append(courses, course)
			json.NewEncoder(res).Encode(course)
			return
		}
	}

	//TODO: send a res when id is not found
}

func deleteOneCourse(res http.ResponseWriter, req *http.Request)  {
	res.Header().Set("Content-Type","application/json")
	
	params:=mux.Vars(req)
	
	for index,course := range courses{
		if course.CourseID == params["id"]{
			courses = append(courses[:index],courses[index+1:]... )
			res.Write([]byte("The Course is Successfully deleted"))
			break
		}
	}
}