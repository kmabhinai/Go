package main

import (
	"fmt"
	"sync"
)

func main() {
	myCh := make(chan int, 1) // ,1 is for 1 buffered channel
	wg := &sync.WaitGroup{}

	// myCh <- 5
	// fmt.Println(<-myCh)

	wg.Add(2)
	go func(ch <-chan int, wg *sync.WaitGroup) { //send only channel
		fmt.Println(<-myCh)
		wg.Done()
	}(myCh, wg)

	go func(ch chan<- int, wg *sync.WaitGroup) { // receive only channel
		myCh <- 5
		myCh <- 6
		close(myCh)
		wg.Done()
	}(myCh, wg)
	wg.Wait()
}
