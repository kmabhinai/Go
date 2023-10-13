package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main()  {
	router := mux.NewRouter()
	router.HandleFunc("/",serveHome).Methods("GET")

	// http.ListenAndServe(":4000",router)
	fmt.Println("Server is running on port 4000")
	log.Fatal(http.ListenAndServe(":4000",router))
}

func serveHome(res http.ResponseWriter,req *http.Request )  {
	res.Write([]byte("<h1>Kambhampati</h1>"))
}