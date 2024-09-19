package main

import "strconv"

func main() {
	myInt, err := strconv.ParseInt("0", 8, 32)
	if err != nil {
		panic(err)
	}

	for myInt < 20 {
		println(myInt)
		myInt++
	}
}
