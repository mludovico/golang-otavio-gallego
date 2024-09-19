package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string)

	go writeForever("Hello", channel)
	go writeForever("World", channel)

	fmt.Println("go routines started")

	for message := range channel {
		println(message)
	}

}

func writeForever(text string, channel chan string) {
	for i := 0; i < 5; i++ {
		channel <- text
		time.Sleep(1 * time.Second)
	}
	close(channel)
}
