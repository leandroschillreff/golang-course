package main

import (
	"fmt"
)

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculosMatematicos(n1, n2 int8) (int8, int8, int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	multiplacao := n1 * n2
	divisao := n1 / n2
	return soma, subtracao, multiplacao, divisao
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func(texto string) string {
		fmt.Println(texto)
		return texto
	}

	resultado := f("Texto da função")
	fmt.Println(resultado)

	resultadoSoma, resultadoSubtracao, resultadoMultiplicacao, resultadoDivisao := calculosMatematicos(3, 3)
	fmt.Println(resultadoSoma)
	fmt.Println(resultadoSubtracao)
	fmt.Println(resultadoMultiplicacao)
	fmt.Println(resultadoDivisao)

	resultadoSoma1, _, _, _ := calculosMatematicos(3, 3)
	fmt.Println(resultadoSoma1)

}
