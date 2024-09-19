package main

func main() {
	println(soma(1, 2, 3, 4, 5))
	println(soma(1, 2, 3))
	println(soma(1, 2))
	println(soma(1))
	println(soma())

	println(conta("soma", 1, 2, 3, 4, 5))

	slice := []int{1, 2, 3, 4, 5}
	println(subtracao(slice...))

	s := []string{"a", "b", "c"}
	t := []string{"test"}
	u := append(t, s...)
	println(&u)
}

func soma(numeros ...int) int {
	soma := 0
	for _, n := range numeros {
		soma += n
	}
	return soma
}

func subtracao(numeros ...int) int {
	subtracao := 0
	for i, n := range numeros {
		if i == 0 {
			subtracao = n
		} else {
			subtracao -= n
		}
	}
	return subtracao
}

func conta(operacao string, numeros ...int) int {
	switch operacao {
	case "soma":
		return soma(numeros...)
	case "subtracao":
		return subtracao(numeros...)
	default:
		return 0
	}
}
