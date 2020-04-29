package db

import (
	"context"
	"time"

	"github.com/erikyvanov/API-Users-Posts/config"
	"github.com/erikyvanov/API-Users-Posts/models"
	"go.mongodb.org/mongo-driver/bson"
)

// ReadFollowedPosts lee los post de los usuarios siguiendo
func ReadFollowedPosts(ID string, page int) ([]models.ReturnFollowePost, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	dataBase := MongoConnection.Database(config.DBName)
	collection := dataBase.Collection("relations")

	skip := (page - 1) * 20
	conditions := make([]bson.M, 0)

	conditions = append(conditions, bson.M{"$match": bson.M{"userid": ID}})
	conditions = append(conditions, bson.M{
		"$lookup": bson.M{
			"from":         "posts",
			"localField":   "relationid",
			"foreignField": "userid",
			"as":           "posts",
		},
	})
	conditions = append(conditions, bson.M{"$unwind": "$posts"})
	conditions = append(conditions, bson.M{"$sort": bson.M{"post.date": -1}})
	conditions = append(conditions, bson.M{"$skip": skip})
	conditions = append(conditions, bson.M{"$limit": 20})

	cursor, err := collection.Aggregate(ctx, conditions)
	var results []models.ReturnFollowePost
	err = cursor.All(ctx, &results)
	if err != nil {
		return results, false
	}

	return results, true
}
