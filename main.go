package main

import (
	"domain-checker/app"
	"log"
	"os"
)

// chama função gerar do package app
func main() {
	aplicacao := app.Gerar()
	if erro := aplicacao.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}

}
