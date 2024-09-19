package main

func calculosAlgebricos(a, b int) (soma, subtracao int) {
	soma = a + b
	subtracao = a - b
	return
}

func main() {
	soma, subtracao := calculosAlgebricos(10, 5)
	println(soma, subtracao)
}
