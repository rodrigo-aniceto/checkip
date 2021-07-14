package main

import (
	"checkip/app"
	"log"
	"os"
)

/*
exemplos de chamada:
./checkip ip --host google.com.br
./checkip servidores --host google.com.br

*/
func main() {

	aplicacao := app.Gerar()
	if erro := aplicacao.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}

}
