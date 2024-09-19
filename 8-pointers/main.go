package main

import "fmt"

type user struct {
	name string
}

func main() {
	fmt.Println("Pointers")

	var var1 int = 10
	var var2 int = var1

	fmt.Println(var1, var2)

	var2++

	fmt.Println(var1, var2)

	user1 := user{name: "John"}
	user2 := user1

	fmt.Println(user1, user2)

	user2.name = "Doe"

	fmt.Println(user1, user2)

	var valor int = 10
	var ponteiro *int = &valor

	fmt.Println(valor, ponteiro)

	*ponteiro = 20

	fmt.Println(valor, ponteiro)

	fmt.Println(&valor, &user1.name)
}
