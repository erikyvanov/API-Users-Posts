package db

import (
	"context"
	"time"

	"github.com/erikyvanov/API-Users-Posts/models"
	"go.mongodb.org/mongo-driver/bson"
)

// CheckUserExists recibe un email y dice si ya esta registrado.
func CheckUserExists(email string) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	dataBase := MongoConnection.Database("app1")
	collection := dataBase.Collection("users")

	filter := bson.M{"email": email}

	var result models.User

	err := collection.FindOne(ctx, filter).Decode(&result)
	ID := result.ID.Hex()

	if err != nil {
		return result, false, ID
	}

	return result, true, ID
}
