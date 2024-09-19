package main

import "errors"

func main() {
	resultado := somar(10, 20)
	resultadoSoma, resultadoSubtracao, resultadoMultiplicacao, resultadoDivisao, erro := calculos(10, 0)
	if erro == nil {
		println(resultado, resultadoSoma, resultadoSubtracao, resultadoMultiplicacao, resultadoDivisao)
	} else {
		println("Erro ao executar a função calculos -", erro.Error())
	}
}

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculos(n1, n2 int8) (int8, int8, int8, float32, error) {
	var erro error = nil
	soma := n1 + n2
	subtracao := n1 - n2
	multiplicacao := n1 * n2
	var divisao float32
	if n2 == 0 {
		erro = errors.New("division by zero")
	} else {
		divisao = float32(n1) / float32(n2)
	}
	return soma, subtracao, multiplicacao, divisao, erro
}
