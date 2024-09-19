package main

import (
	"fmt"
)

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Println("Salvando dados do usuário: ", u.nome)
}

func (u usuario) isOver18() bool {
	return u.idade >= 18
}

func (u usuario) completeBirthday() {

}

func completeBirthday(u *usuario) {
	u.idade++
}

func main() {
	usuario1 := usuario{"Usuario 1", 20}
	usuario1.salvar()
	fmt.Println(usuario1.isOver18())

	fmt.Println(usuario1)
	completeBirthday(&usuario1)
	fmt.Println(usuario1)
}
