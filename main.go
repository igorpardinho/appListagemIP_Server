package main

import (
	"linha-de-comando/app"
	"log"
	"os"
)

func main(){

	aplicaçao := app.Gerar()

	if erro := aplicaçao.Run(os.Args); erro != nil{
		log.Fatal(erro)
	}

}
