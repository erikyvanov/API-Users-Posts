package db

import (
	"context"
	"log"
	"time"

	"github.com/erikyvanov/API-Users-Posts/config"
	"github.com/erikyvanov/API-Users-Posts/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ReadPosts lee todos los posts de un usuario en la base de datos
func ReadPosts(ID string, page int64) ([]*models.ReturnPost, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	dataBase := MongoConnection.Database(config.DBName)
	collection := dataBase.Collection("posts")

	var results []*models.ReturnPost

	filter := bson.M{"userid": ID}

	options := options.Find()
	//Dice cual es el limite de post a enviar
	options.SetLimit(20)
	//Va a ordenar los documentos por fecha y orden descendente
	options.SetSort(bson.D{{Key: "date", Value: -1}})
	//Uso de paginacion
	options.SetSkip((page - 1) * 20)

	cursor, err := collection.Find(ctx, filter, options)
	if err != nil {
		log.Fatal(err.Error())
		return results, false
	}

	// recorre los resultados para gardarlos en el arreglo
	for cursor.Next(context.TODO()) {
		var post models.ReturnPost
		err := cursor.Decode(&post)
		if err != nil {
			return results, false
		}

		results = append(results, &post)
	}

	return results, true
}
