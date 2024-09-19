package main

import (
	"fmt"
)

func teste1(i int) {
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
		break
	case 2:
		fmt.Println("two")
		break
	case 3:
		fmt.Println("three")
		break
	default:
		fmt.Println("unknown")
	}
}

func teste2(i int) {
	switch {
	case i == 1:
		fmt.Println("one")
		break
	case i == 2:
		fmt.Println("two")
		break
	case i == 3:
		fmt.Println("three")
		break
	case i > 3:
		fmt.Println("greater than three")
	}
}
func main() {
	teste1(1)
	teste2(4)
}
