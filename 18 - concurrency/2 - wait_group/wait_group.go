package main

import (
	"sync"
	"time"
)

func main() {
	waitGroup := new(sync.WaitGroup)

	waitGroup.Add(4)

	go func() {
		writeFiveTimes("Hello")
		waitGroup.Done()
	}()

	go func() {
		writeFiveTimes("World")
		waitGroup.Done()
	}()

	go func() {
		writeFiveTimes("Go routine 3")
		waitGroup.Done()
	}()

	go func() {
		writeFiveTimes("Go routine 4")
		waitGroup.Done()
	}()

	waitGroup.Wait()
}

func writeFiveTimes(text string) {
	for i := 0; i < 5; i++ {
		println(text)
		time.Sleep(1 * time.Second)
	}
}
