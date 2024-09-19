package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
	end   endereco
}

type endereco struct {
	logradouro string
}

func main() {
	fmt.Println("Arquivo structs")

	var u usuario
	u.nome = "Davi"
	u.idade = 21
	fmt.Println(u)

	u2 := usuario{nome: "Davi", idade: 21, end: endereco{logradouro: "Rua dos Bobos"}}
	fmt.Println(u2.end.logradouro)
}
