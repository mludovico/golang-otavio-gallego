package main

import "fmt"

func main() {
	fmt.Println("Control structures")

	numero := 10
	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor que 15")
	}

	if outroNumero := 20; outroNumero > 15 {
		fmt.Println("20 é maior que 15")
	}

	// fmt.Println(outroNumero) // erro: undefined: outroNumero

}
