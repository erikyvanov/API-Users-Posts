package db

import (
	"github.com/erikyvanov/API-Users-Posts/models"
	"golang.org/x/crypto/bcrypt"
)

// LoginTry Devuelve un usuario y dice si el usuario y contraseña coinciden
func LoginTry(email string, password string) (models.User, bool) {
	user, found, _ := CheckUserExists(email)

	if !found {
		return user, false
	}

	// Comparar contraseñas con bcrypt
	passwordBytes := []byte(password)
	passwordDB := []byte(user.Password)

	err := bcrypt.CompareHashAndPassword(passwordDB, passwordBytes)
	if err != nil {
		return user, false
	}

	return user, true
}
