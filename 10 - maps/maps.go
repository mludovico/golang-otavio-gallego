package main

import (
	"fmt"
)

func main() {
	usuario1 := map[string]int{
		"key1": 1,
		"key2": 2,
	}

	usuario1["key3"] = 3

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "João",
			"segundo":  "Silva",
		},
		"endereco": {
			"rua":  "Rua 1",
			"cep":  "12345-678",
			"pais": "Brasil",
		},
	}

	usuario2["contato"] = map[string]string{
		"email": "mm@mm.com",
	}

	delete(usuario2, "endereco")
	fmt.Println(usuario1, usuario2)
}
