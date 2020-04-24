package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/erikyvanov/API-Users-Posts/db"

	"github.com/erikyvanov/API-Users-Posts/models"
)

// NewPost crea una nueva publicacion
func NewPost(w http.ResponseWriter, r *http.Request) {
	var m models.SendPost
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		http.Error(w, "Error en el json "+err.Error(), http.StatusBadRequest)
		return
	}

	post := models.Post{
		UserID: IDUser,
		Body:   m.Message,
		Date:   time.Now(),
	}

	_, status, err := db.InsertPost(post)
	if err != nil {
		http.Error(w, "Error al insertar post "+err.Error(), http.StatusBadRequest)
		return
	}

	if !status {
		http.Error(w, "No se inserto post", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
