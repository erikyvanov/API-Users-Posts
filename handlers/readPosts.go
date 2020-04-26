package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/erikyvanov/API-Users-Posts/db"
)

// ReadPosts Lee los posts de un usuario en la base de datos y se los envia al usuario
func ReadPosts(w http.ResponseWriter, r *http.Request) {
	// Leer parametros de la URL
	ID := r.URL.Query().Get("id")
	if len(ID) == 0 {
		http.Error(w, "No hay ID", http.StatusBadRequest)
		return
	}

	if len(r.URL.Query().Get("page")) == 0 {
		http.Error(w, "No hay pagina", http.StatusBadRequest)
		return
	}

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "Numero de pagina no valida", http.StatusBadRequest)
		return
	}

	// Vamos a base de datos
	result, success := db.ReadPosts(ID, int64(page))
	if !success {
		http.Error(w, "Error al leer los posts", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(result)
}
