package main

import "fmt"

func main() {
	func() {
		println("Hello from anonymous function without parameters and return")
	}()

	func(a int) {
		fmt.Printf(
			"Hello from anonymous function with parameter %d and without return\n",
			a,
		)
	}(10)

	s := func(a int) string {
		return fmt.Sprintf(
			"Hello from anonymous function with parameter %d and returning this string\n",
			a,
		)
	}(10)
	println(s)
}
