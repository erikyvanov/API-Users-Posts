package handlers

import (
	"errors"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/erikyvanov/API-Users-Posts/config"
	"github.com/erikyvanov/API-Users-Posts/db"
	"github.com/erikyvanov/API-Users-Posts/models"
)

// Email guarda el email del usuario que esta siendo procesado (autenticacion token)
var Email string

// IDUser guarda el ID del usuario que esta siendo procesado (autenticacion token)
var IDUser string

// ProcessToken valida un token
func ProcessToken(tk string) (*models.Claim, bool, string, error) {
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de token invalido")
	}
	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return config.TokenKey, nil
	})

	if err == nil {
		_, exist, _ := db.CheckUserExists(claims.Email)
		if exist {
			Email = claims.Email
			IDUser = claims.ID.Hex()
		}

		return claims, exist, IDUser, nil
	}

	if !tkn.Valid {
		return claims, false, string(""), errors.New("token invalido")
	}

	return claims, false, string(""), err
}
