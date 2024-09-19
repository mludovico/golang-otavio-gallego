package main

func main() {
	f := closure()
	f()
}

func closure() func() {
	x := 10
	return func() {
		println(x)
	}
}
