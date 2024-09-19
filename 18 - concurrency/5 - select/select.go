package main

import "time"

func main() {
	channel1, channel2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(500 * time.Millisecond)
			channel1 <- "Hello"
		}
	}()

	go func() {
		for {
			time.Sleep(2 * time.Second)
			channel2 <- "World"

		}
	}()

	for {
		select {
		case message := <-channel1:
			println(message)
		case message := <-channel2:
			println(message)
		}
	}
}
