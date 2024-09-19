package main

func inverterSinal(numero int) int {
	return numero * -1
}

func inverterSinalComPonteiro(numero *int) int {
	*numero = *numero * -1
	return *numero
}

func main() {
	numero := 20
	numeroInvertido := inverterSinal(numero)
	println(numeroInvertido)
	println(numero)

	numeroInvertidoComPonteiro := inverterSinalComPonteiro(&numero)
	println(numero)
	println(numeroInvertidoComPonteiro)
}
