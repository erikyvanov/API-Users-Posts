package db

import (
	"context"
	"time"

	"github.com/erikyvanov/API-Users-Posts/config"
	"github.com/erikyvanov/API-Users-Posts/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// SearchProfile busca un perfil y lo devuelve
func SearchProfile(ID string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	dataBase := MongoConnection.Database(config.DBName)
	col := dataBase.Collection("users")

	var profile models.User

	IDbson, _ := primitive.ObjectIDFromHex(ID)

	filter := bson.M{
		"_id": IDbson,
	}

	err := col.FindOne(ctx, filter).Decode(&profile)
	profile.Password = ""
	if err != nil {
		return profile, err
	}

	return profile, nil
}
