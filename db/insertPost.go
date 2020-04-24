package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/erikyvanov/API-Users-Posts/config"
	"github.com/erikyvanov/API-Users-Posts/models"
	"go.mongodb.org/mongo-driver/bson"
)

// InsertPost mete una publiccacion de un usuario a a la base de datos
// Regresa el ID del post, si se inserto en la base de datos y un error
func InsertPost(p models.Post) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	dataBase := MongoConnection.Database(config.DBName)
	collection := dataBase.Collection("posts")

	register := bson.M{
		"userid": p.UserID,
		"body":   p.Body,
		"date":   p.Date,
	}

	result, err := collection.InsertOne(ctx, register)
	if err != nil {
		return string(""), false, err
	}

	ObjectID, _ := result.InsertedID.(primitive.ObjectID)

	return ObjectID.Hex(), true, nil
}
