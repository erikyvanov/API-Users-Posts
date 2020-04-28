package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/erikyvanov/API-Users-Posts/db"
)

// ListUsers lista los usuarios
func ListUsers(w http.ResponseWriter, r *http.Request) {
	typeUSer := r.URL.Query().Get("type")
	pageStr := r.URL.Query().Get("page")
	search := r.URL.Query().Get("search")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		http.Error(w, "La pagina debe ser un entero mayor a 0", http.StatusBadRequest)
		return
	}

	result, status := db.ReadUsers(IDUser, int64(page), search, typeUSer)
	if !status {
		http.Error(w, "Error en la busqueda de usuarios", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}
