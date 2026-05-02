package main

import "fmt"

func main() {
	var variavel1 string = "variável 1"
	fmt.Println(variavel1)

	variavel2 := "variável 2"
	fmt.Println(variavel2)

	var (
		variavel3 string = "variável 3"
		variavel4 string = "variável 4"
	)
	fmt.Println(variavel3)
	fmt.Println(variavel4)

	variavel5, variavel6 := "variável 5", "variável 6"
	fmt.Println(variavel5)
	fmt.Println(variavel6)

	const constante1 string = "constante 1"
	fmt.Println(constante1)

	const (
		constante2 string = "constante 2"
		constante3 string = "constante 3"
	)
	fmt.Println(constante2)
	fmt.Println(constante3)

	const constante4, constante5 string = "constante 4", "constante 5"
	fmt.Println(constante4)
	fmt.Println(constante5)
}
