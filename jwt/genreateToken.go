package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/erikyvanov/API-Users-Posts/config"
	"github.com/erikyvanov/API-Users-Posts/models"
)

// GenerateJWT Genera un token para uusuario
func GenerateJWT(u models.User) (string, error) {
	payload := jwt.MapClaims{
		"_id":         u.ID.Hex(),
		"name":        u.Name,
		"lastName":    u.LastName,
		"description": u.Description,
		"birthday":    u.Birthday,

		"email": u.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(config.TokenKey)

	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}
