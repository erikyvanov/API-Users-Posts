package routers

import (
	"log"
	"net/http"
	"os"

	"github.com/erikyvanov/API-Users-Posts/handlers"
	"github.com/erikyvanov/API-Users-Posts/middlewares"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// Routers configura el puerto y los routers
func Routers() {
	router := mux.NewRouter()

	//Crear rutas
	router.HandleFunc("/user", middlewares.CheckDB(handlers.Register)).Methods("POST")
	router.HandleFunc("/login", middlewares.CheckDB(handlers.Login)).Methods("POST")
	router.HandleFunc("/view-profile", middlewares.CheckDB(middlewares.ValidateJWT(handlers.ViewProfile))).Methods("GET")
	router.HandleFunc("/modify-profile", middlewares.CheckDB(middlewares.ValidateJWT(handlers.ModifyProfile))).Methods("PUT")

	router.HandleFunc("/post", middlewares.CheckDB(middlewares.ValidateJWT(handlers.NewPost))).Methods("POST")
	router.HandleFunc("/posts", middlewares.CheckDB(middlewares.ValidateJWT(handlers.ReadPosts))).Methods("GET")
	router.HandleFunc("/post", middlewares.CheckDB(middlewares.ValidateJWT(handlers.DeletePost))).Methods("DELETE")

	router.HandleFunc("/uploadAvatar", middlewares.CheckDB(middlewares.ValidateJWT(handlers.UploadAvatar))).Methods("POST")
	router.HandleFunc("/getAvatar", middlewares.CheckDB(handlers.GetAvatar)).Methods("GET")

	router.HandleFunc("/relation", middlewares.CheckDB(middlewares.ValidateJWT(handlers.Relation))).Methods("POST")
	router.HandleFunc("/relation", middlewares.CheckDB(middlewares.ValidateJWT(handlers.DelRelation))).Methods("DELETE")
	router.HandleFunc("/consultRelation", middlewares.CheckDB(middlewares.ValidateJWT(handlers.ConsultRelationHandler))).Methods("GET")

	//Creamos el PORT
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
