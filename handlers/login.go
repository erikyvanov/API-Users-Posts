package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/erikyvanov/API-Users-Posts/db"
	"github.com/erikyvanov/API-Users-Posts/jwt"
	"github.com/erikyvanov/API-Users-Posts/models"
)

// Login es el manejador para iniciar sesion en la app
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var u models.User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, "Usuario y/o Contraseña invalidos "+err.Error(), http.StatusBadRequest)
		return
	}
	if len(u.Email) == 0 {
		http.Error(w, "Email invalido.", http.StatusBadRequest)
		return
	}

	doc, exist := db.LoginTry(u.Email, u.Password)
	if !exist {
		http.Error(w, "Usuario y/o Contraseña invalidos.", http.StatusBadRequest)
		return
	}

	//Si se logro logear se envia un token
	jwtKey, err := jwt.GenerateJWT(doc)
	if err != nil {
		http.Error(w, "Error al generar el token "+err.Error(), http.StatusBadRequest)
		return
	}

	resp := models.LoginResponse{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}
