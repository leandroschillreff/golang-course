package main

import (
	"fmt"
)

type usuario struct {
	nome      string
	sobrenome string
	idade     uint8
}

func main() {
	var u usuario
	u.nome = "Leandro"
	u.sobrenome = "Schillreff"
	u.idade = 25
}
