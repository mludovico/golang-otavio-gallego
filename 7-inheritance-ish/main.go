package main

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	p1 := pessoa{"João", "Pedro", 20, 178}
	p2 := estudante{p1, "Engenharia", "USP"}
	p3 := estudante{
		pessoa{
			nome:      "Maria",
			sobrenome: "Silva",
			idade:     21,
			altura:    175,
		},
		"Medicina",
		"UNESP",
	}

	println(p1.nome)
	println(p2.nome)
	println(p2.faculdade)
	println(p3.nome)
}
