package main

import (
	"log"

	"github.com/celemasmut/MicrobloggingProject/bd"
	"github.com/celemasmut/MicrobloggingProject/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexion con la BD")
		return
	}
	handlers.Manejadores() // muestra que la conexion fue exitosa.
}
