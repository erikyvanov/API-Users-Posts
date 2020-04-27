package handlers

import (
	"net/http"

	"github.com/erikyvanov/API-Users-Posts/db"
)

// DeletePost hace una peticion a la base de datos para borrar un post
func DeletePost(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) == 0 {
		http.Error(w, "No hay postID", http.StatusBadRequest)
		return
	}

	err := db.DeletePost(ID, IDUser)
	if err != nil {
		http.Error(w, "Error al borrar el post "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
