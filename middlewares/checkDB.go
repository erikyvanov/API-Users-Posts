package middlewares

import (
	"net/http"

	"github.com/erikyvanov/API-Users-Posts/db"
)

// CheckDB es un middleware que comprueba que haya conexión con la base de datos
func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !db.ConnectionOK() {
			http.Error(w, "Error de conexión con la base de datos.", http.StatusInternalServerError)
			return
		}
		next.ServeHTTP(w, r)
	}
}
