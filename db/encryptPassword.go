package db

import "golang.org/x/crypto/bcrypt"

// EncryptPassword encripta una contraseña tipo string y la devuelve en string
func EncryptPassword(pass string) (string, error) {
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(pass), 6)

	return string(encryptedPassword), err
}
