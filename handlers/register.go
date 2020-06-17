package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/erikyvanov/API-Users-Posts/db"
	"github.com/erikyvanov/API-Users-Posts/models"
)

// Register es un endpoint para poder registrarse a la app.
func Register(w http.ResponseWriter, r *http.Request) {
	var newUser models.User

	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Error al recibir los datos: "+err.Error(), http.StatusBadRequest)
		return
	}

	if len(newUser.Email) == 0 {
		http.Error(w, "Se requiere un email.", http.StatusBadRequest)
		return
	}
	if len(newUser.Email) < 6 {
		http.Error(w, "La contraseña tiene que tener minimo 6 caracteres.", http.StatusBadRequest)
		return
	}

	_, exists, _ := db.CheckUserExists(newUser.Email)
	if exists {
		http.Error(w, "Ya hay un usuario registrado con ese email", http.StatusBadRequest)
		return
	}

	_, status, err := db.InsertUser(newUser)

	if err != nil {
		http.Error(w, "Hubo un error al crear el usuario: "+err.Error(), http.StatusBadRequest)
		return
	}

	if !status {
		http.Error(w, "No se inserto el usuario: ", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
