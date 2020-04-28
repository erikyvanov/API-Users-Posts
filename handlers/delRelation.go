package handlers

import (
	"net/http"

	"github.com/erikyvanov/API-Users-Posts/db"

	"github.com/erikyvanov/API-Users-Posts/models"
)

// DelRelation borra una relacion
func DelRelation(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "No hay ID", http.StatusBadRequest)
		return
	}

	var t models.Relation
	t.UserID = IDUser
	t.RelationID = ID

	_, err := db.DeleteRelation(t)
	if err != nil {
		http.Error(w, "No se borro la relacion "+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
