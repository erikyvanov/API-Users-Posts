package handlers

import (
	"io"
	"net/http"
	"os"

	"github.com/erikyvanov/API-Users-Posts/db"
)

// GetAvatar obtiene el avatar del usuario en la base de datos
func GetAvatar(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "No hay id", http.StatusBadRequest)
		return
	}

	user, err := db.SearchProfile(ID)
	if err != nil {
		http.Error(w, "No existe ese usuario "+err.Error(), http.StatusBadRequest)
		return
	}

	openAvatar, err := os.Open("uploads/avatars/" + user.Avatar)
	if err != nil {
		http.Error(w, "Imagen no encontrada", http.StatusBadRequest)
		return
	}
	defer openAvatar.Close()

	_, err = io.Copy(w, openAvatar)
	if err != nil {
		http.Error(w, "Error al envia la imagen", http.StatusBadRequest)
		return
	}
}
