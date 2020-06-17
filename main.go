package main

import (
	"log"

	"github.com/erikyvanov/API-Users-Posts/db"
	"github.com/erikyvanov/API-Users-Posts/routers"
)

func main() {
	if !db.ConnectionOK() {
		log.Fatal("No hay conexión a la base de datos.")
	} else {
		routers.Routers()
	}
}
