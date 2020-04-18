package db

import (
	"context"
	"time"

	"github.com/erikyvanov/API-Users-Posts/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// InsertUser mete a la base de datos un usuario.
// Regresa (ID insertado, si se inserto el usuario, error)
func InsertUser(u models.User) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	dataBase := MongoConnection.Database("app1")
	collection := dataBase.Collection("users")

	// Cifrar password
	u.Password, _ = EncryptPassword(u.Password)

	// Meter usuario en la base de datos
	result, err := collection.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}

	// ID del usuario insertado
	objID, _ := result.InsertedID.(primitive.ObjectID)

	return objID.String(), true, nil
}
