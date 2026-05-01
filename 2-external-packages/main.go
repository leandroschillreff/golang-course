package main

import (
	"fmt"
	"github.com/badoux/checkmail"
	"module/auxiliary"
)

func main() {
	fmt.Println("Escrevendo do arquivo main.go")
	auxiliary.ToWrite()

	erro := checkmail.ValidateFormat("leandroschillreff@gmail.com")
	fmt.Println(erro)
}
