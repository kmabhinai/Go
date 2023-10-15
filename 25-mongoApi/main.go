package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kmabhinai/go-mongoapi/router"
)

func main()  {
	fmt.Println("MongoDb Api")
	r := router.Router()
	fmt.Println("Server is getting started...")
	log.Fatal(http.ListenAndServe(":3000",r))
	fmt.Println("Listening to the port 3000...")
}