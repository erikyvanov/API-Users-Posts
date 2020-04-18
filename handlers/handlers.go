package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/erikyvanov/API-Users-Posts/middlewares"
	"github.com/erikyvanov/API-Users-Posts/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// Handlers configura el puerto y los handlers
func Handlers() {
	router := mux.NewRouter()

	//Crear rutas
	router.HandleFunc("/user", middlewares.CheckDB(routers.Register)).Methods("POST")

	//Creamos el PORT
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
