package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func cleanup() {
	fmt.Println("Time for cleanup")
	time.Sleep(1 * time.Second) // or runtime.Gosched() or similar per @misterbee
}

func main() {
	s1 := make(chan os.Signal, 1)
	signal.Notify(s1, syscall.SIGTERM, os.Interrupt)

	go func() {
		sig := <-s1
		fmt.Printf("This program is going to exit now because of %v...\n", s1)
		switch sig {
		case syscall.SIGTERM:
			fmt.Println("Someone sent us a SIGTERM, time for exit gracefully...")
		case os.Interrupt:
			fmt.Println("Ctrl-C received, time for exit gracefully...")
		}
		cleanup()
		os.Exit(1)
	}()

	for {
		fmt.Println("sleeping...")
		time.Sleep(10 * time.Second) // or runtime.Gosched() or similar per @misterbee
	}
}
