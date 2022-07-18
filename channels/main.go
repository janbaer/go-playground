package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan string)

	wg.Add(2)

	go func(ch chan string, wg *sync.WaitGroup) {
		fmt.Println(<-ch)
		wg.Done()
	}(ch, wg)

	go func(ch chan string, wg *sync.WaitGroup) {
		ch <- "Hello world"
		wg.Done()
	}(ch, wg)

	wg.Wait()
}
