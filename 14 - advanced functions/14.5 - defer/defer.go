package main

import "fmt"

func main() {
	defer fmt.Println("Hello from defer")
	fmt.Println("Hello from main")
	fmt.Println(getNum())
}

func getNum() int {
	defer fmt.Println("Returning soon")
	defer fmt.Println("Be patient")
	fmt.Println("now we're in the middle of the function")
	return 10
}
