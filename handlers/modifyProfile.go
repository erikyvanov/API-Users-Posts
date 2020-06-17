package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/erikyvanov/API-Users-Posts/db"
	"github.com/erikyvanov/API-Users-Posts/models"
)

// ModifyProfile es el handler para modificar el perfil
func ModifyProfile(w http.ResponseWriter, r *http.Request) {
	var u models.User

	err := json.NewDecoder(r.Body).Decode(&u)

	if err != nil {
		http.Error(w, "Datos invalidos "+err.Error(), http.StatusBadRequest)
		return
	}

	status, err := db.ModifyRegister(u, IDUser)
	if err != nil {
		http.Error(w, "Hubo un error al modificar los datos "+err.Error(), http.StatusBadRequest)
		return
	}

	if !status {
		http.Error(w, "No se logro modificar los datos en la base de datos", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
