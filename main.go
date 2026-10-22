package main

import (
	"log"
	"os"

	"teste/app"
)

func main() {
	aplicacao := app.Gerar()
	erro := aplicacao.Run(os.Args)
	if erro != nil {
		log.Fatal(erro)
	}
}
