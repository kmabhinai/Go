package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}
var wg sync.WaitGroup //pointer
var mut sync.Mutex    //pointer

func main() {
	// go print6("Hello")
	// print6("World")

	websiteList := []string{
		"https://google.com",
		"https://fb.com",
		"https://github.com",
		"https://go.dev",
		"https://knls.in",
	}

	for _, web := range websiteList {
		go getStatusCode(web)
		wg.Add(1)
	}

	wg.Wait()
	fmt.Println(signals)
}

// func print6(s string){
// 	for i:=0;i<6;i++{
// 		time.Sleep(3*time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endpoint string) {
	defer wg.Done() // decrement the counter when function returns or panics
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("failed to fetch endpoint")
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()

		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
	}
}
