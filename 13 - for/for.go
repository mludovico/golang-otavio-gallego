package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("for loop")

	i := 0
	for i < 5 {
		i++
		time.Sleep(100 * time.Millisecond)
		fmt.Println(i)
	}

	for j := 0; j < 5; j++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(j)
	}

	fibonaci(34)

	nome := [3]string{"João", "Maria", "José"}
	for _, n := range nome {
		fmt.Println(n)
	}

	for indice, letra := range "PALAVRA" {
		fmt.Printf("Índice: %d, Letra: %c\n", indice, letra)
	}

	usuario := map[string]string{
		"nome":      "João",
		"sobrenome": "da Silva",
	}

	for chave, valor := range usuario {
		fmt.Printf("%s = %s\n", chave, valor)
	}

	for {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(i)
		i++
		if i == 50 {
			break
		}
	}

	value := true
	for value {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(i)
		i++
		if i == 70 {
			value = false
		}
	}
}

func fibonaci(n int) {
	for i := 0; i <= n; {
		println(i)
		if i == 0 {
			i++
		} else {
			i += i
		}
	}
}
