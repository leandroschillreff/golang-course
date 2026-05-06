package main

import (
	"fmt"
)

type usuario struct {
	nome  string
	idade uint8
}

func main() {
	var u usuario
	u.nome = "Leandro Schillreff"
	u.idade = 25
}
