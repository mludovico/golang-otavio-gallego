package main

import "time"

func main() {
	go writeForever("Hello")
	go writeForever("World")
}

func writeForever(text string) {
	for {
		println(text)
		time.Sleep(1 * time.Second)
	}
}
