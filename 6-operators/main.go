package main

import (
	"fmt"
)

func main() {
	// ARITMETICOS
	// +, -, /, *, %

	soma := 1 + 2
	subtracao := 1 - 2
	multiplicacao := 1 * 2
	divisao := 10 / 2
	restoDaDivisao := 10 / 3
	fmt.Println(soma, subtracao, multiplicacao, divisao, restoDaDivisao)

	var numero1 int16 = 10
	var numero2 int32 = 25
	soma1 := numero1 + int16(numero2)
	fmt.Println(soma1)

	//Operadores relacionais
	fmt.Println(1 > 2)
	fmt.Println(1 < 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 != 2)

	//Operadores lógicos
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)

	//Operadores unários
	numero := 10
	numero++
	numero += 10
	fmt.Println(numero)
	numero --
	numero -= 5

	//Operador ternário
	
}
