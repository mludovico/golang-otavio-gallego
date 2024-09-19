package main

import (
	"fmt"
	"time"
)

func main() {
	channel := writer("Hello")
	for i := 0; i < 5; i++ {
		fmt.Println(<-channel)
	}
}

func writer(text string) <-chan string {
	channel := make(chan string)

	go func() {
		for {
			channel <- fmt.Sprintf("Value received: %s", text)
			time.Sleep(500 * time.Millisecond)
		}
	}()

	return channel
}
