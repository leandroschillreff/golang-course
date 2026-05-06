package main

import (
	"errors"
	"fmt"
)

func main() {
	//int8, int16, int32, int64
	//int8 -> até 8 bits
	//int16 -> até 16 bits
	//int32 -> até 32 bits
	//int64 -> até 64 bits

	var numero int16 = 100
	fmt.Println(numero)

	var numero2 uint64 = 23456
	fmt.Println(numero2)

	// alias
	//INT32 = RUNE
	var numero3 rune = 123456
	fmt.Println(numero3)

	//alias
	//BYTE = UINT8
	var numero4 byte = 255
	fmt.Println(numero4)

	//float32, float64

	var numeroreal1 float32 = 123.45
	fmt.Println(numeroreal1)

	var numeroreal2 float64 = 1230000000.45
	fmt.Println(numeroreal2)

	var str string = "Texto"
	fmt.Println(str)

	var booleano bool
	fmt.Println(booleano)

	var erro error = errors.New("Erro interno")
	fmt.Println(erro)

}
