package main

import (
	"fmt"
	"time"
)

func main() {
	presentTime := time.Now()
	fmt.Println(presentTime)

	//Default values to print the time
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createDate := time.Date(2022, time.November, 18, 44, 40, 0, 0, time.UTC)
	fmt.Println(createDate)
}
