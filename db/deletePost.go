package db

import (
	"context"
	"errors"
	"time"

	"github.com/erikyvanov/API-Users-Posts/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// DeletePost Borra un post de la base de datos
func DeletePost(IDPost string, IDUser string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	dataBase := MongoConnection.Database(config.DBName)
	collection := dataBase.Collection("posts")

	objID, _ := primitive.ObjectIDFromHex(IDPost)
	filter := bson.M{
		"_id":    objID,
		"userid": IDUser,
	}

	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("no se encontro el post")
	}

	return nil
}
