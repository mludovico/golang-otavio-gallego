package main

import (
	"errors"
	"fmt"
	"reflect"
)

func main() {
	numero1 := 1000000
	numero2 := 10000000000000000000.0
	fmt.Println(reflect.TypeOf(numero1) == reflect.TypeOf(numero2))

	var str string = "Texto"
	fmt.Println(str)

	char := 'B'
	fmt.Println(char, ": ", reflect.TypeOf(char))

	var texto string
	fmt.Println(texto, reflect.TypeOf(texto))

	var err error = errors.New("Erro interno")
	fmt.Println(err, reflect.TypeOf(err))
}
