package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/erikyvanov/API-Users-Posts/db"
)

// ReadPostsRelations busca los posts con las relaciones
func ReadPostsRelations(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "No hay pagina", http.StatusBadRequest)
		return
	}

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "La pagina debe ser un numero entero mayor a 0", http.StatusBadRequest)
		return
	}

	res, status := db.ReadFollowedPosts(IDUser, page)
	if !status {
		http.Error(w, "Error al leer los posts", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}
