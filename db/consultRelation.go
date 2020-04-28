package db

import (
	"context"
	"time"

	"github.com/erikyvanov/API-Users-Posts/config"
	"github.com/erikyvanov/API-Users-Posts/models"
	"go.mongodb.org/mongo-driver/bson"
)

// ConsultRelation dice si hay un relacion entre 2 usuarios
func ConsultRelation(t models.Relation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	dataBase := MongoConnection.Database(config.DBName)
	collection := dataBase.Collection("relations")

	filter := bson.M{
		"userid":     t.UserID,
		"relationid": t.RelationID,
	}

	var result models.Relation
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return false, err
	}

	return true, nil
}
