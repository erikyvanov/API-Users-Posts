package handlers

import (
	"net/http"

	"github.com/erikyvanov/API-Users-Posts/db"
	"github.com/erikyvanov/API-Users-Posts/models"
)

// Relation crea una relacion entre usuarios
func Relation(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "No hay ID", http.StatusBadRequest)
		return
	}

	var t models.Relation
	t.UserID = IDUser
	t.RelationID = ID

	success, err := db.InsertRelation(t)
	if err != nil {
		http.Error(w, "Error al insertar relacion en la base de datos"+err.Error(), http.StatusBadRequest)
		return
	}

	if !success {
		http.Error(w, "Error al insertar relacion en la base de datos", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
