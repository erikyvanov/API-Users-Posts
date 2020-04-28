package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/erikyvanov/API-Users-Posts/db"
	"github.com/erikyvanov/API-Users-Posts/models"
)

// ConsultRelationHandler dice si los usuarios estan relacionados o no
func ConsultRelationHandler(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "No hay ID", http.StatusBadRequest)
		return
	}

	var t models.Relation
	t.UserID = IDUser
	t.RelationID = ID

	var result models.ConsultRelationResponse
	status, err := db.ConsultRelation(t)

	if err != nil || !status {
		result.Status = false
	} else {
		result.Status = true
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}
