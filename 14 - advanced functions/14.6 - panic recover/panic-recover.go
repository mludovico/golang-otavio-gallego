package main

func main() {
	defer func() {
		if r := recover(); r != nil {
			println("Recovered from panic")
		}
	}()
	panic("Panic!")
}
